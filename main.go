package main

import (
	"fiber-teste/fuzzy"
	"fmt"

	"github.com/PaesslerAG/gval"
)

// https://github.com/PaesslerAG/gval

func main() {
	s := ("70.12 and ( 85. or ( 7. and 90. ) ) or ( 70. and 100. )")
	//s := "71. or 80.3"
	exp := fuzzy.FuzzyLogicalInference(s)
	fmt.Println(s, exp)
	//exp = "9>4 || 2==10"
	i, err := gval.Evaluate(exp, nil)
	fmt.Println(i, err)
}
