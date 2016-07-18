package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/willf/bloom"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check domain presence in the bloom filter",
}

func init() {
	initCommonFlags(checkCmd)

	checkCmd.RunE = check
}

func check(cmd *cobra.Command, args []string) error {
	if err := InitializeConfig(checkCmd); err != nil {
		return err
	}

	// Initialize bloom filter 1M max elements, 5 hash
	bf := bloom.New(1000000, 5)

	// Check if bloom filter file exists
	fileName := viper.GetString("DatabasePath")
	if _, err := os.Stat(fileName); err == nil {
		bff, err2 := os.Open(fileName)
		if err2 != nil {
			panic(err2)
		}
		defer bff.Close()

		if _, err3 := bf.ReadFrom(bff); err != nil {
			panic(err3)
		}
	} else {
		fmt.Println("You must build a filter first, use 'build' command.")
		return err
	}

	if len(args) > 0 {
		for _, domain := range args {
			fmt.Printf("%s,%v\n", domain, bf.TestString(domain))
		}
	}

	return nil
}
