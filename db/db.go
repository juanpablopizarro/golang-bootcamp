package db

//db global storage in memory
var db = make(map[int64]interface{})

//id auto incremental value
var index int64 = 0

//Entity is an interface that allows to interact with the database
type Entity struct{}

//Create a new record in the database
func (e Entity) Create(ent interface{}) interface{} {
	index++
	db[index] = ent
	return ent
}

//Update a record by id in the database
func (e Entity) Update(ent interface{}, id int64) bool {
	_, ok := db[id]

	if !ok {
		return ok
	}

	db[id] = ent

	return ok
}

//Get a record by id in the database
func (e Entity) Get(id int64) interface{} {
	ent, ok := db[id]

	if !ok {
		return Entity{}
	}

	return ent
}

//Delete a record by id in the database
func (e Entity) Delete(id int64) bool {
	_, ok := db[id]

	if ok {
		delete(db, id)
	}

	return ok
}

//Connection emulates returning a real connection to the database
func Connection() *Entity {
	//emulates a connection
	return &Entity{}
}
