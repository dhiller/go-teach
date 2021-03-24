package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := "http.log.gz"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var r io.Reader = file
	if strings.HasSuffix(fileName, ".gz") {
		r, err = gzip.NewReader(file)
		if err != nil {
			log.Fatal(err)
		}
	}

	hash := sha1.New() // reminder: naming - not "NewSha1"
	if _, err := io.Copy(r, hash); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", hash.Sum(nil))

	// FIXME the checksum is not correct!
}

func maybeClose(r io.Reader) error {
	// c := r.(io.Closer) // will panic if r does not implement Closer
	c, ok := r.(io.Closer) // type assertion
	if !ok {
		return nil
	}

	return c.Close()
}
