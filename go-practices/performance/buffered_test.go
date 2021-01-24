package performance

import (
	"bufio"
	"os"
	"testing"
)

func BenchmarkWriteFile(b *testing.B) {

	for n := 0; n < b.N; n++ {
		f, err := os.Create("/tmp/tests.txt")
		if err != nil {
			panic(err)
		}
		for i := 0; i < 10000; i++ {
			f.WriteString("some text\n")
		}
		f.Close()
	}
}

func BenchmarkWriteFileBuffered(b *testing.B) {
	for n := 0; n < b.N; n++ {
		f, err := os.Create("/tmp/tests.txt")
		if err != nil {
			panic(err)
		}

		w := bufio.NewWriter(f)
		for i := 0; i < 10000; i++ {
			w.WriteString("some text!\n")
		}

		w.Flush()
		f.Close()
	}
}
