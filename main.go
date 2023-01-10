package main

import (
	"cargo/model"
	"cargo/service"
	"fmt"
	"time"
)

func main(){
	// init
	
	// a cargo
	cargo := model.NewCargo("cargo one")
	
	// startlocation
	loc1 := model.NewLocation("zz","郑州")
	loc2 := model.NewLocation("bj","北京")
	loc3 := model.NewLocation("hz","杭走")

	// a carrier
	carrier1 := model.NewCarrier(loc1,loc2, time.Now(), time.Now())
	carrier2 := model.NewCarrier(loc2,loc3, time.Now(), time.Now())
	carrier3 := model.NewCarrier(loc3,loc1, time.Now(), time.Now())

	service.CargoBook(cargo, carrier1)
	service.CargoBook(cargo, carrier2)
	service.CargoBook(cargo, carrier3)

	// 获取货物线路状态
	for _,i:= range model.CargoDeliveryHistory[cargo.ID]{
		fmt.Println(i.Marshal())
	}

}