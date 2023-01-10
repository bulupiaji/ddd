package model

type Location struct{
	Code string
	Name string 
}

func NewLocation(name, code string) Location{
	return Location{Name: name, Code: code}
}
