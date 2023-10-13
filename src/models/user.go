package models

import (
	"GuTikTok/src/storage/database"
	"GuTikTok/utils"
	"fmt"
	"net/url"
	"regexp"
	"time"

	"gorm.io/gorm"
)

type (
	// User 用户信息表
	User struct {
		Model
		Name            string     `json:"name" gorm:"index:,unique;size:32;comment:用户名称"`
		Pawd            string     `json:"-" gorm:"size:128;comment:用户密码"`
		Avatar          string     `json:"avatar" gorm:"comment:用户头像"`
		BackgroundImage string     `json:"background_image" gorm:"comment:用户个人页顶部大图"`
		Signature       string     `json:"signature" gorm:"default:此人巨懒;comment:个人简介"`
		WorkCount       int64      `json:"work_count" gorm:"default:0;comment:作品数量"`
		Follow          []*User    `json:"follow,omitempty" gorm:"many2many:UserFollow;comment:关注列表"`
		Favorite        []*Video   `json:"like_list,omitempty" gorm:"many2many:UserFavorite;comment:喜欢列表"`
		Videos          []*Video   `json:"video_list,omitempty" gorm:"many2many:UserCreation;comment:作品列表"`
		Comment         []*Comment `json:"comment_list,omitempty" gorm:"comment:评论列表"`
		FollowCount     int64      `json:"follow_count" gorm:"-"`       // 关注总数
		FollowerCount   int64      `json:"follower_count" gorm:"-"`     // 粉丝总数
		TotalFavorited  int64      `json:"total_favorited" gorm:"-"`    // 获赞数量
		FavoriteCount   int64      `json:"favorite_count" gorm:"-"`     // 点赞数量
		IsFollow        bool       `json:"is_follow" gorm:"-"`          // 是否关注
		Follower        []*User    `json:"follower,omitempty" gorm:"-"` // 粉丝列表
		Friend          []*User    `json:"friend,omitempty" gorm:"-"`   // 好友列表
	}
	// FriendUser 好友结构体
	FriendUser struct {
		User
		Message string `json:"message"`  // 和该好友的最新聊天消息
		MsgType int    `json:"msg_type"` // 0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
	}
	// UserCreation 联合作者
	UserCreation struct {
		VideoID   int64          `json:"video_id,omitempty" gorm:"primaryKey"`
		UserID    int64          `json:"author_id" gorm:"primaryKey"`
		Type      string         `json:"type" gorm:"comment:创作者类型"` // Up主, 参演，剪辑，录像，道具，编剧，打酱油
		CreatedAt time.Time      `json:"created_at"`
		DeletedAt gorm.DeletedAt `json:"-"`
	}
)

var userCountKey = make([]byte, 0, 50)

// IsNameEmail 判断用户的名称是否为邮箱。
func (u *User) IsNameEmail() bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(u.Name)
}

func (u *User) IsDirty() bool {
	return u.Name != ""
}

func (u *User) GetID() int64 {
	return u.Model.ID
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.ID = utils.GetId(2, 114514)
	}
	if u.Avatar == "" {
		u.Avatar = fmt.Sprintf("https://api.multiavatar.com/%s.png", url.QueryEscape(u.Name))
	}
	if u.BackgroundImage == "" {
		u.BackgroundImage = "https://api.paugram.com/wallpaper/"
	}
	return
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	return
}

func init() {
	if err := database.Client.AutoMigrate(&User{}, &UserCreation{}); err != nil {
		panic(err)
	}
}
