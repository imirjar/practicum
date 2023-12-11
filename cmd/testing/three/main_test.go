package main

import "testing"

func TestAddNew(t *testing.T) {
	family := Family{}
	person := Person{"Bob", "Jones", 12}
	if err := family.AddNew(Father, person); err != nil {
		t.Errorf("Куда тебе столько отцов?!")
	}
	if err := family.AddNew(Father, person); err == nil {
		t.Errorf("Куда тебе столько отцов?!")
	}
}
