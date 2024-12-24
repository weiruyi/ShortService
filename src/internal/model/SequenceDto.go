package model

import "sync"

type SequenceDto struct {
	Sequence  Sequence
	BatchSize uint64
	Current   uint64
	End       uint64
	SegmentMu sync.Mutex
}

//func NewSequenceDto(name string, batchSize uint64) *SequenceDto {
//	return &SequenceDto{Sequence : Sequence{Name: name}}
//}
