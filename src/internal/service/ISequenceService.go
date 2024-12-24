package service

import (
	"ShortService/src/common"
	"ShortService/src/internal/model"
	"ShortService/src/internal/repository"
	"context"
	"fmt"
	"sync/atomic"
)

type ISequenceService interface {
	//初始化或者获取新的发号区间
	NextBatch(ctx context.Context, sequenceDto model.SequenceDto) (resSequenceDto model.SequenceDto, err error)
	//下一个号码
	Next() (uint64, error)
}

type SequenceServiceImpl struct {
	sequenceRepo repository.SequenceRepo
}

func NewSequenceService(sequenceRepo repository.SequenceRepo) ISequenceService {
	return &SequenceServiceImpl{sequenceRepo: sequenceRepo}
}

// 给发号器分配新的区间
func (c *SequenceServiceImpl) NextBatch(ctx context.Context, sequenceDto model.SequenceDto) (resSequenceDto model.SequenceDto, err error) {
	common.CommonSequenceDto.SegmentMu.Lock()
	defer common.CommonSequenceDto.SegmentMu.Unlock()
	if common.CommonSequenceDto.Current != 0 && common.CommonSequenceDto.Current <= common.CommonSequenceDto.End {
		return resSequenceDto, nil
	}
	resSequenceDto, err = c.sequenceRepo.QueryRow(ctx, sequenceDto)
	return resSequenceDto, err
}

// 获取号码
func (c *SequenceServiceImpl) Next() (uint64, error) {
	//code := common.CommonSequenceDto.Current
	//common.CommonSequenceDto.Current++
	code := atomic.AddUint64(&common.CommonSequenceDto.Current, 1)
	code--
	if code > common.CommonSequenceDto.End {
		return 0, fmt.Errorf("发号器出错")
	}
	return code, nil
}
