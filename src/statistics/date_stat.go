package statistics

import (
	"ter_dbf_research/src/godbf"
	"time"
)

// FieldDateStat is structure for store date format statistic for concret field.
type FieldDateStat struct {
	Formats map[string]uint64
	Other   uint64
	Empty   uint64
}

// DateStat is structure for store statistics about date formats.
type DateStat struct {
	CommonStat
	Gr         FieldDateStat
	CbDate     FieldDateStat
	CeDate     FieldDateStat
	layouts    []string
	dbfTable   *godbf.DbfTable
	rowDataMap *rowDataMap
}

// NewDateStat is the constructor for DateStat structure.
func NewDateStat(dbfTable *godbf.DbfTable, rowDataMap *rowDataMap) *DateStat {
	ds := &DateStat{dbfTable: dbfTable, rowDataMap: rowDataMap}
	ds.setNumFields(dbfTable, rowDataMap)

	ds.layouts = []string{
		"20060102",
		"02.01.06",
		"02.01.2006",
	}

	ds.initFormatsForField(&ds.Gr)
	ds.initFormatsForField(&ds.CbDate)
	ds.initFormatsForField(&ds.CeDate)

	return ds
}

// StatsForAll calculates statistics for all records
func (ds *DateStat) StatsForAll() error {
	statSet := []struct {
		fieldName string
		statStore *FieldDateStat
	}{
		{"GR", &ds.Gr},
		{"CB_DATE", &ds.CbDate},
		{"CE_DATE", &ds.CeDate},
	}

	for i := 0; i < ds.dbfTable.NumberOfRecords(); i++ {
		for _, set := range statSet {
			val, err := ds.dbfTable.FieldValueByName(i, set.fieldName)
			if err != nil {
				return err
			}

			if err := ds.setDateStat(val, set.statStore); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ds *DateStat) initFormatsForField(statStore *FieldDateStat) {
	statStore.Formats = make(map[string]uint64)
	for _, layout := range ds.layouts {
		statStore.Formats[layout] = 0
	}
}

func (ds *DateStat) setDateStat(dateStr string, statStore *FieldDateStat) error {
	if dateStr == "" {
		statStore.Empty++
		return nil
	}

	failedLayouts := 0
	for _, layout := range ds.layouts {
		if _, err := time.Parse(layout, dateStr); err != nil {
			failedLayouts++
			continue
		}

		statStore.Formats[layout]++
	}

	if failedLayouts == len(ds.layouts) {
		statStore.Other++
	}

	return nil
}
