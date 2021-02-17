package main

import (
	"fmt"
	"github.com/dhiller/go-teach/nlp"
	"os"
	"strings"
)

func main() {
	fmt.Println(nlp.Tokenize(strings.Join(os.Args[1:], " ")))
}
