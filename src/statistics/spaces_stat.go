package statistics

import (
	"github.com/LindsayBradford/go-dbf/godbf"
)

// FieldSpaceStat is structure for store field statictic.
type FieldSpaceStat struct {
	LeadingSpace    uint64
	TrailingSpace   uint64
	BorderingSpaces uint64
}

// SpacesStat is struct for store spaces statictics about field which may be separeted.
type SpacesStat struct {
	CommonStat
	Nameu      FieldSpaceStat
	Descript   FieldSpaceStat
	Amr        FieldSpaceStat
	Address    FieldSpaceStat
	Mr         FieldSpaceStat
	Director   FieldSpaceStat
	Founder    FieldSpaceStat
	Terrtype   FieldSpaceStat
	dbfTable   *godbf.DbfTable
	rowDataMap *rowDataMap
}

// NewSpaceStat is the constructor for SpacesStat structure.
func NewSpaceStat(dbfTable *godbf.DbfTable, rowDataMap *rowDataMap) *SpacesStat {
	ss := &SpacesStat{dbfTable: dbfTable, rowDataMap: rowDataMap}
	ss.setNumFields(dbfTable, rowDataMap)

	return ss
}

// StatsForAll calculates statistics for all records
func (ss *SpacesStat) StatsForAll() error {
	for i := 0; i < ss.dbfTable.NumberOfRecords(); i++ {
		nameu, err := ss.dbfTable.FieldValueByName(i, "NAMEU")
		if err != nil {
			return err
		}
		countSpaces(nameu, &ss.Nameu)

		descript, err := ss.dbfTable.FieldValueByName(i, "DESCRIPT")
		if err != nil {
			return err
		}
		countSpaces(descript, &ss.Descript)

		amr, err := ss.dbfTable.FieldValueByName(i, "AMR")
		if err != nil {
			return err
		}
		countSpaces(amr, &ss.Amr)

		address, err := ss.dbfTable.FieldValueByName(i, "ADRESS")
		if err != nil {
			return err
		}
		countSpaces(address, &ss.Address)

		mr, err := ss.dbfTable.FieldValueByName(i, "MR")
		if err != nil {
			return err
		}
		countSpaces(mr, &ss.Mr)

		director, err := ss.dbfTable.FieldValueByName(i, "DIRECTOR")
		if err != nil {
			return err
		}
		countSpaces(director, &ss.Director)

		founder, err := ss.dbfTable.FieldValueByName(i, "FOUNDER")
		if err != nil {
			return err
		}
		countSpaces(founder, &ss.Founder)

		terrtype, err := ss.dbfTable.FieldValueByName(i, "TERRTYPE")
		if err != nil {
			return err
		}
		countSpaces(terrtype, &ss.Terrtype)
	}

	return nil
}

func countSpaces(str string, fieldStat *FieldSpaceStat) {
	if len(str) == 0 {
		return
	}

	var res uint8

	if str[0] == ' ' {
		res++
	}
	if str[len(str)-1:] == " " {
		res += 2
	}

	switch res {
	case 1:
		fieldStat.LeadingSpace++
	case 2:
		fieldStat.TrailingSpace++
	case 3:
		fieldStat.BorderingSpaces++
	}
}
