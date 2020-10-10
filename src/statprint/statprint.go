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
	printSpaceSatFieldData("Descript", spaceStat.Descript)
	printSpaceSatFieldData("Amr", spaceStat.Amr)
	printSpaceSatFieldData("Address", spaceStat.Address)
	printSpaceSatFieldData("Mr", spaceStat.Mr)
	printSpaceSatFieldData("Director", spaceStat.Director)
	printSpaceSatFieldData("Founder", spaceStat.Founder)
	printSpaceSatFieldData("Terrtype", spaceStat.Terrtype)
}

func printSpaceSatFieldData(name string, fieldStat statistics.FieldSpaceStat) {
	prefTab := strings.Repeat(" ", 4)
	fmt.Printf("%s:\n", name)
	fmt.Printf("%sLeadingSpace:    %d\n", prefTab, fieldStat.LeadingSpace)
	fmt.Printf("%sTrailingSpace:   %d\n", prefTab, fieldStat.TrailingSpace)
	fmt.Printf("%sBorderingSpaces: %d\n", prefTab, fieldStat.BorderingSpaces)
}
