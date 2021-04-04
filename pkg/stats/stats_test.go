package stats

import (
	"reflect"
	"testing"

	"github.com/ToirovSadi/bank/v2/pkg/types"
)

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
