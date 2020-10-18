package main

import (
	"fmt"
	"io"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// configure clusterIP, podIP, port and protocol
	clusterIP := "100.65.248.49"
	podIP := "100.96.2.50"
	port := 80
	proto := "tcp"
	fmt.Println("Starting node proxy")
	addRedirectRules(clusterIP, port, proto)
	createProxy(podIP, port, proto)
}

func addRedirectRules(clusterIP string, port int, proto string) error {
	p := strconv.Itoa(port)
	cmd := exec.Command("iptables", "-t", "nat", "-A", "OUTPUT", "-p", "tcp",
		"-d", clusterIP, "--dport", p, "-j", "REDIRECT", "--to-port", p)
	return cmd.Run()
}

func createProxy(podIP string, port int, proto string) {
	host := ""
	listener, err := net.Listen(proto, net.JoinHostPort(host, strconv.Itoa(port)))
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("successfully created tcp server ...")

	for {
		inConn, _ := listener.Accept()
		// if err != nil {
		// 	panic(err)
		// }
		fmt.Println("connection received ...")

		outConn, err := net.Dial(proto, net.JoinHostPort(podIP, strconv.Itoa(port)))
		if err != nil {
			panic(err)
		}
		fmt.Println("sucessfully dialed pod ....")

		go func(in, out *net.TCPConn) {
			var wg sync.WaitGroup
			wg.Add(2)
			fmt.Printf("Proxying %v <-> %v <-> %v <-> %v\n",
				in.RemoteAddr(), in.LocalAddr(), out.LocalAddr(), out.RemoteAddr())
			go copyBytes(in, out, &wg)
			go copyBytes(out, in, &wg)
			wg.Wait()
		}(inConn.(*net.TCPConn), outConn.(*net.TCPConn))
	}

}

func copyBytes(dst, src *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()
	if _, err := io.Copy(dst, src); err != nil {
		if !strings.HasSuffix(err.Error(), "use of closed network connection") {
			fmt.Printf("io.Copy error: %v", err)
		}
	}
	dst.Close()
	src.Close()
}
