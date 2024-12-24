package initialize

import (
	"ShortService/src/global"
	"ShortService/src/internal/repository/dao"
	"ShortService/src/internal/service"
	"context"
	"fmt"
	"github.com/bits-and-blooms/bloom"
	"github.com/redis/go-redis/v9"
)

func initRedis() *redis.Client {
	redisOpt := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisOpt.Host, redisOpt.Port),
		Password: redisOpt.Password, // no password set
		DB:       redisOpt.DataBase, // use default DB
	})
	ping := client.Ping(context.Background())
	err := ping.Err()
	if err != nil {
		panic(err)
	}
	return client
}

func initBloom() *bloom.BloomFilter {
	filter := bloom.NewWithEstimates(1000000, 0.01)
	shortUrlService := service.NewShortUrlService(dao.NewShortUrlsDao(global.DB))
	shortUrls, err := shortUrlService.List(context.Background())
	if err != nil {
		global.Log.Error("init bloom short url err:", err.Error())
	}
	for _, shortUrl := range shortUrls {
		shortCode := shortUrl.Short
		shortUrl := global.EncoderBase62(shortCode)
		filter.Add([]byte(shortUrl))
	}
	global.Log.Info("init bloom short success")
	return filter
}
