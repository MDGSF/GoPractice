package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (Product) TableName() string {
	return "product"
}

func main() {
	log, _ := zap.NewDevelopment()
	defer log.Sync()

	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gormdemo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error("failed to open mysql", zap.Any("err", err))
		return
	}
	defer db.Close()
	log.Info("open mysql success")

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "tb_" + defaultTableName
	}

	// auto generate table.
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "L1212", Price: 1000})
}
