package cmd

import (
	"bytes"
	"testing"
)

func TestWc(t *testing.T) {
	words := bytes.NewBufferString("word1 word2 word3")

	exp := 3
	res := Wc(words, false)

	if exp != res {
		t.Errorf("Expected %d, got %d instead\n", exp, res)
	}
}

func TestWcLines(t *testing.T) {
	lines := bytes.NewBufferString("line01\nline02\nline03\nline04\nline05")

	exp := 5
	res := Wc(lines, true)

	if exp != res {
		t.Errorf("Expected %d, got %d instead\n", exp, res)
	}
}
