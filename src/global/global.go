package global

import (
	"ShortService/src/config"
	"ShortService/src/logger"
	"github.com/bits-and-blooms/bloom"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config *config.AllConfig
	DB     *gorm.DB
	Log    logger.ILog
	Redis  *redis.Client
	Bloom  *bloom.BloomFilter
)
