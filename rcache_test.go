package rcache

import (
	"labix.org/v2/mgo/bson"
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

func TestBsonIds(t *testing.T) {
	key := bson.NewObjectId()
	value := &TaxiDriver{"Georges", "Bizet", "", "203330222", "georges.bizet@gmail.com", "24410131", "88684734", ""}
	err := Set(key, value)
	if err != nil {
		t.Errorf("Throwing error in Set(key, value) when using bson object: %v", err)
	}

	valueRead := &TaxiDriver{}
	err = Get(key, &valueRead)
	if err != nil {
		t.Errorf("Throwing error in Get(key, value) when using bson object: %v", err)
	}
	if valueRead.FirstName != "Georges" {
		t.Errorf("Invalid value returned by GET %v", value)
	}
}

func TestDel(t *testing.T) {
	key := "delkey"
	value := &TaxiDriver{"Greivin", "López", "Paniagua", "205480941", "greivin.lopez@gmail.com", "24410131", "88684734", ""}

	err := Set(key, value)
	if err != nil {
		t.Errorf("Throwing error in Set(key, value): %v", err)
	}

	err = Del(key)

	valueRead := &TaxiDriver{}
	err = Get(key, &valueRead)
	if err == nil {
		t.Errorf("Failed to delete from Redis: %v", key)
	}
}
