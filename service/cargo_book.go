package service

import "cargo/model"

func CargoBook(cargo *model.Cargo, carrier *model.Carrier) {
	// some rules
	model.CargoDeliveryHistory[cargo.ID] = append(model.CargoDeliveryHistory[cargo.ID], carrier)

}
