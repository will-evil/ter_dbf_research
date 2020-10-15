// Package statistics provide functional for collect different statistics about dbf files
// which contain data about terrorists.
package statistics

import (
	"errors"
	"strconv"

	"ter_dbf_research/src/godbf"
)

const sepFieldMaxLength = 254

// Struct for store info about row from dbf terrorist file.
type rowData struct {
	index int
	rowID uint64
}

type rowDataMap map[uint64][]rowData

// Stat is structure for store files list and provide functional for collect statistics.
type Stat struct {
	dbfTable   *godbf.DbfTable
	rowDataMap *rowDataMap
}

// NewStat is the constructor for Stat structure.
func NewStat(filePath string) (*Stat, error) {
	if filePath == "" {
		return nil, errors.New("filePath can not be empty")
	}

	dbfTable, err := godbf.NewFromFile(filePath, "866")
	if err != nil {
		return nil, err
	}

	rowDataMap := make(rowDataMap)

	return &Stat{
		dbfTable:   dbfTable,
		rowDataMap: &rowDataMap,
	}, nil
}

// Spaces returns statistics about spaces in fields which may be separeted.
func (s *Stat) Spaces() (*SpacesStat, error) {
	rowDataMap, err := s.getRowDataMap()
	if err != nil {
		return nil, err
	}
	ss := NewSpaceStat(s.dbfTable, rowDataMap)
	if err := ss.StatsForAll(); err != nil {
		return nil, err
	}

	return ss, nil
}

// Dup counts duplicates in separated fields.
func (s *Stat) Dup() (*DupStat, error) {
	rowDataMap, err := s.getRowDataMap()
	if err != nil {
		return nil, err
	}
	ds := NewDupStat(s.dbfTable, rowDataMap)
	if err := ds.StatsForAll(); err != nil {
		return nil, err
	}

	return ds, nil
}

func (s *Stat) getRowDataMap() (*rowDataMap, error) {
	var err error

	if len(*s.rowDataMap) == 0 {
		err = s.setRowDataMap()
	}

	return s.rowDataMap, err
}

// Set map with data about records from dbf file.
func (s *Stat) setRowDataMap() error {
	rowDataMap := make(rowDataMap)

	for i := 0; i < s.dbfTable.NumberOfRecords(); i++ {
		numberStr, err := s.dbfTable.FieldValueByName(i, "NUMBER")
		if err != nil {
			return err
		}
		rowIDStr, err := s.dbfTable.FieldValueByName(i, "ROW_ID")
		if err != nil {
			return err
		}

		number, err := strconv.ParseUint(numberStr, 10, 64)
		if err != nil {
			return err
		}
		rowID, err := strconv.ParseUint(rowIDStr, 10, 64)
		if err != nil {
			return err
		}

		data := rowData{index: i, rowID: rowID}
		rowDataMap[number] = append(rowDataMap[number], data)
	}

	s.rowDataMap = &rowDataMap

	return nil
}
