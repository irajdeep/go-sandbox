package linecount

import (
	"bufio"
	"bytes"
)

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		*l++
	}
	return int(*l), nil
}
