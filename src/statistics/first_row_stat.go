package statistics

import (
	"sort"
	"ter_dbf_research/src/godbf"
)

// FirstRowStat is structure for store statistics about correct and not correct values in first row of record.
type FirstRowStat struct {
	CommonStat
	NotCorrect       uint64
	NotCorrectTerror uint64
	NotCorrectTu     uint64
	NotCorrectKd     uint64
	dbfTable         *godbf.DbfTable
	rowDataMap       *rowDataMap
}

// NewFirstRowStat is the constructor for FirstRowStat structure.
func NewFirstRowStat(dbfTable *godbf.DbfTable, rowDataMap *rowDataMap) *FirstRowStat {
	frs := &FirstRowStat{dbfTable: dbfTable, rowDataMap: rowDataMap}
	frs.setNumFields(dbfTable, rowDataMap)

	return frs
}

// StatsForAll calculates statistics for all records
func (frs *FirstRowStat) StatsForAll() error {
	var rowIsCorrect bool

	checkFieldsData := []struct {
		name          string
		allowedValues []string
		counter       *uint64
	}{
		{"TERROR", []string{"0", "1"}, &frs.NotCorrectTerror},
		{"TU", []string{"1", "2", "3"}, &frs.NotCorrectTu},
		{"KD", []string{"0", "01", "02", "03", "04"}, &frs.NotCorrectKd},
	}

	for _, rowDataSlice := range *frs.rowDataMap {
		sort.SliceStable(rowDataSlice, func(i, j int) bool {
			return rowDataSlice[i].rowID < rowDataSlice[j].rowID
		})

		rowIsCorrect = true
		for _, field := range checkFieldsData {
			res, err := frs.checkField(rowDataSlice[0].index, field.name, field.allowedValues)
			if err != nil {
				return err
			}
			if !res {
				rowIsCorrect = false
				*field.counter++
			}
		}

		if !rowIsCorrect {
			frs.NotCorrect++
		}
	}

	return nil
}

func (frs *FirstRowStat) checkField(index int, name string, allowedValues []string) (bool, error) {
	val, err := frs.dbfTable.FieldValueByName(index, name)
	if err != nil {
		return false, err
	}

	for _, allowed := range allowedValues {
		if val == allowed {
			return true, nil
		}
	}

	return false, nil
}
