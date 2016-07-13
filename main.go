package main

import (
	"fmt"
	"os"
	"time"

	"github.com/willf/bloom"

	"zenithar.org/go/dgabloom/generators"
	// Import all generators to register them
	_ "zenithar.org/go/dgabloom/generators/locky"
)

func main() {
	// It is a good day to die !
	now := time.Now().UTC()

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

	// 30 days loop
	for day := 0; day < 30; day++ {

		// For each registered generator
		for name, generator := range generators.Generators {
			domains, err := generator.Generate(now, nil)
			if err != nil {
				panic(err)
			}

			// For each generated domains
			for _, domain := range domains {
				// Print CSV
				fmt.Printf("%s,%s,%s\n", name, now.Format("2006-01-02"), domain)
				// Add to bloom filter
				bf.AddString(domain)
			}
		}

		// Substract 1 day
		now = now.AddDate(0, 0, -1)
	}

	// Save bloom filter
	bff, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer bff.Close()

	// Dump to file
	bf.WriteTo(bff)
}
