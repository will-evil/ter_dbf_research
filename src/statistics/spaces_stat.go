package statistics

import (
	"sort"
	"ter_dbf_research/src/godbf"
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
	for _, rowDataSlice := range *ss.rowDataMap {
		sort.SliceStable(rowDataSlice, func(i, j int) bool {
			return rowDataSlice[i].rowID < rowDataSlice[j].rowID
		})

		sliceSize := len(rowDataSlice)
		for i := 0; i < sliceSize; i++ {
			rowData := rowDataSlice[i]
			err := ss.countSpacesForRow(
				rowData.index,
				nextRowIndex(i, rowDataSlice),
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (ss *SpacesStat) countSpacesForRow(rowIndex int, nextRowIndex *int) error {
	var nextStrPointer *string

	fieldsSet := []struct {
		name      string
		fieldStat *FieldSpaceStat
	}{
		{"NAMEU", &ss.Nameu},
		{"DESCRIPT", &ss.Descript},
		{"AMR", &ss.Amr},
		{"ADRESS", &ss.Address},
		{"MR", &ss.Mr},
		{"DIRECTOR", &ss.Director},
		{"FOUNDER", &ss.Founder},
		{"TERRTYPE", &ss.Terrtype},
	}

	for _, set := range fieldsSet {
		value, err := ss.dbfTable.FieldValueByName(rowIndex, set.name)
		if err != nil {
			return err
		}

		if nextRowIndex != nil {
			nextStr, err := ss.dbfTable.FieldValueByName(*nextRowIndex, set.name)
			if err != nil {
				return err
			}
			nextStrPointer = &nextStr
		} else {
			nextStrPointer = nil
		}

		countSpaces(value, nextStrPointer, set.fieldStat)
	}

	return nil
}

func nextRowIndex(currentIndex int, rowDataSlice []rowData) *int {
	nextIndex := currentIndex + 1
	if nextIndex < len(rowDataSlice) {
		return &rowDataSlice[nextIndex].index
	}

	return nil
}

func countSpaces(str string, nextStr *string, fieldStat *FieldSpaceStat) {
	if len(str) == 0 {
		return
	}

	var res uint8

	if str[0] == ' ' {
		res++
	}

	if nextStr != nil && len(*nextStr) > 0 && len(str) == (sepFieldMaxLength-1) {
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
