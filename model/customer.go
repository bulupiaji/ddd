package model

import (
	"cargo/common"
	"errors"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID   string
	Name string
}

func NewCustomer(name string) *Customer{
	var customer Customer
	if err := customer.Find(name); errors.Is(err, gorm.ErrRecordNotFound){

		customer = Customer{ID: common.RandomString(5,"customer-"), Name: name}
		customer.Add()
	}
	return &customer
}

func(c *Customer) Add(){
	CargoDB.Create(c)
}

func(c *Customer) Find(name string) error{
	result := CargoDB.Where("name = ?", name).Take(c)
	return result.Error
}

