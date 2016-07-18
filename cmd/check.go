package cmd

import (
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
	checkCmd.RunE = build
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
		bff, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer bff.Close()

		if _, err := bf.ReadFrom(bff); err != nil {
			panic(err)
		}
	}

	// Save bloom filter
	bff, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer bff.Close()

	// Dump to file
	_, err = bf.WriteTo(bff)
	return err
}
