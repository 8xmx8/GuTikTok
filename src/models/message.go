package models

import (
	"GuTikTok/src/storage/database"
	"GuTikTok/utils"
	"gorm.io/gorm"
)

type (
	// Message 消息表
	Message struct {
		ID         int64  `json:"id" gorm:"primarykey;comment:主键"`
		CreatedAt  int64  `json:"create_time" gorm:"autoUpdateTime:milli"`
		ToUserID   int64  `json:"to_user_id" gorm:"primaryKey;comment:该消息接收者的id"`
		FromUserID int64  `json:"from_user_id" gorm:"primaryKey;comment:该消息发送者的id"`
		ToUser     User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		FromUser   User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Content    string `json:"content" gorm:"comment:消息内容"`
	}
)

func (m *Message) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == 0 {
		m.ID = utils.GetId(3, 114514)
	}
	// 来自一个天坑
	// m.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return
}

func init() {
	if err := database.Client.AutoMigrate(&Message{}); err != nil {
		panic(err)
	}

}
