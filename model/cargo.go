package model

import (
	"cargo/common"

	"gorm.io/gorm"
)

// cargo 应该是cargo的传输记录

// var CargoDeliveryHistory = make(map[string] []*Carrier)


type Cargo struct{
	gorm.Model
	TrackID string 
	Describe string 
	ShipperID string
	RiceverID  string 
	PayerID string
	SrcID string
	DesID string
}


func NewCargo(desc,shipper_id, ricever_id, payer_id, src_id , des_id string) *Cargo{

	var cargo = Cargo{TrackID: common.RandomString(5,"Cargo-"), 
	Describe: desc,
	ShipperID: shipper_id,
	RiceverID: ricever_id,
	PayerID: payer_id,
	SrcID: src_id,
	DesID: des_id}

	cargo.Add()
	return &cargo
}

func(c *Cargo) Add(){
	CargoDB.Create(c)
}

func(c *Cargo) Find(trackid string) error{
	result:=CargoDB.Where("trackid = ?", trackid).Take(c)
	return result.Error
}