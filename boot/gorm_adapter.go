package boot

import (
	"github.com/jinzhu/gorm"
	"log"
)

type GormAdapter struct {
	*gorm.DB
}

func NewGormAdapter() *GormAdapter {
	db, err := gorm.Open("mysql", "go_study:Zfy123456@tcp(rm-uf6sy00x06ad0d7q5go.mysql.rds.aliyuncs.com:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
	return &GormAdapter{DB: db}
}
