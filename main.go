package main

import (
	"fmt"
	"gocalc/calc"
)

func main() {
	cs := []string{
		"1+1-1",
		"1+6*7-8",
		"1*2+3",
		"2+2*2",
		"1+6*5/7",
		"20+20*20",
		"1*2+3*4",
		"10^10",
		"1/0",
		"100",
		"5+(6-3)",
		"(1-1)+(1-1)",
		"(1+2)-3",
		"1+(2+3+(4+5)+((6+7)+8)+9)+10",
		"(2+2)*2",
	}

	for _, c := range cs {
		t := calc.Calculate(c)
		fmt.Printf("%s = %f\n", c, t)
	}
}
