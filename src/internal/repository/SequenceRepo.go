package repository

import (
	"ShortService/src/internal/model"
	"context"
)

type SequenceRepo interface {
	QueryRow(ctx context.Context, sequenceDto model.SequenceDto) (model.SequenceDto, error)
}
