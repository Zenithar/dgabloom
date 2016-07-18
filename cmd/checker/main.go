package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/willf/bloom"
)

var (
	domainStr = flag.String("domain", "", "Domain to check.")
)

func init() {
	flag.Parse()
}

func main() {
	// Initialize bloom filter 1M max elements, 5 hash
	bf := bloom.New(1000000, 5)

	// Check if bloom filter file exists
	fileName := os.ExpandEnv("$HOME/.dgabloom")
	if _, err := os.Stat(fileName); err == nil {
		bff, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer bff.Close()

		if _, err := bf.ReadFrom(bff); err != nil {
			panic(err)
		}
	}

	// Test DGA to the bloom filter
	fmt.Printf("%v", bf.TestString(*domainStr))
}
