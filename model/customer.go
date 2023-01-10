package model

import "cargo/common"

type Customer struct {
	Name string
	ID   string
}

func NewCustomer(name string) *Customer{
	return &Customer{Name: name , ID: common.RandomString(5,"customer")}
}



