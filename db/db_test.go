package db_test

import (
	"github.com/luciodesimone/golang-bootcamp/db"
	"testing"
)

var mockBeer = db.Beer{
	ID:          1,
	Desc:        "mock-beer-description",
	AlcoholCont: 5.5,
	IBU:         90,
	SRM:         41,
	AvgScore:    74.2,
}

func TestDb(t *testing.T) {
	//Create
	b := db.Create(mockBeer)

	if b != mockBeer {
		t.Errorf("DB creation failed, expecting: %v", mockBeer)
	}

	if b.ID != int64(1) {
		t.Errorf("DB creation failed, expecting id to be: 1, recieved: %d", b.ID)
	}

	//Get
	record := db.Get(1)

	if record != mockBeer {
		t.Errorf("DB get failed, expecting: %v", mockBeer)
	}

	record = db.Get(2)

	if record == mockBeer {
		t.Errorf("DB get failed, expecting not to be defined")
	}

	//Update
	newBeer := db.Beer{
		Desc:        "mock-new-beer-description",
		AlcoholCont: 10,
		IBU:         48,
		SRM:         80,
		AvgScore:    54.2,
	}

	succ := db.Update(newBeer, 1)

	if !succ {
		t.Errorf("DB update failed, expecting update to success")
	}

	succ = db.Update(newBeer, 15)

	if succ {
		t.Errorf("DB update failed, expecting update to fail")
	}

	//Delete
	succ = db.Delete(2)

	if succ {
		t.Errorf("DB delete failed, expecting update to fail")
	}

	succ = db.Delete(1)

	if !succ {
		t.Errorf("DB delete failed, expecting update to success")
	}
}
