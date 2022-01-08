package evenup

import (
	"fmt"
)

type InitialLedger struct {
	Names         []string
	PaymentValues []float32
}

// for a given inital ledger (names and initial payments), return a list of transactions that evens everyone up
func CalculateTransactions(initialLedger InitialLedger) ([]string, error) {
	debts := calculateDebts(initialLedger.PaymentValues)
	err := validateDebts(debts)
	if err != nil {
		return []string{}, err
	}

	// buildTransactionsMap - maybe name this better -> buildTransactionList
	return []string{}, nil
}

// for all initial payment values, calculate the amount the person owes (positive value) or is owed (negative value)
func calculateDebts(payments []float32) []float32 {
	debts := make([]float32, len(payments))
	evenAmount := sum(payments) / float32(len(payments))

	for i, payment := range payments {
		debts[i] = evenAmount - payment
	}

	return debts
}

func buildTransactionList(debts []float32, friends []string) []string {
	return []string{}
}

func validateDebts(debts []float32) error {
	if !isEffectivelyZero(sum(debts)) {
		return fmt.Errorf("debts must add up to zero")
	}

	return nil
}

// the granularity we care about is a max differential of 0.01 between values. float arithmetic may cause problems if we don't specify this
func isEffectivelyZero(num float32) bool {
	return num > -0.01 && num < 0.01
}

func sum(vals []float32) float32 {
	var sum float32
	for _, v := range vals {
		sum += v
	}

	return sum
}
