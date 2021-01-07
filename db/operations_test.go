package db_test

import (
	"github.com/luciodesimone/golang-bootcamp/beer"
	"github.com/luciodesimone/golang-bootcamp/db"
	"testing"
)

//we only use a copy of this variable to prevent data dependency
//between test cases
var mockBeer = beer.Beer{
	ID:          "mock-id",
	Desc:        "mock-beer-description",
	AlcoholCont: 5.5,
	IBU:         90,
	SRM:         41,
	AvgScore:    74.2,
}

func newMockStorage() map[string]beer.Beer {
	mockDB := make(map[string]beer.Beer)
	mockDB["mock-id"] = mockBeer
	return mockDB
}

func TestNewStorage(t *testing.T) {
	s := db.NewStorage("str")

	_, ok := s.(db.DB)

	if !ok {
		t.Errorf("A new storage must be a DB interface")
	}
}

func TestNOpenStorage(t *testing.T) {
	s := newMockStorage()
	open := db.OpenStorage(&s)
	b := mockBeer

	_, ok := s["mock-id"]

	if !ok {
		t.Errorf("The storage open storage should contain the data that is pointing the mock db")
	}

	bStored := beer.Beer{}

	err := open.GetBeer("mock-id", &bStored)

	if err != nil {
		t.Errorf("The storages doesn't point to the same location")
	}

	if bStored != b {
		t.Errorf("The storages doesn't have the same content")
	}
}

func TestCreateBeer(t *testing.T) {
	bMock := mockBeer

	s := db.NewStorage("str")

	b := s.CreateBeer(bMock)

	if b.Desc != b.Desc {
		t.Errorf("Expected description to be: %s recieved: %s", b.Desc, b.Desc)
	}

	if b.IBU != b.IBU {
		t.Errorf("Expected IBU to be: %d recieved: %d", b.IBU, b.IBU)
	}

	if b.AlcoholCont != b.AlcoholCont {
		t.Errorf("Expected alcohol content to be: %f recieved: %f", b.AlcoholCont, b.AlcoholCont)
	}

	if b.SRM != b.SRM {
		t.Errorf("Expected SRM to be: %d recieved: %d", b.SRM, b.SRM)
	}

	if b.AvgScore != b.AvgScore {
		t.Errorf("Expected average score to be: %f recieved: %f", b.AvgScore, b.AvgScore)
	}

}

func TestUpdateBeer(t *testing.T) {
	mockDB := newMockStorage()

	s := db.OpenStorage(&mockDB)
	b := mockBeer

	err := s.UpdateBeer("mock-id", b)

	if err != nil {
		t.Errorf("Error updating: %s", err.Error())
	}
}

func TestUpdateBeerNotFound(t *testing.T) {
	s := db.NewStorage("str")
	b := mockBeer

	err := s.UpdateBeer("non-existent", b)

	if err == nil {
		t.Errorf("Error updating, expected the record to don't exist")
	}
}

func TestGet(t *testing.T) {
	mockDB := newMockStorage()

	s := db.OpenStorage(&mockDB)
	b := beer.Beer{}

	err := s.GetBeer("mock-id", &b)

	if err != nil {
		t.Errorf("Error getting the record: %s", err.Error())
	}

	if b.ID != "mock-id" {
		t.Errorf("Error the record ID getted is incorrect: %s", err.Error())
	}

}

func TestGetNotFound(t *testing.T) {
	s := db.NewStorage("str")
	b := beer.Beer{}

	err := s.GetBeer("non-existent", &b)

	if err == nil {
		t.Errorf("Error get, expected the record to don't exist")
	}

	if b.ID != "" {
		t.Errorf("Error get, expected to don't return a beer")
	}
}

func TestDelete(t *testing.T) {
	mockDB := newMockStorage()

	s := db.OpenStorage(&mockDB)

	err := s.DeleteBeer("mock-id")

	_, ok := mockDB["mock-id"]

	if ok {
		t.Errorf("Error deleting the record: %s", err.Error())
	}
}

func TestDeleteNotFound(t *testing.T) {
	s := db.NewStorage("str")

	err := s.DeleteBeer("non-existent")

	if err == nil {
		t.Errorf("Error delete, expected the record to don't exist")
	}
}
