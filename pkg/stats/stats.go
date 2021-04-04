package stats

import (
	"github.com/Ulugbek999/bank1/pkg/types"
)












//Avg расчитовает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	countPayments := types.Money(len(payments))
	sumPayments := types.Money(0)
	for _, payment := range payments {
		moneyPayments := payment.Amount
		sumPayments += moneyPayments
	}
	return sumPayments / countPayments
}


// TotalInCategory находит 	сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sumPayments := types.Money(0)
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		moneyPayments := payment.Amount
		sumPayments += moneyPayments
	}
	return sumPayments
}