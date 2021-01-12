package main

import (
	"fmt"
	"github.com/luciodesimone/golang-bootcamp/beer"
	"github.com/luciodesimone/golang-bootcamp/db"
	"os"
)

func main() {
	stout := beer.New("Creamy black dense beer with coffee notes, ideal before dinner", 42, 8, 40, 92)
	bitter := beer.New("dark amber light beer, ideal for summer", 57, 5.5, 10, 63)
	barley := beer.New("Strong, dense red/dark beer, ideal for late night with friends", 70, 12, 22, 86)
	newB := beer.Beer{}

	//you must provide a valid json file
	conn, err := db.Open("beers.json")

	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		os.Exit(1)
	}

	conn.CreateBeer(&stout)
	conn.CreateBeer(&bitter)
	conn.CreateBeer(&barley)

	err = conn.GetBeer(stout.ID, &newB)

	if err != nil {
		fmt.Printf("Error getting beer: %s\n", err)
	} else {
		fmt.Printf("Beer!: %v\n", newB)
	}

	defer conn.Close()

	err = conn.DeleteBeer(stout.ID)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		err = conn.GetBeer(stout.ID, &newB)

		if err != nil {
			fmt.Printf("Poof no more beer: %s\n", err)
		}
	}

	err = conn.DeleteBeer("not found")

	if err != nil {
		fmt.Printf("Not found as expected: %s\n", err)
	}

	creamStout := beer.Beer{
		Desc:        "More creamy less black with coffee notes",
		IBU:         26,
		AlcoholCont: 5,
		SRM:         34,
		AvgScore:    78,
	}

	err = conn.UpdateBeer(bitter.ID, creamStout)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	cs := beer.Beer{}

	err = conn.GetBeer(bitter.ID, &cs)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("Found: %+v\n", cs)
}
