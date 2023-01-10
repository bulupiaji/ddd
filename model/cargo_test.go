package model

import (
	"testing"
)

func TestCargo(t *testing.T){
	c := NewCargo("A TEST CAGO")
	t.Error(c)
}