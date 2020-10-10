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

// DupStat prints statistics about duplications in fields.
func DupStat(dupStat *statistics.DupStat, fullPrint bool) {
	fmt.Println("NumRows:", dupStat.NumRows)
	fmt.Println("NumRecords:", dupStat.NumRecords)
	printFieldDupStat("Nameu", dupStat.Nameu, fullPrint)
	printFieldDupStat("Descript", dupStat.Descript, fullPrint)
	printFieldDupStat("Amr", dupStat.Amr, fullPrint)
	printFieldDupStat("Address", dupStat.Address, fullPrint)
	printFieldDupStat("Mr", dupStat.Mr, fullPrint)
	printFieldDupStat("Director", dupStat.Director, fullPrint)
	printFieldDupStat("Founder", dupStat.Founder, fullPrint)
	printFieldDupStat("Terrtype", dupStat.Terrtype, fullPrint)
}

func printSpaceSatFieldData(name string, fieldStat statistics.FieldSpaceStat) {
	prefTab := strings.Repeat(" ", 4)
	fmt.Printf("%s:\n", name)
	fmt.Printf("%sLeadingSpace:    %d\n", prefTab, fieldStat.LeadingSpace)
	fmt.Printf("%sTrailingSpace:   %d\n", prefTab, fieldStat.TrailingSpace)
	fmt.Printf("%sBorderingSpaces: %d\n", prefTab, fieldStat.BorderingSpaces)
}

func printFieldDupStat(name string, fieldStat statistics.FieldDupStat, fullPrint bool) {
	prefTab := strings.Repeat(" ", 4)
	fmt.Printf("%s:\n", name)
	fmt.Printf("%sCount: %d\n", prefTab, fieldStat.Count)

	if !fullPrint {
		return
	}

	fmt.Printf("%sCountByNumbers:\n", prefTab)
	prefTab = strings.Repeat(prefTab, 2)
	for number, count := range fieldStat.NumberCountMap {
		fmt.Printf("%s%d: %d\n", prefTab, number, count)
	}
}
