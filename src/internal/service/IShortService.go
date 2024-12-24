package service

import (
	"ShortService/src/common"
	"ShortService/src/common/e"
	"ShortService/src/global"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type IShortService interface {
	CreateShort(ctx context.Context, url string) (shortUrl string, err error)
	ParseUrl(ctx context.Context, shortUrl string) (url string, err error)
}

type ShortServiceImpl struct {
	sequenceService  ISequenceService
	shortUrlsService IShortUrlService
}

func NewShortService(sequenceService ISequenceService, shortUrlsService IShortUrlService) IShortService {
	return &ShortServiceImpl{sequenceService: sequenceService, shortUrlsService: shortUrlsService}
}

func (c *ShortServiceImpl) CreateShort(ctx context.Context, url string) (shortUrl string, err error) {
	//0、先查redis看是否存在,布隆过滤器?

	//1、首先判断当前链接是否存在,如果存在直接返回对应的短链接
	shortUrls, err := c.shortUrlsService.FindByLongUrl(url)
	if shortUrls.Id != 0 {
		shortUrl = global.EncoderBase62(shortUrls.Short)
		global.Log.Info("短链接已经存在,原始长链接:%s,shortUrl:%s", url, shortUrl)
		return shortUrl, err
	}
	//2、通过发号器获取编号
	//2.1 判断发号器是否已经初始化,如果Current=0,则没有初始化,同时还要判断号码是否已经发完,调用NextBatch()进行初始化
	commonSequenceDto := &common.CommonSequenceDto
	if commonSequenceDto.Current == 0 || commonSequenceDto.End < commonSequenceDto.Current {
		common.CommonSequenceDto, err = c.sequenceService.NextBatch(ctx, common.CommonSequenceDto)
	}
	//2.2 获取号码
	var code uint64 = 0
	for count := 0; count < 3; count++ {
		code, err = c.sequenceService.Next()
		if err == nil {
			break
		}
		if count == 2 {
			return "", fmt.Errorf("创建失败")
		}
	}

	//3、向数据库中写入长链接以及对应的编号
	shortUrls.Short = code
	shortUrls.LongUrl = url
	err = c.shortUrlsService.Save(shortUrls)
	if err != nil {
		global.Log.Error("添加到数据库失败")
		return shortUrl, err
	}

	//4、对编号进行base62编码,并返回
	shortUrl = global.EncoderBase62(code)

	global.Log.Info("长链接:%s, 短链接:%s", url, shortUrl)
	return shortUrl, err
}

func (c *ShortServiceImpl) ParseUrl(ctx context.Context, shortUrl string) (url string, err error) {
	//1、布隆过滤器
	if !global.Bloom.Test([]byte(shortUrl)) {
		return "", fmt.Errorf("short code %s not found (Bloom filter rejection)", shortUrl)
	}

	//2、 查询redis
	cc := context.Background()
	get, rerr := global.Redis.Get(cc, shortUrl).Result()
	if rerr == nil && rerr != redis.Nil {
		//命中
		global.Log.Info(fmt.Sprintf("缓存命中,shortUrl:%s, longUrl:%s", shortUrl, get))
		return get, nil
	}
	global.Log.Info(fmt.Sprintf("缓存未命中,shortUrl:%s, longUrl:%s", shortUrl, get))

	//3、查询数据库
	shortcode := global.DecoderBase62(shortUrl)
	shortUrls, err := c.shortUrlsService.FindByShortCode(shortcode)
	if err != nil || shortUrls.Short != shortcode {
		global.Log.Info("该链接不存在")
		return "", fmt.Errorf("无效链接")
	}

	// 4、redis中不存在,写入redis
	if rerr == redis.Nil {
		err := global.Redis.Set(cc, shortUrl, shortUrls.LongUrl, e.REDIS_DURATION*time.Second).Err()
		if err != nil {
			global.Log.Error(fmt.Sprintf("shortUrl:%s, shortCode:%d, longUrl:%s,写入redis失败", shortUrl, shortcode, shortUrls.LongUrl))
		} else {
			global.Log.Info(fmt.Sprintf("shortUrl:%s, shortCode:%d, longUrl:%s,写入redis成功", shortUrl, shortcode, shortUrls.LongUrl))
		}

	}
	return shortUrls.LongUrl, err
}
