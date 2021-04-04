package stats

import (
	"reflect"
	"testing"

	"github.com/ToirovSadi/bank/v2/pkg/types"
)

func TestPeriodsDynamic_excess1(t *testing.T) {
	res := PeriodsDynamic(map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}, map[types.Category]types.Money{
		"food": 20,
	})

	expect := map[types.Category]types.Money{
		"auto": -10,
		"food": 0,
	}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Error:PeriodsDynamic_excess1():\nwant:\n%v\ngot:\n%v\n", expect, res)
	}
}

func TestPeriodsDynamic_excess2(t *testing.T) {
	res := PeriodsDynamic(map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}, map[types.Category]types.Money{
		"auto":   10,
		"food":   25,
		"mobile": 5,
	})

	expect := map[types.Category]types.Money{
		"auto":   0,
		"food":   5,
		"mobile": 5,
	}
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Error:PeriodsDynamic_excess1():\nwant:\n%v\ngot:\n%v\n", expect, res)
	}
}

func TestCategoriesAvg(t *testing.T) {
	res := CategoriesAvg([]types.Payment{
		{ID: 1, Amount: 10, Category: "internet"},
		{ID: 2, Amount: 20, Category: "cafe"},
		{ID: 3, Amount: 30, Category: "auto"},
		{ID: 4, Amount: 40, Category: "internet"},
		{ID: 5, Amount: 50, Category: "cafe"},
	})

	expect := map[types.Category]types.Money{
		"internet": 25,
		"cafe":     35,
		"auto":     30,
	}

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("error:CategoriesAvg():\nwant: %v\n,got:%v\n", expect, res)
	}
}
