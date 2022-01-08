package evenup

import (
	"fmt"
	"math"
	"sort"
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

	transactions := buildTransactionList(debts, initialLedger.Names)

	return transactions, nil
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
	sortFriendsByDebt(friends, debts)
	var transactions []string

	i := 0              // left side pointer
	j := len(debts) - 1 // right side pointer
	for i < j {
		transactions = clearEasyMatches(debts, friends, transactions, i, j)

		if isZero(debts[i]) && isZero(debts[j]) {
			i++
			j--
			continue
		}

		largerAbsValIdx := -1
		smallerAbsValIdx := -1

		if abs(debts[i]) <= abs(debts[j]) {
			smallerAbsValIdx = i
			largerAbsValIdx = j
		} else {
			smallerAbsValIdx = j
			largerAbsValIdx = i
		}

		txAmount := abs(debts[smallerAbsValIdx])

		// adjust values based on transaction amount.
		// make sure our signs are right by checking whether or not the left side is positive or negative
		// (aka i or j -> implicit: val at index i should always be neg [is owed], val at index j should always be pos [owes])
		alteration := float32(txAmount) // this is the amount that we alter the larger value by
		if largerAbsValIdx == i {
			alteration = alteration * -1
		}

		debts[largerAbsValIdx] -= alteration
		debts[smallerAbsValIdx] = 0

		transactions = append(transactions, fmt.Sprintf("%s pays %s $%.2f", friends[j], friends[i], txAmount))

		// adjust left and right pointers
		for isZero(debts[i]) && i < len(debts)-1 {
			i++
		}

		for isZero(debts[j]) && j > 0 {
			j--
		}
	}

	return transactions
}

// sort both debts and friends based on debts in ascending order (neg to pos === [is owed] to [owes])
func sortFriendsByDebt(friends []string, debts []float32) {
	// group the 2 fields together so we can sort them together
	friendDebts := make([]struct {
		friend string
		debt   float32
	}, len(friends))

	for i := range friends {
		friendDebts[i].friend = friends[i]
		friendDebts[i].debt = debts[i]
	}

	sort.Slice(friendDebts, func(i, j int) bool {
		return friendDebts[i].debt < friendDebts[j].debt // asc
	})

	// alter the underlying friends and debts data structures
	for i := range friendDebts {
		friends[i] = friendDebts[i].friend
		debts[i] = friendDebts[i].debt
	}
}

// clear the easy matches (any owe <--> owed equalities). These are the quick low hanging fruit that the main alg may miss, without which could result in extra transactions.
// modifies debts
// assumes debts is sorted asc
func clearEasyMatches(debts []float32, friends []string, transactions []string, i, j int) []string {
	for debts[i] < 0 {
		for debts[j] >= abs(debts[i]) {
			if isZero(debts[j] - abs(debts[i])) {
				transactions = append(transactions, fmt.Sprintf("%s pays %s $%.2f", friends[j], friends[i], debts[j]))

				debts[i] = 0
				debts[j] = 0

				break
			}
			j--
		}
		i++
	}

	return transactions
}

func validateDebts(debts []float32) error {
	if !isZero(sum(debts)) {
		return fmt.Errorf("debts must add up to zero")
	}

	return nil
}

// the granularity we care about is a max differential of 0.01 between values. float arithmetic may cause problems if we don't specify this.
// this is *effectively* zero for our purposes
func isZero(num float32) bool {
	return num > -0.01 && num < 0.01
}

func sum(vals []float32) float32 {
	var sum float32
	for _, v := range vals {
		sum += v
	}

	return sum
}

func abs(num float32) float32 {
	return float32(math.Abs(float64(num)))
}
