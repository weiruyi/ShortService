package common

import "ShortService/src/internal/model"

var (
	CommonSequenceDto = model.SequenceDto{
		BatchSize: SequenceBatchSize,
		Current:   0,
		End:       0,
		Sequence: model.Sequence{
			Name: SequenceName,
		},
	}
)
