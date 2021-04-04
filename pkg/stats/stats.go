package stats

import (
	"github.com/ToirovSadi/bank/v2/pkg/types"
)

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	res := map[types.Category]types.Money{}
	cnt := map[types.Category]types.Money{}

	for _, payment := range payments {
		res[payment.Category] += types.Money(payment.Amount)
		cnt[payment.Category] += 1
	}

	for key := range res {
		res[key] /= cnt[key]
	}

	return res
}

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
