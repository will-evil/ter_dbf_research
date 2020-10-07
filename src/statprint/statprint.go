package statprint

import (
	"fmt"
	"strings"
	"ter_dbf_research/src/statistics"
)

// SpacesStat prints statistics about spaces.
func SpacesStat(spaceStat *statistics.SpacesStat) {
	fmt.Println("NumRows:", spaceStat.NumRows)
	fmt.Println("NumRecords:", spaceStat.NumRecords)
	printSpaceSatFieldData("Nameu", spaceStat.Nameu)
}

func printSpaceSatFieldData(name string, fieldStat statistics.FieldSpaceStat) {
	prefTab := strings.Repeat(" ", 4)
	fmt.Printf("%s:\n", name)
	fmt.Printf("%sLeadingSpace:    %d\n", prefTab, fieldStat.LeadingSpace)
	fmt.Printf("%sTrailingSpace:   %d\n", prefTab, fieldStat.TrailingSpace)
	fmt.Printf("%sBorderingSpaces: %d\n", prefTab, fieldStat.BorderingSpaces)
}
