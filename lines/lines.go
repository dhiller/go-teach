package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	s := NewLineScanner(os.Stdin)
	for s.Scan() {
		fmt.Printf("LINE %d: %s\n", s.LineNum, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

func (s *LineScanner) Scan() bool {
	s.LineNum++
	return s.Scanner.Scan()
	// return s.Scan() // infinite recursion!
}

func NewLineScanner(r io.Reader) *LineScanner {
	scanner := bufio.NewScanner(r)
	return &LineScanner{Scanner: *scanner}
}

type LineScanner struct {
	LineNum int
	bufio.Scanner
}
