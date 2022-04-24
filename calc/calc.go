package calc

import (
	"fmt"
	"math"
	"strconv"
)

type operation = string

const (
	EXPONENT       = operation("^")
	MULTIPLICATION = operation("*")
	DIVISION       = operation("/")
	ADDITION       = operation("+")
	SUBTRACTION    = operation("-")
)

const (
	ExponentPriority       = 1
	MultiplicationPriority = 2
	DivisionPriority       = 2
	AdditionPriority       = 3
	SubtractionPriority    = 3
)

var numRunes = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var opRunes = []rune{'*', '/', '+', '-', '^'}

type calcTree struct {
	op           operation
	leftVal      float64
	rightVal     float64
	leftSubCalc  *calcTree
	rightSubCalc *calcTree
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

	return 0, ""
}

func getOpPriority(op operation) int {
	opPriorityMap := map[operation]int{
		EXPONENT:       ExponentPriority,
		MULTIPLICATION: MultiplicationPriority,
		DIVISION:       DivisionPriority,
		ADDITION:       AdditionPriority,
		SUBTRACTION:    SubtractionPriority,
	}
	return opPriorityMap[op]
}

func doOp(left, right float64, op operation) float64 {
	var val float64
	switch op {
	case MULTIPLICATION:
		val = left * right
		break
	case DIVISION:
		val = left / right
		break
	case ADDITION:
		val = left + right
		break
	case SUBTRACTION:
		val = left - right
		break
	case EXPONENT:
		val = math.Pow(left, right)
		break
	default:
		val = 0.0
	}
	return val
}

func (c *calcTree) calc() float64 {
	var leftVal, rightVal float64

	if c.leftSubCalc != nil {
		leftVal = c.leftSubCalc.calc()
	} else {
		leftVal = c.leftVal
	}

	if c.rightSubCalc != nil {
		rightVal = c.rightSubCalc.calc()
	} else {
		rightVal = c.rightVal
	}

	return doOp(leftVal, rightVal, c.op)
}

func (c *calcTree) stringVal() string {
	var lString, rString string

	if c.leftSubCalc != nil {
		lString = c.leftSubCalc.stringVal()
	} else {
		lString = fmt.Sprintf("%f", c.leftVal)
	}

	if c.rightSubCalc != nil {
		rString = c.rightSubCalc.stringVal()
	} else {
		rString = fmt.Sprintf("%f", c.rightVal)
	}

	return fmt.Sprintf("(%s %s %s)", lString, c.op, rString)
}

func (c *calcTree) print() {
	fmt.Println(c.stringVal())
}

func (c *calcTree) getRightest() *calcTree {
	head := c
	for head.rightSubCalc != nil {
		head = head.rightSubCalc
	}
	return head
}

func parseCalcTree(calc string) *calcTree {
	var num float64
	var head, curr, prev *calcTree
	currNumString := ""

	for _, r := range calc {
		t, o := validRune(r)

		if t == 0 {
			continue
		}

		if t == 1 {
			currNumString += string(r)
		}

		if t == 2 {
			num, _ = strconv.ParseFloat(currNumString, 64)
			currNumString = ""

			if head == nil {
				head = &calcTree{op: o, leftVal: num}
				continue
			} else {
				head.getRightest().rightVal = num
			}

			curr = head
			for curr != nil && getOpPriority(o) < getOpPriority(curr.op) {
				curr, prev = curr.rightSubCalc, curr
			}
			newCalcTree := &calcTree{
				op:          o,
				leftSubCalc: curr,
			}
			if prev == nil {
				head = newCalcTree
			} else {
				newCalcTree.leftVal = prev.rightVal
				prev.rightSubCalc = newCalcTree
			}
			curr, prev = nil, nil
		}
	}

	num, _ = strconv.ParseFloat(currNumString, 64)
	head.getRightest().rightVal = num

	return head
}

func Calculate(calc string) float64 {
	t := parseCalcTree(calc)
	r := t.calc()
	return r
}
