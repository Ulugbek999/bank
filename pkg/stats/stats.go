package stats

import (
	"github.com/Ulugbek999/bank1/v2/pkg/types"
	
	
	
	
)












//Avg расчитовает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	sumPayments := types.Money(0)
	indexPayments := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		moneyPayments := payment.Amount
		sumPayments += moneyPayments
		indexPayments++
	}
	return sumPayments / indexPayments
}


// TotalInCategory находит 	сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sumPayments := types.Money(0)
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		if payment.Status == types.StatusFail {
			continue
		}
		moneyPayments := payment.Amount
		sumPayments += moneyPayments
	}
	return sumPayments
}