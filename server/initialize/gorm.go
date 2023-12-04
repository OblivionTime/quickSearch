package initialize

import (
	"fmt"
	"os"
	"quickSearch/global"
	"quickSearch/model/quickSearch"
	"time"

	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/sqlite"
)

func GormSqlite() *gorm.DB {
	fmt.Println("init sqlite3")
	var err error
	DB, err := gorm.Open("sqlite3", "./quickSearch.db")
	if err != nil {
		panic(err)
	}
	//允许单表创建
	DB.SingularTable(true)
	//关闭sql语句日志
	DB.LogMode(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	DB.DB().SetConnMaxLifetime(time.Hour)

	return DB
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		quickSearch.FileInfo{},
		quickSearch.Category{},
	).Error
	//判断数据库是否存在,如果不存在则初始化相关数据
	if result := db.Table("category").First(&quickSearch.Category{}); result.RowsAffected == 0 {
		db.Create(&quickSearch.Category{
			Name: "不知名的分组",
		})
	}
	if err != nil {
		fmt.Println("register table failed", err)
		os.Exit(0)
	}
	fmt.Println("register table success")
}
