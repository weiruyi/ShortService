package dao

import (
	"ShortService/src/global"
	"ShortService/src/internal/model"
	"ShortService/src/internal/repository"
	"context"
	"gorm.io/gorm"
	"time"
)

type ShortUrlsDao struct {
	db *gorm.DB
}

func NewShortUrlsDao(db *gorm.DB) repository.ShortUrlsRepo {
	return &ShortUrlsDao{db: db}
}

// 根据长链接来查询
func (c *ShortUrlsDao) FindByLongUrl(longUrl string) (model.ShortUrls, error) {
	shortUrls := model.ShortUrls{}
	err := c.db.Where("long_url = ?", longUrl).Find(&shortUrls).Error
	if err != nil {
		global.Log.Info("链接不存在")
	}
	return shortUrls, err
}

// 插入
func (c *ShortUrlsDao) Save(shortUrl model.ShortUrls) error {
	err := c.db.Create(&shortUrl).Error
	return err
}

// 根据短链接查询
func (c *ShortUrlsDao) FindByShortCode(shortCode uint64) (model.ShortUrls, error) {
	shortUrls := model.ShortUrls{}
	err := c.db.Where("short = ?", shortCode).Find(&shortUrls).Error
	return shortUrls, err
}

// 查询全部未删除的
func (c *ShortUrlsDao) List(ctx context.Context) ([]model.ShortUrls, error) {
	var shortUrls []model.ShortUrls
	err := c.db.Where("delete_time IS NULL OR delete_time > ?", time.Now()).Find(&shortUrls).Error
	return shortUrls, err
}
