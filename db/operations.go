package db

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/luciodesimone/golang-bootcamp/beer"
)

//DB define all the operations that can be done with database
type DB interface {
	ConnectionString() string
	CreateBeer(beer.Beer) beer.Beer
	UpdateBeer(string, beer.Beer) error
	GetBeer(string, *beer.Beer) error
	DeleteBeer(string) error
	Close() error
}

//ConnectionString returns the filename of the database
func (s storage) ConnectionString() string {
	return s.ConnectionString()
}

//CreateBeer a new beer record in the database, returns error if the UUID creation fails
func (s *storage) CreateBeer(b beer.Beer) beer.Beer {
	uid := uuid.New()

	b.ID = uid.String()
	s.db[b.ID] = b

	return b
}

//UpdateBeers a beer record by id in the database
func (s *storage) UpdateBeer(id string, b beer.Beer) error {
	_, ok := s.db[id]

	if !ok {
		return fmt.Errorf("The ID entered doesn't exist")
	}

	b.ID = id
	s.db[id] = b

	return nil
}

//GetBeer a new record by id from the database, the beer will be returned in b returns an error if not found
func (s storage) GetBeer(id string, b *beer.Beer) error {
	record, ok := s.db[id]

	if !ok {
		return fmt.Errorf("No records returned")
	}

	*b = record

	return nil
}

//DeleteBeer phisically a record by id from the database, returns an error if not found
func (s *storage) DeleteBeer(id string) error {
	_, ok := s.db[id]

	if !ok {
		return fmt.Errorf("No records found")
	}

	delete(s.db, id)

	return nil
}
