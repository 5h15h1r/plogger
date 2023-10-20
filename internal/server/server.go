package server

import (
	"fmt"
	"sync"
)

type Record struct {
	Value		[]byte `json:"value"`
	Offset		uint64 `json:"offset"`
}

type Log struct{
	mut   		sync.Mutex
	records 	[]Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error){
	c.mut.Lock()
	defer c.mut.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
	c.mut.Lock()
	defer c.mut.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
		}
		return c.records[offset], nil
}
var ErrOffsetNotFound = fmt.Errorf("Offset not found")