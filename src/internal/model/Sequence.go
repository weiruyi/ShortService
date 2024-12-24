package model

import (
	"gorm.io/gorm"
	"time"
)

type Sequence struct {
	Id           uint64    `json:"id"`
	Name         string    `json:"name"`
	CurrentValue uint64    `json:"currentValue"`
	Increment    uint64    `json:"increment"`
	CreateTime   time.Time `json:"createTime"`
	UpdateTime   time.Time `json:"updateTime"`
	DeleteTime   time.Time `json:"deleteTime"`
}

func (c *Sequence) BeforeCreate(tx *gorm.DB) error {
	// 自动填充 创建时间、创建人、更新时间、更新用户
	c.CreateTime = time.Now()
	c.UpdateTime = time.Now()
	// 从上下文获取用户信息
	//value := tx.Statement.Context.Value(enum.CurrentId)
	//if uid, ok := value.(uint64); ok {
	//	c.CreateUser = uid
	//	c.UpdateUser = uid
	//}
	return nil
}

func (c *Sequence) BeforeUpdate(tx *gorm.DB) error {
	// 在更新记录千自动填充更新时间
	c.UpdateTime = time.Now()
	//// 从上下文获取用户信息
	//value := tx.Statement.Context.Value(enum.CurrentId)
	//if uid, ok := value.(uint64); ok {
	//	c.UpdateUser = uid
	//}
	return nil
}
