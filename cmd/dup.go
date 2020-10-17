package cmd

import (
	"fmt"
	"log"
	"ter_dbf_research/src/statistics"
	"ter_dbf_research/src/statprint"

	"github.com/spf13/cobra"
)

var dupShort bool

var dupCmd = &cobra.Command{
	Use:   "dup",
	Short: "Counts duplicates for separated fields for one record in dbf file",
	Run: func(cmd *cobra.Command, args []string) {
		provider := getProvider(args)

		for _, path := range provider.Files {
			stat, err := statistics.NewStat(path)
			if err != nil {
				log.Fatal(err)
			}
			dupStat, err := stat.Dup()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Statistics for file %s\n", path)
			statprint.DupStat(dupStat, !dupShort)
		}
	},
}

func getDupCmd() *cobra.Command {
	dupCmd.Flags().BoolVar(&dupShort, "short", false, "print short statistics")

	return dupCmd
}
