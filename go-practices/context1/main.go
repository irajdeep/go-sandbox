// a way to terminate go routines
package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8090/hello", nil)
	if err != nil {
		log.Println("err:", err)
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
