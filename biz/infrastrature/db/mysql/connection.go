package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

// todo
func init_not_use() {
	dsn := "root:5rgzH9l3LJ678_oyBnrk@tcp(172.30.179.120:6606)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Connection = db
}
