package cmd

import (
	"fmt"
	"log"
	"ter_dbf_research/src/statistics"
	"ter_dbf_research/src/statprint"

	"github.com/spf13/cobra"
)

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "Analyzes date formats in fields GR, CB_DATE and CE_DATE",
	Run: func(cmd *cobra.Command, args []string) {
		provider := getProvider(args)

		for _, path := range provider.Files {
			stat, err := statistics.NewStat(path)
			if err != nil {
				log.Fatal(err)
			}
			dateStat, err := stat.Date()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Statistics for file %s\n", path)
			statprint.DateStat(dateStat)
		}
	},
}
