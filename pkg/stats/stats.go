package stats

import (
	"github.com/ToirovSadi/bank/pkg/types"
)

func Avg(payment []types.Payment) types.Money {
	var sum types.Money = 0
	var n int = len(payment)
	for _, payment := range payment {
		sum += payment.Amount
	}
	return sum / types.Money(n)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money = 0
	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}
