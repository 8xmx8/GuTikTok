package model

import (
	"gorm.io/gorm"
	"log"
	"reflect"
)

// Init 初始化数据库服务
func Init(db *gorm.DB) {
	for _, m := range GetMigrate() {
		err := db.AutoMigrate(m)
		if err != nil {
			log.Fatalf("%s 模型自动迁移失败: %s", reflect.TypeOf(m), err.Error())
		}
	}
	err := db.SetupJoinTable(&Video{}, "CoAuthor", &UserCreation{})
	if err != nil {
		log.Fatalf("自定义连接表设置失败,Video: %s", err)
	}
	err = db.SetupJoinTable(&User{}, "Videos", &UserCreation{})
	if err != nil {
		log.Fatalf("自定义连接表设置失败,User: %s", err)
	}

}

// id 快捷用法返回一个Model{id:val}
func id(val int64) Model {
	return Model{ID: val}
}
