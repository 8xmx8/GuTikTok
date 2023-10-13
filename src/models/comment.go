package models

import (
	"GuTikTok/src/storage/database"
	"GuTikTok/utils"
	"gorm.io/gorm"
	"time"
)

type (
	// Comment 评论表
	Comment struct {
		Model
		UserID     int64  `json:"-" gorm:"index:idx_uvid;comment:评论用户信息"`
		VideoID    int64  `json:"-" gorm:"index:idx_uvid;comment:评论视频信息"`
		User       User   `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Video      Video  `json:"video" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Content    string `json:"content" gorm:"comment:评论内容"`
		CreateDate string `json:"create_date" gorm:"comment:评论发布日期"` // 格式 mm-dd
		// 自建字段
		ReplyID int64 `json:"reply_id" gorm:"index;comment:回复ID"`
	}
)

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == 0 {
		c.ID = utils.GetId(2, 20230724)
	}
	c.CreateDate = time.Now().Format("01-02")
	return
}

func init() {
	if err := database.Client.AutoMigrate(&Comment{}); err != nil {
		panic(err)
	}
}
