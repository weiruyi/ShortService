package dao

import (
	"ShortService/src/global"
	"ShortService/src/internal/model"
	"ShortService/src/internal/repository"
	"context"
	"gorm.io/gorm"
)

type SequenceDao struct {
	db *gorm.DB
}

func (c *SequenceDao) QueryRow(ctx context.Context, sequenceDto model.SequenceDto) (model.SequenceDto, error) {
	var sequence model.Sequence
	tx := c.db.Begin()
	if tx.Error != nil {
		global.Log.Error("failed to start transaction")
	}

	err := c.db.WithContext(ctx).Where("name=?", sequenceDto.Sequence.Name).First(&sequence).Error
	if err != nil {
		global.Log.Error("error for check for db")
	}

	err = c.db.WithContext(ctx).Model(&sequence).Where("name=?", sequence.Name).Update("current_value", sequence.CurrentValue+sequenceDto.BatchSize).Error
	if err != nil {
		global.Log.Error("error for check for db")
	}

	sequenceDto.Current = sequence.CurrentValue - sequenceDto.BatchSize
	sequenceDto.Sequence = sequence
	sequenceDto.End = sequence.CurrentValue
	err = tx.Commit().Error
	if err != nil {
		global.Log.Error("failed to commit transaction")
	}
	return sequenceDto, err
}

func NewSequenceDao(db *gorm.DB) repository.SequenceRepo {
	return &SequenceDao{db: db}
}
