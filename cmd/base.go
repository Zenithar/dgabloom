package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// MainCmd is DgaBloom's root command.
// Every other command attached to HugoCmd is a child command to it.
var MainCmd = &cobra.Command{
	Use:   "dgabloom",
	Short: "Check dga status quickly using bloom filters",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := InitializeConfig(); err != nil {
			return err
		}

		return cmd.Usage()
	},
}

// Flag values
var (
	databasePath string
)

func init() {
	MainCmd.PersistentFlags().StringVarP(&databasePath, "dbpath", "d", "$HOME/.dgabloom", "filesystem path where to store/read the bloom filter.")

	// Bash completion
	MainCmd.PersistentFlags().SetAnnotation("dbpath", cobra.BashCompFilenameExt, []string{})
}

// InitializeConfig prepares the flag environment
func InitializeConfig(subCmdVs ...*cobra.Command) error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("dgabloom")

	loadDefaultSettings()

	if databasePath != "" {
		viper.Set("DatabasePath", os.ExpandEnv(databasePath))
	}

	return nil
}

func initCommonFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&databasePath, "dbpath", "d", "$HOME/.dgabloom", "filesystem path whete to store/read the bloom filter.")

	// Set bash-completion
	cmd.Flags().SetAnnotation("dbpath", cobra.BashCompSubdirsInDir, []string{})
}

func loadDefaultSettings() {
	viper.SetDefault("DatabasePath", os.ExpandEnv("$HOME/.dgabloom"))
}

// Execute adds all child commands to the root command HugoCmd and sets flags appropriately.
func Execute() {

	AddCommands()

	if err := MainCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// AddCommands adds child commands to the root command MainCmd.
func AddCommands() {
	MainCmd.AddCommand(buildCmd)
	MainCmd.AddCommand(checkCmd)
}
