package main

import (
	"fmt"
	"github.com/luciodesimone/golang-bootcamp/beer"
	"github.com/luciodesimone/golang-bootcamp/db"
)

func main() {
	stout := beer.New("Creamy black dense beer with coffee notes, ideal before dinner", 42, 8, 40, 92)
	bitter := beer.New("dark amber light beer, ideal for summer", 57, 5.5, 10, 63)
	barley := beer.New("Strong, dense red/dark beer, ideal for late night with friends", 70, 12, 22, 86)

	s := db.NewStorage()

	st := s.Create(stout)
	bi := s.Create(bitter)
	ba := s.Create(barley)

	fmt.Printf("Created: %+v\n", st)
	fmt.Printf("Created: %+v\n", bi)
	fmt.Printf("Created: %+v\n", ba)

	gst := beer.Beer{}
	gbi := beer.Beer{}
	gba := beer.Beer{}
	gnf := beer.Beer{}

	err := s.Get(st.ID, &gst)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	err = s.Get(bi.ID, &gbi)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	err = s.Get(ba.ID, &gba)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	err = s.Get("not-found", &gnf)

	if err != nil {
		fmt.Printf("Not found as expected: %s\n", err)
	}

	fmt.Printf("Found: %+v\n", gst)
	fmt.Printf("Found: %+v\n", gbi)
	fmt.Printf("Found: %+v\n", gba)

	err = s.Delete(bi.ID)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	err = s.Delete("not found")

	if err != nil {
		fmt.Printf("Not found as expected: %s\n", err)
	}

	err = s.Get(bi.ID, &gbi)

	if err != nil {
		fmt.Printf("Deleted as expected: %s\n", err)
	}

	creamStout := beer.Beer{
		Desc:        "More creamy less black with coffee notes",
		IBU:         26,
		AlcoholCont: 5,
		SRM:         34,
		AvgScore:    78,
	}

	err = s.Update(st.ID, creamStout)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	cs := beer.Beer{}

	err = s.Get(st.ID, &cs)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("Found: %+v\n", cs)
}
