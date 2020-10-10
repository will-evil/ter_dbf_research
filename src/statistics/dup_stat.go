package statistics

import (
	"fmt"
	"strconv"

	"github.com/LindsayBradford/go-dbf/godbf"
)

// FieldDupStat is structure for store duplication statistic for concret field.
type FieldDupStat struct {
	Count          uint64
	NumberCountMap map[uint64]uint64
}

// DupStat is structure for store statistics about duplicates in fields for one record.
type DupStat struct {
	CommonStat
	Nameu      FieldDupStat
	Descript   FieldDupStat
	Amr        FieldDupStat
	Address    FieldDupStat
	Mr         FieldDupStat
	Director   FieldDupStat
	Founder    FieldDupStat
	Terrtype   FieldDupStat
	dbfTable   *godbf.DbfTable
	rowDataMap *rowDataMap
}

// NewDupStat is the constructor for DupStat structure.
func NewDupStat(dbfTable *godbf.DbfTable, rowDataMap *rowDataMap) *DupStat {
	ds := &DupStat{dbfTable: dbfTable, rowDataMap: rowDataMap}
	ds.setNumFields(dbfTable, rowDataMap)

	return ds
}

// StatsForAll calculates statistics for all records
func (ds *DupStat) StatsForAll() error {
	for i := 0; i < ds.dbfTable.NumberOfRecords(); i++ {
		numberStr, err := ds.dbfTable.FieldValueByName(i, "NUMBER")
		if err != nil {
			return err
		}
		number, err := strconv.ParseUint(numberStr, 10, 64)
		if err != nil {
			return err
		}

		if err := ds.processRowForDup(number); err != nil {
			return nil
		}
	}

	return nil
}

func (ds *DupStat) processRowForDup(number uint64) error {
	rowPartsData := (*ds.rowDataMap)[number]
	if len(rowPartsData) < 2 {
		return nil
	}

	fields := []string{
		"NAMEU",
		"DESCRIPT",
		"AMR",
		"ADRESS",
		"MR",
		"DIRECTOR",
		"FOUNDER",
		"TERRTYPE",
	}

	rowFieldMap := make(map[string][]string)
	for _, rowPartData := range rowPartsData {
		for _, field := range fields {
			value, err := ds.dbfTable.FieldValueByName(rowPartData.index, field)
			if err != nil {
				return err
			}
			if value == "" {
				continue
			}
			if existsInStringSlice(rowFieldMap[field], value) {
				if err := ds.addDupForField(field, number); err != nil {
					return err
				}
			}
			rowFieldMap[field] = append(rowFieldMap[field], value)
		}
	}

	return nil
}

func (ds *DupStat) addDupForField(field string, number uint64) error {
	incrementStat := func(dupStat *FieldDupStat) {
		dupStat.Count++
		if len(dupStat.NumberCountMap) == 0 {
			dupStat.NumberCountMap = make(map[uint64]uint64)
		}
		dupStat.NumberCountMap[number]++
	}

	switch field {
	case "NAMEU":
		incrementStat(&ds.Nameu)
	case "DESCRIPT":
		incrementStat(&ds.Descript)
	case "AMR":
		incrementStat(&ds.Amr)
	case "ADRESS":
		incrementStat(&ds.Address)
	case "MR":
		incrementStat(&ds.Mr)
	case "DIRECTOR":
		incrementStat(&ds.Director)
	case "FOUNDER":
		incrementStat(&ds.Founder)
	case "TERRTYPE":
		incrementStat(&ds.Terrtype)
	default:
		return fmt.Errorf("field with name '%s' not supported", field)
	}

	return nil
}

func existsInStringSlice(items []string, target string) bool {
	for _, val := range items {
		if val == target {
			return true
		}
	}
	return false
}
