package model

import "cargo/common"

var CargoDeliveryHistory = make(map[string] []*Carrier)


type Cargo struct{
	ID string
	Describe string
}


func NewCargo(desc string) *Cargo{
	return &Cargo{ID: common.RandomString(5,"Cargo-"), Describe:desc}
}
