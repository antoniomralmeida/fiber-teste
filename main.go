package main

import (
	"fiber-teste/rnpgo"
)

// https://github.com/celioyutaka/rpn-go
func main() {

	expression := "100.0 and (97.7 or 81.8) "
	rpn := rnpgo.RpnGo{}
	rpn.SetDebug(false)
	rpn.CalculateExpression(expression)

	rpn.ShowResult()

}
