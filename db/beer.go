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
