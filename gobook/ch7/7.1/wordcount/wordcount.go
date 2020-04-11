package wordcount

import (
	"bufio"
	"strings"
)

//ByteWordsCounter : counts number of words in words separated by space
type ByteWordsCounter int

func (b *ByteWordsCounter) Write(p []byte) (int, error) {
	inp := string(p)

	scanner := bufio.NewScanner(strings.NewReader(inp))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*b += ByteWordsCounter(len(scanner.Text()))
	}
	return len(p), nil
}
