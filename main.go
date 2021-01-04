package main

import (
	"github.com/luciodesimone/golang-bootcamp/db"
	"log"
)

func main() {
	stout := db.Beer{
		Desc:        "Creamy black dense beer with coffee notes, ideal before dinner",
		IBU:         42,
		AlcoholCont: 8,
		SRM:         40,
		AvgScore:    92,
	}

	bitter := db.Beer{
		Desc:        "dark amber light beer, ideal for summer",
		IBU:         57,
		AlcoholCont: 5.5,
		SRM:         10,
		AvgScore:    63,
	}

	barley := db.Beer{
		Desc:        "Strong, dense red/dark beer, ideal for late night with friends",
		IBU:         70,
		AlcoholCont: 12,
		SRM:         22,
		AvgScore:    86,
	}

	s := db.Create(stout)
	bi := db.Create(bitter)
	ba := db.Create(barley)

	log.Printf("Created: %v", s)
	log.Printf("Created: %v", bi)
	log.Printf("Created: %v", ba)

	s = db.Get(1)
	bi = db.Get(2)
	ba = db.Get(3)
	nf := db.Get(4)

	log.Printf("Found: %v", s)
	log.Printf("Found: %v", bi)
	log.Printf("Found: %v", ba)
	log.Printf("Not found: %v", nf)

	log.Printf("Found: %v", s)
	log.Printf("Found: %v", bi)
	log.Printf("Found: %v", ba)
	log.Printf("Not found: %v", nf)

	del := db.Delete(2)

	log.Printf("Deleted: %v", del)

	del = db.Delete(7)

	log.Printf("Deleted: %v", del)

	bi = db.Get(2)

	log.Printf("Not found: %v", bi)

	creamStout := db.Beer{
		Desc:        "More creamy less black with coffee notes",
		IBU:         26,
		AlcoholCont: 5,
		SRM:         34,
		AvgScore:    78,
	}

	upd := db.Update(creamStout, 1)

	log.Printf("Updated: %v", upd)

	upd = db.Update(creamStout, 2)

	log.Printf("Updated: %v", upd)

	cs := db.Get(1)

	log.Printf("Changed: %v", cs)

	cs = db.Get(2)

	log.Printf("Not changed: %v", cs)
}
