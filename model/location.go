package model

import (
	"errors"

	"gorm.io/gorm"
)

type Locations []Location

func (ls *Locations)Search() {
	CargoDB.Find(ls)
}

type Location struct{
	gorm.Model 
	Code string 
	Name string 
}

func NewLocation(code,name string) *Location{
	var location Location
	
	if err := location.Find(code); errors.Is(err, gorm.ErrRecordNotFound){
		location = Location{Code: code, Name: name}
		location.Add()
	}
	return &location
}


func(l *Location) Add(){
	CargoDB.Create(l)
}

func(l *Location) First(){
	CargoDB.First(l)
}

func(l *Location) Find(code string) error{
	result := CargoDB.Where("code = ?", code).Take(l)
	return result.Error
}






	