package cmd

import (
	"fmt"
	"log"
	"ter_dbf_research/src/statistics"
	"ter_dbf_research/src/statprint"

	"github.com/spf13/cobra"
)

var spacesCmd = &cobra.Command{
	Use:   "spaces",
	Short: "Does a spaces researching of the file",
	Long:  "Researches if there are spaces in fields, which may be separated in dbf file",
	Run: func(cmd *cobra.Command, args []string) {
		provider := getProvider(args)

		for _, path := range provider.Files {
			stat, err := statistics.NewStat(path)
			if err != nil {
				log.Fatal(err)
			}
			spaceStat, err := stat.Spaces()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Statistics for file %s\n", path)
			statprint.SpacesStat(spaceStat)
		}
	},
}
