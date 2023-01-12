package model

import (

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var CargoDB *gorm.DB

func init(){
	initDB()
	initTable()
}

func initDB() {
	db, err := gorm.Open(sqlite.Open("cargo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connet database")
	}    	
	CargoDB = db
}

func Truncate(db *gorm.DB, dst ...interface{}){
	db.Migrator().DropTable(dst...)
}

func initTable(){
	initLocation()
	initCustomer()
	initCargo()
}

func initCargo(){
	var cargo = &Cargo{}
	if CargoDB.Migrator().HasTable(cargo){
		return 
	}
	CargoDB.AutoMigrate(cargo)
}

func initCustomer(){
	var customer = &Customer{}
	if CargoDB.Migrator().HasTable(customer){
		return 
	}
	CargoDB.AutoMigrate(customer)
}

func initLocation(){
	var location = &Location{}
	if CargoDB.Migrator().HasTable(location){
		return 
	}
	CargoDB.AutoMigrate(location)
	// 初始化location地址
	l1 := NewLocation("ZZ","郑州")
	l2 := NewLocation("HZ","杭州")
	l3 := NewLocation("BJ","北京")

	l1.Add()
	l2.Add()
	l3.Add()
}
