package model

import (
	"GuTikTok/utils"

	"gorm.io/gorm"
)

type (
	// Video 视频表
	Video struct {
		Model
		AuthorID      int64      `json:"-" gorm:"index;notNull;comment:视频作者信息"`
		Author        User       `json:"author"`
		PlayUrl       string     `json:"play_url" gorm:"comment:视频播放地址"`
		CoverUrl      string     `json:"cover_url" gorm:"comment:视频封面地址"`
		Title         string     `json:"title" gorm:"comment:视频标题"`
		Desc          string     `json:"desc" gorm:"comment:简介"`
		Comment       []*Comment `json:"comment,omitempty" gorm:"comment:评论列表"`
		FavoriteUser  []*User    `json:"-" gorm:"many2many:UserFavorite;comment:喜欢用户列表"`
		IsFavorite    bool       `json:"is_favorite" gorm:"-"`    // 是否点赞
		PlayCount     int64      `json:"play_count" gorm:"-"`     // 视频播放量
		FavoriteCount int64      `json:"favorite_count" gorm:"-"` // 视频的点赞总数
		CommentCount  int64      `json:"comment_count" gorm:"-"`  // 视频的评论总数
		// 自建字段
		TypeOf   string  `json:"typeOf" gorm:"comment:视频类型"`
		CoAuthor []*User `json:"authors,omitempty" gorm:"many2many:UserCreation;"` // 联合投稿
	}
)

func (v *Video) AfterFind(tx *gorm.DB) (err error) {

	return
}

func (v *Video) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == 0 {
		v.ID = utils.GetId(3, 20230724)
	}
	return
}

func init() {
	addMigrate(&Video{})

}
