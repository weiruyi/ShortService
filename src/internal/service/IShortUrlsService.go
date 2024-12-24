package service

import (
	"ShortService/src/internal/model"
	"ShortService/src/internal/repository"
	"context"
)

type IShortUrlService interface {
	FindByLongUrl(url string) (model.ShortUrls, error)
	Save(shortUrls model.ShortUrls) error
	FindByShortCode(code uint64) (model.ShortUrls, error)
	List(ctx context.Context) ([]model.ShortUrls, error)
}

type ShortUrlServiceImpl struct {
	shortUrlsDao repository.ShortUrlsRepo
}

func NewShortUrlService(shortUrlsDao repository.ShortUrlsRepo) IShortUrlService {
	return &ShortUrlServiceImpl{shortUrlsDao: shortUrlsDao}
}

func (s *ShortUrlServiceImpl) FindByLongUrl(url string) (model.ShortUrls, error) {
	return s.shortUrlsDao.FindByLongUrl(url)
}

func (c *ShortUrlServiceImpl) FindByShortCode(code uint64) (model.ShortUrls, error) {
	return c.shortUrlsDao.FindByShortCode(code)
}

func (s *ShortUrlServiceImpl) Save(shortUrls model.ShortUrls) error {
	return s.shortUrlsDao.Save(shortUrls)
}

func (c *ShortUrlServiceImpl) List(ctx context.Context) ([]model.ShortUrls, error) {
	return c.shortUrlsDao.List(ctx)
}
