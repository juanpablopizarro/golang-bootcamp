package db

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/luciodesimone/golang-bootcamp/beer"
)

//DB define all the operations that can be done with database
type DB interface {
	Create(beer.Beer) beer.Beer
	Update(string, beer.Beer) error
	Get(string, *beer.Beer) error
	Delete(string) error
}

type storage struct {
	db map[string]beer.Beer
}

//NewStorage creates a new beer database
func NewStorage() DB {
	return &storage{
		db: make(map[string]beer.Beer),
	}
}

//OpenStorage allows to use the interface of the database
//on a existent one
func OpenStorage(db *map[string]beer.Beer) DB {
	return &storage{
		db: *db,
	}
}

//Create a new beer record in the database, returns error if the UUID creation fails
func (s *storage) Create(b beer.Beer) beer.Beer {
	uid := uuid.New()

	b.ID = uid.String()
	s.db[b.ID] = b

	return b
}

//Updates a beer record by id in the database
func (s *storage) Update(id string, b beer.Beer) error {
	_, ok := s.db[id]

	if !ok {
		return fmt.Errorf("The ID entered doesn't exist")
	}

	b.ID = id
	s.db[id] = b

	return nil
}

//Get a new record by id from the database, the beer will be returned in b returns an error if not found
func (s storage) Get(id string, b *beer.Beer) error {
	record, ok := s.db[id]

	if !ok {
		return fmt.Errorf("No records returned")
	}

	*b = record

	return nil
}

//Delete phisically a record by id from the database, returns an error if not found
func (s *storage) Delete(id string) error {
	_, ok := s.db[id]

	if !ok {
		return fmt.Errorf("No records found")
	}

	delete(s.db, id)

	return nil
}
