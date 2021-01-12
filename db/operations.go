package db

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/luciodesimone/golang-bootcamp/beer"
)

//DB define all the operations that can be done with database
type DB interface {
	CreateBeer(*beer.Beer) beer.Beer
	UpdateBeer(string, beer.Beer) error
	GetBeer(string, *beer.Beer) error
	DeleteBeer(string) error
	Close() error
	flush() error
	RLock(string) (beer.Beer, error)
	WLock(string, beer.Beer)
	DeleteLock(string) error
}

//CreateBeer a new beer record in the database, returns error if the UUID creation fails
func (s *storage) CreateBeer(b *beer.Beer) beer.Beer {
	uid := uuid.New()

	b.ID = uid.String()

	s.WLock(b.ID, *b)
	return *b
}

//UpdateBeers a beer record by id in the database
func (s *storage) UpdateBeer(id string, b beer.Beer) error {
	_, err := s.RLock(id)

	if err != nil {
		return err
	}

	b.ID = id

	s.WLock(id, b)

	return nil
}

//GetBeer a new record by id from the database, the beer will be returned in b returns an error if not found
func (s *storage) GetBeer(id string, b *beer.Beer) error {
	beer, err := s.RLock(id)

	if err != nil {
		return fmt.Errorf("No records returned")
	}

	*b = beer

	return nil
}

//DeleteBeer phisically a record by id from the database, returns an error if not found
func (s *storage) DeleteBeer(id string) error {
	err := s.DeleteLock(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *storage) RLock(id string) (beer.Beer, error) {
	s.db.RLock()
	defer s.db.RUnlock()

	b, ok := s.db.m[id]

	if !ok {
		return beer.Beer{}, fmt.Errorf("No records found")
	}

	return b, nil
}

func (s *storage) WLock(id string, value beer.Beer) {
	s.db.Lock()
	defer s.db.Unlock()
	s.db.m[id] = value
	return
}

func (s *storage) DeleteLock(id string) error {
	s.db.Lock()
	defer s.db.Unlock()

	delete(s.db.m, id)

	_, ok := s.db.m[id]

	if !ok {
		return nil
	}

	return fmt.Errorf("Cant delete the record")
}
