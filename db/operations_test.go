package db_test

import (
	"github.com/luciodesimone/golang-bootcamp/beer"
	"github.com/luciodesimone/golang-bootcamp/db"
	"os"
	"reflect"
	"testing"
)

//we only use a copy of this variable to prevent data dependency
//between test cases
var mockBeer = beer.Beer{
	Desc:        "mock-beer-description",
	AlcoholCont: 5.5,
	IBU:         90,
	SRM:         41,
	AvgScore:    74.2,
}

func newMockStorage() (db.DB, string) {
	f := &os.File{}
	s := db.New("mock-fname", f)
	s.CreateBeer(&mockBeer)
	return s, mockBeer.ID
}

func TestReadLocks(t *testing.T) {
	t.Parallel()
	newMockStorage()

	t.Run("Update beer group", func(t *testing.T) {
		t.Run("Update beer 1", TestUpdateBeerNotFound)
		t.Run("Update beer 2", TestUpdateBeerNotFound)
		t.Run("Update beer 3", TestUpdateBeerNotFound)
	})
}

func TestWriteLocks(t *testing.T) {
	t.Parallel()
	newMockStorage()

	t.Run("Get beer group", func(t *testing.T) {
		t.Run("Get beer 1", TestGetBeerNotFound)
		t.Run("Ger beer 2", TestGetBeerNotFound)
		t.Run("Ger beer 3", TestGetBeerNotFound)
	})
}

func TestDeleteLocks(t *testing.T) {
	t.Parallel()
	newMockStorage()

	t.Run("Delete beer group", func(t *testing.T) {
		t.Run("Delete beer 1", TestDeleteBeerNotFound)
		t.Run("Delete beer 2", TestDeleteBeerNotFound)
		t.Run("Delete beer 3", TestDeleteBeerNotFound)
	})
}

func TestNew(t *testing.T) {
	t.Parallel()

	fName := "mock-file-name"
	mFile := &os.File{}

	s := db.New(fName, mFile)

	if reflect.TypeOf(s.Close).Kind() != reflect.Func {
		t.Errorf("Expected Close to be defined in the new storage")
	}

	if reflect.TypeOf(s.CreateBeer).Kind() != reflect.Func {
		t.Errorf("Expected CreateBeer to be defined in the new storage")
	}

	if reflect.TypeOf(s.DeleteBeer).Kind() != reflect.Func {
		t.Errorf("Expected DeleteBeer to be defined in the new storage")
	}

	if reflect.TypeOf(s.UpdateBeer).Kind() != reflect.Func {
		t.Errorf("Expected UpdateBeer to be defined in the new storage")
	}

	if reflect.TypeOf(s.GetBeer).Kind() != reflect.Func {

	}

}

func TestCreateBeer(t *testing.T) {
	t.Parallel()

	bMock := mockBeer
	s, _ := newMockStorage()
	b := s.CreateBeer(&bMock)

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
	s, beerId := newMockStorage()
	b := mockBeer

	err := s.UpdateBeer(beerId, b)

	if err != nil {
		t.Errorf("Error updating: %s", err.Error())
	}
}

func TestUpdateBeerNotFound(t *testing.T) {
	t.Parallel()

	s, _ := newMockStorage()
	b := mockBeer

	err := s.UpdateBeer("non-existent", b)

	if err == nil {
		t.Errorf("Error updating, expected the record to don't exist")
	}
}

func TestGetBeer(t *testing.T) {
	s, beerId := newMockStorage()

	b := beer.Beer{}

	err := s.GetBeer(beerId, &b)

	if err != nil {
		t.Errorf("Error getting the record: %s", err.Error())
	}

	if b.ID != beerId {
		t.Errorf("Error the record ID getted is incorrect: %s", b.ID)
	}

}

func TestGetBeerNotFound(t *testing.T) {
	t.Parallel()

	s, _ := newMockStorage()
	b := beer.Beer{}

	err := s.GetBeer("non-existent", &b)

	if err == nil {
		t.Errorf("Error get, expected the record to don't exist")
	}

	if b.ID != "" {
		t.Errorf("Error get, expected to don't return a beer")
	}
}

func TestDeleteBeer(t *testing.T) {
	s, beerId := newMockStorage()

	err := s.DeleteBeer(beerId)

	if err != nil {
		t.Errorf("Error deleting the record: %s", err.Error())
	}
}

func TestDeleteBeerNotFound(t *testing.T) {
	t.Parallel()

	s, _ := newMockStorage()

	//if not found it will return no error
	err := s.DeleteBeer("non-existent")

	if err != nil {
		t.Errorf("Error delete, expected the record to don't exist")
	}
}
