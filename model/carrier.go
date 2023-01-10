package model

import (
	"cargo/common"
	"encoding/json"
	"time"
)

type Carrier struct {
	ScheduleId    string  `json:"schedule_id"`
	SrcLocation   Location`json:"src_location"`
	DesLocation   Location`json:"des_location"`
	DepartureTime time.Time`json:"departur_time"`
	ArrivalTime   time.Time`json:"arrival_time"`
}

func NewCarrier(src, des Location, dep, arr time.Time) *Carrier {
	return &Carrier{ScheduleId: common.RandomString(5, "carrier-"),
		SrcLocation:   src,
		DesLocation:   des,
		ArrivalTime:   arr,
		DepartureTime: dep}
}

func(c *Carrier) Marshal() string{
	json_carrier, err := json.Marshal(c)
	if err != nil{
		panic(err)
	}
	return string(json_carrier)
}
