package db

//Beer is an entity that represents a record in the database
type Beer struct {
	id   int64
	Desc string

	//AlcoholCont is a percentage that describes the graduation of pure alcohol per 100ml of beer
	AlcoholCont float32

	//IBU is the index bitter unit that represents the bitterness
	IBU uint

	//SRM is the reference reference measurement for colour
	SRM uint

	//AvgScore is the average score /100 given by customers
	AvgScore float32
}

//Create beer
func Create(beer Beer) Beer {
	index++
	db[index] = beer
	return beer
}

//Update beer
func Update(beer Beer, id int64) bool {
	beer, ok := db[id]

	if !ok {
		return ok
	}

	db[id] = beer

	return ok
}

//Get beer
func Get(id int64) Beer {
	beer, ok := db[id]

	if !ok {
		return Beer{}
	}

	return beer
}

//Delete beer
func Delete(id int64) bool {
	_, ok := db[id]

	if ok {
		delete(db, id)
	}

	return ok
}
