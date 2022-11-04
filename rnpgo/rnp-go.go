package rnpgo

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	BlueColorFont = "\033[1;34m%v\033[0m"
)

type RpnGoer interface {
	SetDebug()
	IsDebugOn()
	printDebug()
	SetExpression()
	GetExpression()
	GetResult()
	AppendRPNItem()
	GetRPNExpression()
	GetRPNStack()
	AppendRPNOperatorItem()
	GetLastOperatorFromStack()
	PopOperatorFromStack()
	PrintOperatorStack()
	PrintOperatorStackBeforeAfter()
	GetOperatorStackLength()
	ConvertExpressionToStack()
	ConvertToRPN()
	CheckPrecedence()
	GetIndexOfStringList()
	IsNumericString()
	IsOperator()
	SimpleCalculate()
	CalculateRPN()
	RemoveFromStackByIndex()
	ShowResult()
	CalculateExpression()
}

type RpnGo struct {
	debug            bool
	expression       string
	expression_stack []string
	operator_stack   []string

	rpn_expression string
	rpn_stack      []string

	result        float64
	result_string string
}

func (r *RpnGo) SetDebug(debug bool) {
	//r.printDebug("METHOD : SetDebug")
	r.debug = debug
}

func (r *RpnGo) IsDebugOn() bool {
	return r.debug
}

func (r *RpnGo) printDebug(i interface{}) {

	if r.IsDebugOn() {

		switch v := i.(type) {
		case string:
			fmt.Printf(BlueColorFont, v)
			fmt.Println()
		case int:
			fmt.Printf(BlueColorFont, v)
			fmt.Println()
		case float64:
			fmt.Printf(BlueColorFont, v)
			fmt.Println()
		default:
			s, _ := json.MarshalIndent(i, "", "\t")
			fmt.Println(string(s))
		}

	}
}

func (r *RpnGo) SetExpression(expression string) {
	//r.printDebug("METHOD : SetExpression")
	expression = strings.ReplaceAll(expression, " ", "")
	expression = strings.ReplaceAll(expression, " ", "")
	expression = "(" + expression + ")"
	expression = strings.TrimSpace(expression)
	r.expression = expression
}

func (r *RpnGo) GetExpression() string {
	//r.printDebug("METHOD : GetExpression")
	return r.expression
}

func (r *RpnGo) GetResult() float64 {
	//r.printDebug("METHOD : GetExpression")
	return r.result
}

func (r *RpnGo) AppendRPNItem(item string) {
	//r.printDebug("METHOD : AppendRPNItem")
	if r.IsDebugOn() {
		r.printDebug("AppendRPNItem: [  " + item + "  ]")
	}
	if item != "(" && item != ")" {
		r.rpn_expression = r.rpn_expression + item + " "
		r.rpn_stack = append(r.rpn_stack, item)
	}
	//r.printDebug(r.rpn_expression)
}

func (r *RpnGo) GetRPNExpression() string {
	return r.rpn_expression
}
func (r *RpnGo) GetRPNStack() []string {
	return r.rpn_stack
}

func (r *RpnGo) AppendRPNOperatorItem(item string) {
	//r.printDebug("METHOD : AppendRPNOperatorItem")
	if r.IsDebugOn() {
		r.printDebug("AppendRPNOperatorItem: [  " + item + "  ]")
	}
	r.operator_stack = append(r.operator_stack, item)

	r.PrintOperatorStack()
}

func (r *RpnGo) GetLastOperatorFromStack() string {
	//r.printDebug("METHOD : GetLastOperatorFromStack")
	if len(r.operator_stack) > 0 {
		//r.printDebug(r.operator_stack[len(r.operator_stack)-1])
		return r.operator_stack[len(r.operator_stack)-1]
	}
	return ""

}

func (r *RpnGo) PopOperatorFromStack() []string {
	//r.printDebug("METHOD : PopOperatorFromStack")
	if r.IsDebugOn() {
		r.printDebug("PopOperatorFromStack")
		aux_op_stack := r.operator_stack
		if len(aux_op_stack) > 0 {
			aux_op_stack = aux_op_stack[:len(aux_op_stack)-1]
		}
		r.PrintOperatorStackBeforeAfter(r.operator_stack, aux_op_stack)
	}
	if len(r.operator_stack) > 0 {
		r.operator_stack = r.operator_stack[:len(r.operator_stack)-1]
	}

	return r.operator_stack
}

func (r *RpnGo) PrintOperatorStack() {

	if r.IsDebugOn() {
		r.printDebug("Operator Stack :")
		if len(r.operator_stack) > 0 {
			fmt.Println("     ")
			fmt.Println("|     |")
			for i := len(r.operator_stack) - 1; i >= 0; i-- {
				item := r.operator_stack[i]
				//fmt.Println(item)
				fmt.Println("|  " + item + "  |")
			}
			fmt.Println("|_____|")
			fmt.Println("       ")
		}
	}
}

func (r *RpnGo) PrintOperatorStackBeforeAfter(before []string, after []string) {

	var aux_debug_before []string
	var aux_debug_after []string

	max_len := len(before)
	if len(after) > max_len {
		max_len = len(after)
	}

	if r.IsDebugOn() {
		if len(before) > 0 {
			for i := len(before) - 1; i >= 0; i-- {
				item := before[i]
				aux_debug_before = append(aux_debug_before, " |  "+item+"  |")
			}
		}

		if len(after) > 0 {
			for i := len(after) - 1; i >= 0; i-- {
				item := after[i]
				aux_debug_after = append(aux_debug_after, "|  "+item+"  |")
			}
		}

		fmt.Println("                    ")
		fmt.Println(" |     |     |     |")
		for i := 0; i < max_len-1; i++ {
			aux_before := ""
			aux_after := ""

			if i < len(before) {
				aux_before = aux_debug_before[i]
			} else {
				aux_before = " |     |"
			}
			if i < len(after) {
				aux_after = aux_debug_after[i]
			} else {
				aux_after = "|     |"
			}
			fmt.Println(aux_before + "     " + aux_after)

		}

		if len(aux_debug_before) == 0 && len(aux_debug_after) == 0 {
			fmt.Println(" |EMPTY|     |EMPTY|")
		}

		fmt.Println(" |_____|     |_____|")
		fmt.Println("  BEFORE      AFTER ")
		/* fmt.Println(len(aux_debug_before))
		fmt.Println(len(aux_debug_after)) */
	}
}

func (r *RpnGo) GetOperatorStackLength() int {
	return len(r.operator_stack)
}

func (r *RpnGo) ConvertExpressionToStack() []string {
	//r.printDebug("METHOD : ConvertExpressionToStack")
	expression := r.expression
	var list []string
	tempStr := ""
	isLastCharNumeric := false

	//walk through expression, check every char
	for i := 0; i < len(expression); i++ {
		//get byte as char
		tempChar := fmt.Sprintf("%c", expression[i])

		//r.printDebug(tempChar)
		//check if char is numeric or dot "." / doest check if is "e"
		if r.IsNumericString(tempChar) {
			//if previous char is numeric OR list is empty
			if isLastCharNumeric || len(list) == 0 {
				tempStr = tempStr + tempChar
			} else {
				tempStr = tempStr + tempChar
			}
			isLastCharNumeric = true
		} else {
			if isLastCharNumeric {
				//add number to list
				list = append(list, tempStr)
			}

			tempStr = ""
			//add char to list
			list = append(list, tempChar)

			//set "previous char is numeric" as false
			isLastCharNumeric = false

		}

		//if is the last char of string
		if i == (len(expression) - 1) {
			//check if it is numeric
			if r.IsNumericString(tempChar) {
				//add number to list
				list = append(list, tempStr)
			} else {
				//add char to list
				list = append(list, tempChar)
			}

		}

	}

	/* for i := range list {
		item := list[i]
		r.printDebug(item)
	} */

	r.expression_stack = list
	return list

}

func (r *RpnGo) ConvertToRPN() string {
	//r.printDebug("METHOD : ConvertToRPN")

	/*
		Walk through list, add number to rpn string,
			add operator to stack,
			check operator precedence
	*/

	expression_list := r.expression_stack
	//operator_precedence := []string{"=", "-", "+", "/", "*", "^"}

	first_i := true

	for i := range expression_list {
		item := expression_list[i]

		r.printDebug("")
		r.printDebug("")
		r.printDebug("==========================================================")
		r.printDebug("============= Now on item: [  " + item + "  ] ")
		r.printDebug("==========================================================")
		r.printDebug("")
		r.printDebug("RPN Expression: " + r.rpn_expression)
		r.printDebug("")
		r.PrintOperatorStack()
		r.printDebug("")
		r.printDebug("-------------------------")
		r.printDebug("")

		if r.IsOperator(item) {
			//if stack of operator is empty, just add operator to stack

			if r.GetOperatorStackLength() == 0 || first_i {
				first_i = false
				r.AppendRPNOperatorItem(item)

			} else {
				//check operator precedence, if actual operator is greater than last operator of stack, add actual operator to stack
				//check operator precedente, while actual operator is equal or  lower than last operator of stack, add last operator of stack to rpn_expression, should use while() go -> for
				//------------------------------------------------------------------------------------ add actual operator to stack
				// important, when hit left bracket, just ignore
				// important, when hit right bracket, remove everything of stack (until hit left bracked in stack), then add to rpn_expression

				/* for item == ")" && r.GetLastOperatorFromStack() != "(" {
					r.printDebug("POP, last operator is '('")
					r.PopOperatorFromStack()
				} */

				//When item is "(", should add to stack, and go to next item"
				if item == "(" || item == " " {
					r.AppendRPNOperatorItem(item)

					continue
				}

				//WHEN HIT A ")", POP AS STACKS UNTIL FIND A "("
				if r.GetOperatorStackLength() > 0 && item == ")" {

					for r.GetOperatorStackLength() > 0 && r.GetLastOperatorFromStack() != "(" {
						r.AppendRPNItem(r.GetLastOperatorFromStack())

						//pop from stack
						r.PopOperatorFromStack()
						//r.GetLastOperatorFromStack()

					}
					//WHEN FIND A "(" POP IT, AND GO TO NEXT CHAR OF EXPRESSION
					if r.GetOperatorStackLength() > 0 && r.GetLastOperatorFromStack() == "(" {
						r.PopOperatorFromStack()
					}
					continue
				}

				poped_loop := false
				//check operator precedente, while actual operator is equal or  lower than last operator of stack, add last operator of stack to rpn_expression, should use while() go -> for
				for r.GetOperatorStackLength() > 0 && (r.CheckPrecedence(item) <= r.CheckPrecedence(r.GetLastOperatorFromStack())) {
					r.AppendRPNItem(r.GetLastOperatorFromStack())

					//pop from stack
					r.PopOperatorFromStack()

					poped_loop = true
				}

				if poped_loop {
					r.AppendRPNOperatorItem(item)

					poped_loop = false
				} else if r.GetOperatorStackLength() > 0 && (r.CheckPrecedence(item) > r.CheckPrecedence(r.GetLastOperatorFromStack())) {
					//check operator precedence, if actual operator is bigger than last operator of stack, add actual operator to stack
					r.AppendRPNOperatorItem(item)

				}
			}

		} else {
			//if its not operator, add item to rpn expression
			r.AppendRPNItem(item)

		}
	}

	for r.GetOperatorStackLength() > 0 {

		r.AppendRPNItem(r.GetLastOperatorFromStack())
		//pop from stack
		r.PopOperatorFromStack()

	}

	/* rpn_expression = strings.ReplaceAll(rpn_expression, "(", "")
	rpn_expression = strings.ReplaceAll(rpn_expression, ")", "")
	r.printDebug(rpn_expression) */

	//r.rpn_expression = rpn_expression

	r.printDebug("==========================================================")
	r.printDebug("======================   END   ===========================")
	r.printDebug("==========================================================")
	r.printDebug("")
	r.printDebug("RPN Expression: " + r.rpn_expression)
	r.printDebug("")
	r.PrintOperatorStack()
	r.printDebug("")
	r.printDebug("=================================")

	r.rpn_expression = strings.Trim(r.rpn_expression, " ")
	r.rpn_expression = strings.TrimRight(r.rpn_expression, " ")
	return r.rpn_expression
}

func (r *RpnGo) CheckPrecedence(item string) int {
	switch item {
	case "^":
		return 40
	case "*":
		return 30
	case "/":
		return 30
	case "+":
		return 20
	case "-":
		return 20
	}
	return 0
}

func (r *RpnGo) GetIndexOfStringList(stringList []string, search string) int {
	//r.printDebug("METHOD : GetIndexOfStringList")
	for i := 0; i < len(stringList); i++ {
		if stringList[i] == search {
			return i
		}
	}
	return -1
}

func (r *RpnGo) IsNumericString(value string) bool {
	//r.printDebug("METHOD : IsNumericString")
	if value == "0" || value == "1" || value == "2" || value == "3" || value == "4" || value == "5" || value == "6" || value == "7" || value == "8" || value == "9" || value == "." {
		return true
	}
	return false
}

func (r *RpnGo) IsOperator(value string) bool {
	//r.printDebug("METHOD : IsOperator")
	if value == "^" || value == "*" || value == "/" || value == "+" || value == "-" || value == "=" || value == ")" || value == "(" {
		return true
	}
	return false
}

func (r *RpnGo) SimpleCalculate(value1 float64, value2 float64, operator string) float64 {
	aux_result := 0.0

	switch operator {
	case "^":
		aux_result = math.Pow(value1, value2)
	case "*":
		aux_result = value1 * value2
	case "/":
		aux_result = value1 / value2
	case "+":
		aux_result = value1 + value2
	case "-":
		aux_result = value1 - value2
	}
	if r.IsDebugOn() {
		aux_debug_string := fmt.Sprintf("Calculate: %f %s %f = %f", value1, operator, value2, aux_result)
		r.printDebug(aux_debug_string)
	}
	return aux_result
}

func (r *RpnGo) CalculateRPN() float64 {

	r.printDebug("")
	r.printDebug("")
	r.printDebug("==========================================================")
	r.printDebug("START - CALCULATE RPN")
	r.printDebug("==========================================================")
	r.printDebug("")
	if r.IsDebugOn() {
		fmt.Println(r.rpn_stack)
	}
	r.printDebug("")
	aux_stack := r.rpn_stack
	for len(aux_stack) > 1 {
		for i := 0; i < len(aux_stack); i++ {
			item := aux_stack[i]
			r.printDebug("-----------> Actual item: " + item)
			if r.IsOperator(item) {
				//if add factorial, this block should get only 1 value before actual index

				r.printDebug("isOperator: " + item)
				if r.IsDebugOn() {
					fmt.Println(aux_stack)
				}

				//Get item of index-2, to get value2
				value1, err := strconv.ParseFloat(aux_stack[i-2], 64)
				if err != nil {
					fmt.Printf("Error value1 as %s", aux_stack[i-2])
				}
				//Get item of index-1, to get value1
				value2, err := strconv.ParseFloat(aux_stack[i-1], 64)
				if err != nil {
					fmt.Printf("Error value1 as %s", aux_stack[i-1])
				}
				//calculate using value1, value2, and item (as operator)
				result_calc := r.SimpleCalculate(value1, value2, item)
				aux_result := fmt.Sprintf("%f", result_calc)

				//replace item (actual index), put result from calc
				aux_stack[i] = aux_result
				//remove value1 and value2 from array
				aux_stack = r.RemoveFromStackByIndex(aux_stack, i-1)
				aux_stack = r.RemoveFromStackByIndex(aux_stack, i-2)

				//set i=0 to check since the start of list
				i = 0

			} else {
				//nothing to do here, go to next item of list
				r.printDebug("isNumeric: " + item)
			}
		}
	}
	if len(aux_stack) == 1 {
		aux_result, err := strconv.ParseFloat(aux_stack[0], 64)
		if err != nil {
			fmt.Printf("Error value1 as %s", aux_stack[0])
		}
		r.result_string = aux_stack[0]
		r.result = aux_result
		return r.result
	}
	return -1
}

func (r *RpnGo) CalculateExpression(expression string) float64 {
	r.SetExpression(expression)
	r.ConvertExpressionToStack()
	r.ConvertToRPN()
	r.CalculateRPN()
	return r.result
}

func (r *RpnGo) RemoveFromStackByIndex(list []string, index int) []string {
	return append(list[:index], list[index+1:]...)
}

func (r *RpnGo) ShowResult() {
	fmt.Println("EXPRESSION: " + r.expression)
	fmt.Println("RPN EXPRESSION: " + r.rpn_expression)
	fmt.Println("RESULT: " + r.result_string)
}

func (r *RpnGo) ShowResultAsTest() {
	temp_result := fmt.Sprintf(`{expression: "%v", rpn_expression: "%v", result: %f, isResultCorrect: true}`,
		r.expression, r.rpn_expression, r.result)
	fmt.Println(temp_result)
}
