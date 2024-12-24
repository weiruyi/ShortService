package repository

import (
	"ShortService/src/internal/model"
	"context"
)

type ShortUrlsRepo interface {
	FindByLongUrl(longUrl string) (model.ShortUrls, error)
	Save(shortUrl model.ShortUrls) error
	FindByShortCode(shortCode uint64) (model.ShortUrls, error)
	List(ctx context.Context) ([]model.ShortUrls, error)
}
