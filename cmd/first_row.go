package cmd

import (
	"fmt"
	"log"
	"ter_dbf_research/src/statistics"
	"ter_dbf_research/src/statprint"

	"github.com/spf13/cobra"
)

var firstRowCmd = &cobra.Command{
	Use:   "first-row",
	Short: "Checks if the first row contains valid data",
	Run: func(cmd *cobra.Command, args []string) {
		provider := getProvider(args)

		for _, path := range provider.Files {
			stat, err := statistics.NewStat(path)
			if err != nil {
				log.Fatal(err)
			}
			firstRowStat, err := stat.FirstRow()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Statistics for file %s\n", path)
			statprint.FirstRowStat(firstRowStat)
		}
	},
}
