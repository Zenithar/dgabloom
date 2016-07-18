package cmd

import (
	"fmt"
	"os"
	"time"

	"zenithar.org/go/dgabloom/generators"
	// Import all generators
	_ "zenithar.org/go/dgabloom/generators/locky"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/willf/bloom"
)

var (
	beforeDayCount int
	afterDayCount  int
	exportCSV      bool
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build or update the bloom filter whith new data.",
}

func init() {
	initCommonFlags(buildCmd)

	buildCmd.Flags().IntVarP(&beforeDayCount, "before", "b", 30, "Sets the number of day before today to generate.")
	buildCmd.Flags().IntVarP(&afterDayCount, "after", "a", 10, "Sets the number of day after today to generate.")
	buildCmd.Flags().BoolVarP(&exportCSV, "csv", "", false, "Export generated data as CSV.")

	buildCmd.RunE = build
}

func dgaGenerate(now time.Time, bf *bloom.BloomFilter) {
	// For each registered generator
	for name, generator := range generators.Generators {
		domains, err := generator.Generate(now, nil)
		if err != nil {
			panic(err)
		}

		// For each generated domains
		for _, domain := range domains {
			// Print CSV
			if exportCSV {
				fmt.Printf("%s,%s,%s\n", name, now.Format("2006-01-02"), domain)
			}
			// Add to bloom filter
			bf.TestAndAddString(domain)
		}
	}
}

func build(cmd *cobra.Command, args []string) error {
	if err := InitializeConfig(buildCmd); err != nil {
		return err
	}

	// It is a good day to die !
	now := time.Now().UTC()

	// Initialize bloom filter 1M max elements, 5 hash
	bf := bloom.New(1000000, 5)

	// Check if bloom filter file exists
	fileName := viper.GetString("DatabasePath")
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

	// future days loop
	for day := 0; day < afterDayCount; day++ {
		dgaGenerate(now, bf)
		// Add 1 day
		now = now.AddDate(0, 0, 1)
	}

	// Reset to today
	now = time.Now().UTC()

	// past days loop
	for day := 0; day < beforeDayCount; day++ {
		dgaGenerate(now, bf)
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
	_, err = bf.WriteTo(bff)
	if err != nil {
		panic(err)
	}

	return nil
}
