package statistics

import "ter_dbf_research/src/godbf"

// CommonStat is struct for store common file statistics data.
type CommonStat struct {
	NumRows    uint64
	NumRecords uint64
}

func (cs *CommonStat) setNumFields(dbfTable *godbf.DbfTable, rowDataMap *rowDataMap) {
	cs.NumRows = uint64(dbfTable.NumberOfRecords())
	cs.NumRecords = uint64(len(*rowDataMap))
}
