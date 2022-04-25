package calc

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

type operation = string

const (
	BRACKET        = operation("(")
	EXPONENT       = operation("^")
	MULTIPLICATION = operation("*")
	DIVISION       = operation("/")
	ADDITION       = operation("+")
	SUBTRACTION    = operation("-")
	NOOP           = operation("_")
)

const (
	BracketPriority        = 0
	ExponentPriority       = 1
	MultiplicationPriority = 2
	DivisionPriority       = 2
	AdditionPriority       = 3
	SubtractionPriority    = 3
	NoopPriority           = -1
)

var numRunes = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var opRunes = []rune{'*', '/', '+', '-', '^'}

var opPriorityMap = map[operation]int{
	BRACKET:        BracketPriority,
	EXPONENT:       ExponentPriority,
	MULTIPLICATION: MultiplicationPriority,
	DIVISION:       DivisionPriority,
	ADDITION:       AdditionPriority,
	SUBTRACTION:    SubtractionPriority,
	NOOP:           NoopPriority,
}

var opNumArgs = map[operation]int{
	EXPONENT:       2,
	MULTIPLICATION: 2,
	DIVISION:       2,
	ADDITION:       2,
	SUBTRACTION:    2,
	NOOP:           1,
}

type calcTree struct {
	op           operation
	numVals      int
	values       []float64
	subCalcTrees []*calcTree
}

func validRune(r rune) (int, operation) {
	for _, c := range numRunes {
		if r == c {
			return 1, ""
		}
	}

	for _, c := range opRunes {
		if r == c {
			return 2, operation(c)
		}
	}

	if r == '(' {
		return 3, BRACKET
	}

	return 0, ""
}

func getOpPriority(op operation) int {
	return opPriorityMap[op]
}

func getOpNumValues(op operation) int {
	return opNumArgs[op]
}

func doOp(values []float64, op operation) float64 {
	var val float64
	switch op {
	case MULTIPLICATION:
		val = values[0] * values[1]
		break
	case DIVISION:
		val = values[0] / values[1]
		break
	case ADDITION:
		val = values[0] + values[1]
		break
	case SUBTRACTION:
		val = values[0] - values[1]
		break
	case EXPONENT:
		val = math.Pow(values[0], values[1])
		break
	case NOOP:
		val = values[0]
		break
	default:
		val = 0.0
	}
	return val
}

func (c *calcTree) calc() float64 {
	values := make([]float64, c.numVals)
	strValues := make([]string, c.numVals)
	copy(values, c.values)
	for i := 0; i < c.numVals; i++ {
		if c.subCalcTrees[i] != nil {
			values[i] = c.subCalcTrees[i].calc()
			strValues[i] = fmt.Sprintf("[%f]", values[i])
		} else {
			strValues[i] = fmt.Sprintf("%f", values[i])
		}
	}

	switch c.op {
	case EXPONENT, MULTIPLICATION, DIVISION, ADDITION, SUBTRACTION:
	}

	return doOp(values, c.op)
}

func (c *calcTree) stringVal() string {
	values := make([]string, len(c.values))

	for i := 0; i < c.numVals; i++ {
		if c.subCalcTrees[i] != nil {
			values[i] = c.subCalcTrees[i].stringVal()
		} else {
			values[i] = fmt.Sprintf("%.3f", c.values[i])
		}
	}

	switch c.op {
	case EXPONENT, MULTIPLICATION, DIVISION, ADDITION, SUBTRACTION:
		return fmt.Sprintf("(%s %s %s)", values[0], c.op, values[1])
	case NOOP:
		return fmt.Sprintf("(%s)", values[0])
	}

	return ""
}

func (c *calcTree) print(w io.Writer) {
	_, _ = fmt.Fprintln(w, c.stringVal())
}

func (c *calcTree) getRightest() *calcTree {
	head := c
	for head.subCalcTrees[head.numVals-1] != nil {
		head = head.subCalcTrees[head.numVals-1]
	}
	return head
}

func parseCalcTree(calc string) *calcTree {
	var num float64
	var head, curr, prev, prevBracketed *calcTree
	waitingBracket := false
	currNumString := ""

	skipUntil := 0

	for i, r := range calc {
		if i < skipUntil {
			continue
		}

		t, o := validRune(r)
		thisPriority := getOpPriority(o)

		if t == 0 {
			continue
		}

		if t == 1 {
			currNumString += string(r)
			continue
		}

		if t == 2 {
			num, _ = strconv.ParseFloat(currNumString, 64)
			currNumString = ""

			opNumArgs := getOpNumValues(o)

			newCalcTree := &calcTree{
				op:           o,
				numVals:      opNumArgs,
				values:       make([]float64, opNumArgs),
				subCalcTrees: make([]*calcTree, opNumArgs),
			}

			if head == nil {
				head = newCalcTree

				if waitingBracket {
					waitingBracket = false
					head.subCalcTrees[0] = prevBracketed
				} else {
					head.values[0] = num
				}
				continue
			} else {
				rightest := head.getRightest()

				if waitingBracket {
					waitingBracket = false
					rightest.subCalcTrees[rightest.numVals-1] = prevBracketed
				} else {
					rightest.values[rightest.numVals-1] = num
				}
			}

			curr = head
			for curr != nil && thisPriority < getOpPriority(curr.op) {
				curr, prev = curr.subCalcTrees[curr.numVals-1], curr
			}
			newCalcTree.subCalcTrees[0] = curr
			if prev == nil {
				head = newCalcTree
			} else {
				newCalcTree.values[0] = prev.values[prev.numVals-1]
				prev.subCalcTrees[prev.numVals-1] = newCalcTree
			}
			curr, prev = nil, nil
		}

		if t == 3 {
			level := 0
			for j := i + 1; j < len(calc); j++ {
				nr := calc[j]

				if nr == ')' && level == 0 {
					prevBracketed = parseCalcTree(calc[i+1 : j])
					skipUntil = j + 1
					waitingBracket = true
					break
				}

				if nr == '(' {
					level += 1
				} else if nr == ')' {
					level -= 1
				}
			}
		}
	}

	if waitingBracket {
		if head == nil {
			return &calcTree{
				op:           NOOP,
				numVals:      1,
				values:       []float64{0},
				subCalcTrees: []*calcTree{prevBracketed},
			}
		}

		rightest := head.getRightest()
		rightest.subCalcTrees[rightest.numVals-1] = prevBracketed

		return head
	}

	num, _ = strconv.ParseFloat(currNumString, 64)

	if head == nil {
		return &calcTree{
			op:           NOOP,
			numVals:      1,
			values:       []float64{num},
			subCalcTrees: []*calcTree{nil},
		}
	}

	rightest := head.getRightest()
	rightest.values[rightest.numVals-1] = num

	return head
}

func Calculate(calc string) float64 {
	t := parseCalcTree(calc)
	r := t.calc()
	return r
}
