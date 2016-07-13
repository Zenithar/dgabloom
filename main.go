package main

import (
	"fmt"
	"time"

	"zenithar.org/go/dgabloom/generators"

	// Import all generators to register them
	_ "zenithar.org/go/dgabloom/generators/locky"
)

func main() {
	now := time.Now().UTC()
	fmt.Printf("%s\n", now.Format("2006-02-01"))

	for name, generator := range generators.Generators {
		fmt.Printf("\n# %s\n", name)

		domains, err := generator.Generate(now, nil)
		if err != nil {
			panic(err)
		}

		for _, domain := range domains {
			fmt.Printf("%s\n", domain)
		}
	}
}
