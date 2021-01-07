package db

import (
	"encoding/json"
	"fmt"
	"github.com/luciodesimone/golang-bootcamp/beer"
	"io/ioutil"
	"os"
)

type storage struct {
	db         map[string]beer.Beer
	connString string
	connection *os.File
}

//Open creates or open an existent file and uploads to memory all the content if it's valid json
func Open(connStr string) (DB, error) {
	//Creates a file with read write permissions
	var db = make(map[string]beer.Beer)
	f, err := os.OpenFile(connStr, os.O_APPEND|os.O_CREATE, os.ModePerm)

	if os.ErrNotExist != err {
		err = readFile(f, &db)
	} else {
		return nil, fmt.Errorf("The file entered doesn't exist: %s", err.Error())
	}

	return &storage{connString: connStr, db: db, connection: f}, nil
}

func readFile(f *os.File, db *map[string]beer.Beer) error {
	if err := json.NewDecoder(f).Decode(&db); err != nil {
		return fmt.Errorf("File format is incorrect or cant read the file: %s", err.Error())
	}

	return nil
}

//Close close the connection with the database and flush the data to disk
func (s *storage) Close() error {
	if s.connection == nil {
		return fmt.Errorf("Cant flush data, file isn't open")
	}

	err := flush(s.connection, s.db)

	if err != nil {
		return fmt.Errorf("An error ocurred flushing the data: %s", err.Error())
	}

	err = s.connection.Close()

	if err != nil {
		return fmt.Errorf("An error ocurred closing the file: %s", err.Error())
	}

	return nil
}

func flush(f *os.File, db map[string]beer.Beer) error {
	//using indent just to debug
	j, err := json.MarshalIndent(db, "", "\t")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(f.Name(), j, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}
