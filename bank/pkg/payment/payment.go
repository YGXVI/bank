package payment

import "github.com/YGXVI/bank/v2/pkg/types"

func Max(payments []types.Payment) types.Payment {
	max := payments[0]

	for _, payment := range payments {
		if payment.Amount > max.Amount {
			max = payment
		}
	}

	return max
}