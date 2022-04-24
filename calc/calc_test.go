package calc

import (
	"fmt"
	"math"
	"testing"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		Equation string
		Expected float64
	}{
		{"1+1-1", 1},
		{"1+6*7-8", 35},
		{"1*2+3", 5},
		{"2+2*2", 6},
		{"1+6*5/7", 1.0 + 6.0*5.0/7.0},
		{"20+20*20", 420},
		{"1*2+3*4", 14},
		{"10^10", 10000000000},
		{"1/0", math.Inf(1)},
		{"100", 100},
		{"5+(6-3)", 8},
		{"(1-1)+(1-1)", 0},
		{"(1+2)-3", 0},
		{"1+(2+3+(4+5)+((6+7)+8)+9)+10", 55},
		{"(2+2)*2", 8},
	}

	for _, tc := range testCases {
		actual := Calculate(tc.Equation)
		if tc.Expected != actual {
			t.Errorf("incorrect calcation of %s, expected: %f, actual: %f", tc.Equation, tc.Expected, actual)
			fmt.Println(parseCalcTree(tc.Equation).stringVal())
		}
	}
}
