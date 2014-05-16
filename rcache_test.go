package rcache

import (
	"testing"
)

type TaxiDriver struct {
	FirstName      string
	LastName       string
	Surname        string
	Identification string
	Email          string
	HomePhone      string
	Mobile         string
	OtherDetails   string
}

func TestSet(t *testing.T) {
	key := "mykey"
	value := &TaxiDriver{"Greivin", "López", "Paniagua", "205480941", "greivin.lopez@gmail.com", "24410131", "88684734", ""}

	err := Set(key, value)
	if err != nil {
		t.Errorf("Throwing error in Set(key, value): %v", err)
	}
}

func TestGet(t *testing.T) {
	key := "getkey"
	value := &TaxiDriver{"Greivin", "López", "Paniagua", "205480941", "greivin.lopez@gmail.com", "24410131", "88684734", ""}

	err := Set(key, value)
	if err != nil {
		t.Errorf("Throwing error in Set(key, value): %v", err)
	}

	value = &TaxiDriver{}
	err = Get(key, &value)
	if err != nil {
		t.Errorf("Throwing error in Get(key, value): %v", err)
	}
	if value.FirstName != "Greivin" {
		t.Errorf("Invalid value returned by GET %v", value)
	}

	err = Get("notexist", &value)
	if err == nil {
		t.Errorf("Not returning error when key not exist: %v", err)
	}
}
