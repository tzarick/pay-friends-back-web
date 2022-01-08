package evenup

import (
	"reflect"
	"testing"
)

type DebtTestCase struct {
	id            int
	inputPayments []float32
	expected      []float32
}

func TestCalculateDebts(t *testing.T) {
	cases := []DebtTestCase{
		{
			id:            1,
			inputPayments: []float32{60, 75, 20, 0},
			expected:      []float32{-21.25, -36.25, 18.75, 38.75},
		},
		{
			id:            2,
			inputPayments: []float32{100, 72.5, 20, 7, 0},
			expected:      []float32{-60.1, -32.6, 19.9, 32.9, 39.9},
		},
	}

	for _, test := range cases {
		t.Logf("Running test case #%v\n", test.id)
		got := calculateDebts(test.inputPayments)
		for i := range test.expected {
			if !isZero(got[i] - test.expected[i]) {
				t.Errorf("Test case #%v failed: Expected: %v, Got: %v", test.id, test.expected, got)
				continue
			}
		}
	}
}

type TxListTestCase struct {
	id           int
	inputDebts   []float32
	inputFriends []string
	expected     []string
}

func TestDebtValidator(t *testing.T) {
	t.Logf("Running debt validator test #1")
	case1 := []float32{30, 20.43, -100, 30, -23.77}
	if err := validateDebts(case1); err == nil {
		t.Errorf("Test case #1: Expected an error to be returned but got none for debt slice: %v", case1)
	}

	case2 := []float32{-21.25, -36.25, 18.75, 38.75}
	if err := validateDebts(case2); err != nil {
		t.Errorf("Test case #2: Did not expect error to be returned but got one for debt slice: %v", case2)
	}
}

func TestBuildTransactionList(t *testing.T) {
	cases := []TxListTestCase{
		{
			id:           1,
			inputDebts:   []float32{-21.25, -36.25, 18.75, 38.75},
			inputFriends: []string{"eve", "mitch", "bethie", "randy"},
			expected: []string{
				"randy pays mitch $36.25",
				"randy pays eve $2.50",
				"bethie pays eve $18.75",
			},
		},
		// this case demonstrate how the clearEasyMatches logic is an important optimization (optimize = guarantee we find the solution with the smallest possible # of transactions. Actual performance will probably take a hit.)
		{
			id:           2,
			inputDebts:   []float32{-50, -20, -10, 20, 60}, // if we were to pay down from the ends without considering the rest of the debts, we would miss the matching $20 debts (an easy win), leading to an extra transaction in the end.
			inputFriends: []string{"mal", "tessa", "sam", "scarlett", "mary"},
			expected: []string{
				"scarlett pays tessa $20.00",
				"mary pays mal $50.00",
				"mary pays sam $10.00",
			},
		},
	}

	for _, test := range cases {
		t.Logf("Running testcase #%v\n", test.id)
		got := buildTransactionList(test.inputDebts, test.inputFriends)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Test case #%v failed: Expected: %v, Got: %v", test.id, test.expected, got)
		}
	}
}
