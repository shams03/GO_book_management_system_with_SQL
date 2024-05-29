package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (db * gorm.DB) //these simple backets can be removed

func Connect(){
	d,err:= gorm.Open("mysql","shams:=xoxo/pikachu?cahrset=utf8&parseTime=True&loc=local")
	if err!=nil{
		panic(err)
	}
	db=d;
}

func GetDb() *gorm.DB{
	return db;
}