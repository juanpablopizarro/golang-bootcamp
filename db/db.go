package db

//db global storage in memory
var db = make(map[int64]Beer)

//id auto incremental value
var index int64 = 0

//Create a new record in the database
func Create(beer Beer) Beer {
	index++
	beer.ID = index
	db[index] = beer
	return beer
}

//Update a record by id in the database
func Update(beer Beer, id int64) bool {
	_, ok := db[id]

	if !ok {
		return ok
	}

	db[id] = beer

	return ok
}

//Get a record by id in the database
func Get(id int64) Beer {
	beer, ok := db[id]

	if !ok {
		return Beer{}
	}

	beer.ID = id

	return beer
}

//Delete a record by id in the database
func Delete(id int64) bool {
	_, ok := db[id]

	if ok {
		delete(db, id)
	}

	return ok
}
