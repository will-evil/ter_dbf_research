package cmd

import (
	"log"

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
}
