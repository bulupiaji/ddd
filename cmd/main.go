package main

import (
	"cargo/model"
	"fmt"
)

func main(){
	// var locations model.Locations
	// locations.Search()
	// fmt.Println(locations)
	
	// one order
	var src = model.NewLocation("SZ","深圳")
	var des = model.NewLocation("BJ","北京")
	var shipper = model.NewCustomer("shipper")
	var receiver = model.NewCustomer("receiver")
	var payer = model.NewCustomer("payer")
	fmt.Println(src,des,shipper,receiver,payer)
	var cargo = model.NewCargo("test cargo order",shipper.ID, receiver.ID, payer.ID, src.Code, des.Code)

	fmt.Println(cargo)

}