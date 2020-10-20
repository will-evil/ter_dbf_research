package cmd

import (
	"log"
	"ter_dbf_research/src/filesprovider"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "dbf-research",
	Short:   "Research dbf file with list of terrorist and terrorist organizations",
	Example: "dbf-research spaces terrorist_list.dbf",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute runs root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(spacesCmd)
	rootCmd.AddCommand(getDupCmd())
	rootCmd.AddCommand(dateCmd)
	rootCmd.AddCommand(firstRowCmd)
}

func getProvider(args []string) *filesprovider.FileProvider {
	provider, err := filesprovider.NewFileProvider(args)
	if err != nil {
		log.Fatal(err)
	}

	if len(provider.Files) == 0 {
		log.Fatalf("you are not provided any files for research")
	}

	return provider
}
