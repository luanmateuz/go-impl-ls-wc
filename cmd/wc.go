package cmd

import (
	"bufio"
	"io"
)

func Wc(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	if countLines {
		scanner.Split(bufio.ScanLines)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}

	return wc
}
