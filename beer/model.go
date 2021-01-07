package beer

//Beer is an entity that represents a record in the database
type Beer struct {
	//CAUTION: Internal value to manage a unique entity in the database
	ID string

	//AlcoholCont is a percentage that describes the graduation of pure alcohol per 100ml of beer
	AlcoholCont float32

	//Desc description of the beer taste or qualities
	Desc string

	//IBU is the index bitter unit that represents the bitterness
	IBU uint

	//SRM is the reference reference measurement for colour
	SRM uint

	//AvgScore is the average score /100 given by customers
	AvgScore float32
}

//New creates a struct beer and returns it
func New(Desc string, IBU uint, AlcoholCont float32, SRM uint, AvgScore float32) Beer {
	return Beer{
		Desc:        Desc,
		AlcoholCont: AlcoholCont,
		IBU:         IBU,
		SRM:         SRM,
		AvgScore:    AvgScore,
	}
}
