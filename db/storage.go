package db

import (
	"encoding/json"
	"fmt"
	"github.com/luciodesimone/golang-bootcamp/beer"
	"io/ioutil"
	"os"
	"sync"
)

//RWMap contains the database structure and locks to prevent data
//dependency on concurrent access
type RWMap struct {
	sync.RWMutex
	m map[string]beer.Beer
}

type storage struct {
	db         *RWMap
	DB         *RWMap
	connString string
	connection *os.File
}

//New creates a storage with the
func New(connString string, connection *os.File) DB {
	return &storage{
		db:         &RWMap{m: make(map[string]beer.Beer)},
		connString: connString,
		connection: connection,
	}
}

//Open creates or open an existent file and uploads to memory all the content if it's valid json
func Open(connStr string) (DB, error) {
	f, err := os.OpenFile(connStr, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if err != nil {
		return nil, fmt.Errorf("Cant open file: %s", err.Error())
	}

	//read the database
	db := New(connStr, f)

	err = json.NewDecoder(f).Decode(&db)

	if err != nil {
		return nil, fmt.Errorf("File format is incorrect or cant read the file: %s", err.Error())
	}

	return db, nil
}

//Close close the connection with the database and flush the data to disk
func (s *storage) Close() error {
	if s.connection == nil {
		return fmt.Errorf("Cant flush data, file isn't open")
	}

	err := s.flush()

	if err != nil {
		return fmt.Errorf("An error ocurred flushing the data: %s", err.Error())
	}

	err = s.connection.Close()

	if err != nil {
		return fmt.Errorf("An error ocurred closing the file: %s", err.Error())
	}

	return nil
}

func (s *storage) flush() error {
	//using indent just to debug
	j, err := json.MarshalIndent(s.db.m, "", "\t")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.connString, j, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
