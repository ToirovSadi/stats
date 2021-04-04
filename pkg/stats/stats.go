package stats

import (
	"github.com/ToirovSadi/bank/v2/pkg/types"
)

func Avg(payment []types.Payment) types.Money {
	var sum types.Money = 0
	var cnt int = 0
	for _, payment := range payment {
		if payment.Status == types.StatusFail {
			continue
		}
		cnt++
		sum += payment.Amount
	}
	return sum / types.Money(cnt)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money = 0
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}
