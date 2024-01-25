package main

import (
	"bufio"
	"io"
	"os"
)

// Counter is a function that takes a bufio.SplitFunc and count the resulting items
type Counter func(io.Reader) (int, error)

type TypeCounter func(io.Reader, bufio.SplitFunc) (int, error)

func counter(f io.Reader, splitter bufio.SplitFunc) (int, error) {
	s := bufio.NewScanner(f)
	s.Split(splitter)
	n := 0
	for s.Scan() {
		n++
	}
	return n, nil
}

func Bytes(f io.Reader) (int, error) {
	return counter(f, bufio.ScanBytes)
}

func Chars(f io.Reader) (int, error) {
	return counter(f, bufio.ScanRunes)
}

func Words(f io.Reader) (int, error) {
	return counter(f, bufio.ScanWords)
}

func Lines(f io.Reader) (int, error) {
	return counter(f, bufio.ScanLines)
}

func Count(fileName string, counter Counter) (int, error) {
	if fileName == "" {
		f := os.Stdin
		return counter(f)
	}

	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return counter(f)
}
