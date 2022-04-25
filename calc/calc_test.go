package calc

import (
	"bytes"
	"fmt"
	"math"
	"testing"
)

func TestCalculate(t *testing.T) {
	tt := []struct {
		Equation string
		Expected float64
	}{
		{
			Equation: "1+1-1",
			Expected: 1,
		},
		{
			Equation: "1+6*7-8",
			Expected: 35,
		},
		{
			Equation: "1*2+3",
			Expected: 5,
		},
		{
			Equation: "2+2*2",
			Expected: 6,
		},
		{
			Equation: "1+6*5/7",
			Expected: 1.0 + 6.0*5.0/7.0,
		},
		{
			Equation: "20+20*20",
			Expected: 420,
		},
		{
			Equation: "1*2+3*4",
			Expected: 14,
		},
		{
			Equation: "10^10",
			Expected: 10000000000,
		},
		{
			Equation: "1/0",
			Expected: math.Inf(1),
		},
		{
			Equation: "100",
			Expected: 100,
		},
		{
			Equation: "5+(6-3)",
			Expected: 8,
		},
		{
			Equation: "(1-1)+(1-1)",
			Expected: 0,
		},
		{
			Equation: "(1+2)-3",
			Expected: 0,
		},
		{
			Equation: "1+(2+3+(4+5)+((6+7)+8)+9)+10",
			Expected: 55,
		},
		{
			Equation: "(2+2)*2",
			Expected: 8,
		},
		{
			Equation: "1 + 1 - 1",
			Expected: 1,
		},
		{
			Equation: "(((5+5)))",
			Expected: 10,
		},
		{
			Equation: "(100)",
			Expected: 100,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Equation, func(t *testing.T) {
			actual := Calculate(tc.Equation)
			if tc.Expected != actual {
				t.Errorf("incorrect calcation of %s, expected: %f, actual: %f", tc.Equation, tc.Expected, actual)
				fmt.Println(parseCalcTree(tc.Equation).stringVal())
			}
		})
	}
}

func TestCalcTree_Print(t *testing.T) {
	tt := []struct {
		Equation string
		Expected string
	}{
		{
			Equation: "1+1-1",
			Expected: "((1.000 + 1.000) - 1.000)\n",
		},
		{
			Equation: "1+6*7-8",
			Expected: "((1.000 + (6.000 * 7.000)) - 8.000)\n",
		},
		{
			Equation: "1*2+3",
			Expected: "((1.000 * 2.000) + 3.000)\n",
		},
		{
			Equation: "2+2*2",
			Expected: "(2.000 + (2.000 * 2.000))\n",
		},
		{
			Equation: "1+6*5/7",
			Expected: "(1.000 + ((6.000 * 5.000) / 7.000))\n",
		},
		{
			Equation: "20+20*20",
			Expected: "(20.000 + (20.000 * 20.000))\n",
		},
		{
			Equation: "1*2+3*4",
			Expected: "((1.000 * 2.000) + (3.000 * 4.000))\n",
		},
		{
			Equation: "10^10",
			Expected: "(10.000 ^ 10.000)\n",
		},
		{
			Equation: "1/0",
			Expected: "(1.000 / 0.000)\n",
		},
		{
			Equation: "100",
			Expected: "(100.000)\n",
		},
		{
			Equation: "5+(6-3)",
			Expected: "(5.000 + (6.000 - 3.000))\n",
		},
		{
			Equation: "(1-1)+(1-1)",
			Expected: "((1.000 - 1.000) + (1.000 - 1.000))\n",
		},
		{
			Equation: "(1+2)-3",
			Expected: "((1.000 + 2.000) - 3.000)\n",
		},
		{
			Equation: "1+(2+3+(4+5)+((6+7)+8)+9)+10",
			Expected: "((1.000 + ((((2.000 + 3.000) + (4.000 + 5.000)) + ((6.000 + 7.000) + 8.000)) + 9.000)) + 10.000)\n",
		},
		{
			Equation: "(2+2)*2",
			Expected: "((2.000 + 2.000) * 2.000)\n",
		},
		{
			Equation: "1 + 1 - 1",
			Expected: "((1.000 + 1.000) - 1.000)\n",
		},
		{
			Equation: "(((5+5)))",
			Expected: "((((5.000 + 5.000))))\n",
		},
		{
			Equation: "(100)",
			Expected: "((100.000))\n",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Equation, func(t *testing.T) {
			var actual bytes.Buffer
			ct := parseCalcTree(tc.Equation)
			ct.print(&actual)
			if tc.Expected != actual.String() {
				t.Errorf("incorrect print of %s,\n\texpected: %s\n\t  actual: %s", tc.Equation, tc.Expected, actual.String())
			}
		})
	}
}

func TestCalcTree_stringValInvalidOp(t *testing.T) {
	tree := &calcTree{
		op:           operation("#"),
		numVals:      1,
		values:       []float64{0.0},
		subCalcTrees: []*calcTree{nil},
	}

	actual := tree.stringVal()

	if actual != "" {
		t.Errorf("unexpected output")
	}
}

func Test_doOpInvalidOp(t *testing.T) {
	expected := 0.0
	actual := doOp([]float64{1.0}, operation("#"))

	if expected != actual {
		t.Errorf("unexpected output")
	}
}

func BenchmarkCalculate(b *testing.B) {
	bt := []struct {
		Equation string
	}{
		{
			Equation: "1+1-1",
		},
		{
			Equation: "1+6*7-8",
		},
		{
			Equation: "1*2+3",
		},
		{
			Equation: "2+2*2",
		},
		{
			Equation: "1+6*5/7",
		},
		{
			Equation: "20+20*20",
		},
		{
			Equation: "1*2+3*4",
		},
		{
			Equation: "10^10",
		},
		{
			Equation: "1/0",
		},
		{
			Equation: "100",
		},
		{
			Equation: "5+(6-3)",
		},
		{
			Equation: "(1-1)+(1-1)",
		},
		{
			Equation: "(1+2)-3",
		},
		{
			Equation: "1+(2+3+(4+5)+((6+7)+8)+9)+10",
		},
		{
			Equation: "(2+2)*2",
		},
		{
			Equation: "1 + 1 - 1",
		},
		{
			Equation: "(((5+5)))",
		},
		{
			Equation: "(100)",
		},
	}

	for _, bc := range bt {
		b.Run(bc.Equation, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Calculate(bc.Equation)
			}
		})
	}
}
