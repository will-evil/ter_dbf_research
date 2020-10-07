// Package statistics provide functional for collect different statistics about dbf files
// which contain data about terrorists.
package statistics

import (
	"errors"
	"strconv"

	"github.com/LindsayBradford/go-dbf/godbf"
)

// Struct for store info about row from dbf terrorist file.
// This struct stores row index in file and value of column ROW-ID.
type rowData struct {
	index int
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
	dbfTable.NumberOfRecords()
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
		number, err := strconv.ParseUint(numberStr, 10, 64)
		if err != nil {
			return err
		}

		data := rowData{index: i}
		rowDataMap[number] = append(rowDataMap[number], data)
	}

	s.rowDataMap = &rowDataMap

	return nil
}
