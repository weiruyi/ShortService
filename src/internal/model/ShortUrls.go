package model

import (
	"gorm.io/gorm"
	"time"
)

type ShortUrls struct {
	Id         uint64     `json:"id"`
	CreateTime *time.Time `json:"createTime"`
	UpdateTime *time.Time `json:"updateTime"`
	DeleteTime *time.Time `json:"deleteTime" gorm:"default:null"`
	LongUrl    string     `json:"longUrl"`
	Short      uint64     `json:"short"`
}

func (c *ShortUrls) BeforeCreate(tx *gorm.DB) error {
	// 自动填充 创建时间、创建人、更新时间、更新用户
	now := time.Now()
	c.CreateTime = &now
	c.UpdateTime = &now
	// 从上下文获取用户信息
	//value := tx.Statement.Context.Value(enum.CurrentId)
	//if uid, ok := value.(uint64); ok {
	//	c.CreateUser = uid
	//	c.UpdateUser = uid
	//}
	return nil
}

func (c *ShortUrls) BeforeUpdate(tx *gorm.DB) error {
	// 在更新记录千自动填充更新时间
	now := time.Now()
	c.UpdateTime = &now
	//// 从上下文获取用户信息
	//value := tx.Statement.Context.Value(enum.CurrentId)
	//if uid, ok := value.(uint64); ok {
	//	c.UpdateUser = uid
	//}
	return nil
}
