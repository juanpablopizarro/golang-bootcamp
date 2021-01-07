package beer_test

import (
	"github.com/luciodesimone/golang-bootcamp/beer"
	"testing"
)

func TestNew(t *testing.T) {
	expect := beer.Beer{
		Desc:        "desc",
		IBU:         1,
		AlcoholCont: 2.3,
		SRM:         4,
		AvgScore:    5,
	}

	b := beer.New("desc", 1, 2.3, 4, 5)

	if b.Desc != expect.Desc {
		t.Errorf("Expected description to be: %s recieved: %s", expect.Desc, b.Desc)
	}

	if b.IBU != expect.IBU {
		t.Errorf("Expected IBU to be: %d recieved: %d", expect.IBU, b.IBU)
	}

	if b.AlcoholCont != expect.AlcoholCont {
		t.Errorf("Expected alcohol content to be: %f recieved: %f", expect.AlcoholCont, b.AlcoholCont)
	}

	if b.SRM != expect.SRM {
		t.Errorf("Expected SRM to be: %d recieved: %d", expect.SRM, b.SRM)
	}

	if b.AvgScore != expect.AvgScore {
		t.Errorf("Expected average score to be: %f recieved: %f", expect.AvgScore, b.AvgScore)
	}

}
