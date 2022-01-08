package evenup

import (
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
		t.Logf("Running test case #%v...\n", test.id)
		got := calculateDebts(test.inputPayments)
		for i := range test.expected {
			if !isEffectivelyZero(got[i] - test.expected[i]) {
				t.Errorf("Test case #%v failed: Expected: %v, Got: %v", test.id, test.expected, got)
				t.FailNow()
			}
		}
	}
}
