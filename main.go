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
	}

	for _, c := range cs {
		t := calc.Calculate(c)
		fmt.Printf("%s = %f\n", c, t)
	}
}
