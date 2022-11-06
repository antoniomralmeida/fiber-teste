package main

import (
	"fmt"
	"strings"

	"github.com/PaesslerAG/gval"
)

func SliceToString(slice []string) string {
	ret := ""
	for x := range slice {
		if slice[x] == "and" || slice[x] == "or" {
			ret = ret + " " + slice[x] + " "
		} else {
			ret = ret + slice[x]
		}
	}
	return ret
}

func ToPrefixSingleLogic(exp string) (string, bool) {
	s := strings.Split(exp, " ")
	for x := range s {
		switch s[x] {
		case "or":
			s[x] = s[x-1]
			s[x-1] = "(" + s[x-1] + ">" + s[x+1] + "?"
			s[x+1] = ":" + s[x+1] + ")"
			return SliceToString(s), true
		}
	}
	for x := range s {
		switch s[x] {
		case "and":
			s[x] = s[x-1]
			s[x-1] = "(" + s[x-1] + "<" + s[x+1] + "?"
			s[x+1] = ":" + s[x+1] + ")"
			return SliceToString(s), true
		}
	}
	return exp, false
}

func ToFuzzyExpression(exp string) string {
	stack := make(map[string]string)
	i_stack := 1000
	key := "base"
	stack[key] = exp

oulter:
	for {
		i := -1
		j := -1
		np := 0
	inter:
		for x := range stack[key] {
			switch stack[key][x] {
			case '(':
				np++
				if np == 1 {
					i = x
				}
			case ')':
				np--
				if np == 0 {
					j = x
					break inter
				}
			}
		}
		if i != -1 {
			i_stack++
			k := fmt.Sprintf("zz%v", i_stack)
			stack[k] = stack[key][i+1 : j]
			stack[key] = strings.Replace(stack[key], "("+stack[k]+")", k, 1)
			if !strings.Contains(stack[key], "(") {
				todo := false
			todo:
				for k := range stack {
					if strings.Contains(stack[k], "(") {
						key = k
						todo = true
						break todo
					}
				}
				if !todo {
					break oulter
				}
			}

		} else {
			break
		}
	}
	keys := []string{}
	for key := range stack {
		keys = append(keys, key)
	}

	for {
		var found bool
		for i := len(keys) - 1; i >= 0; i-- {
			var f bool
			stack[keys[i]], f = ToPrefixSingleLogic(stack[keys[i]])
			found = found || f
		}
		if !found {
			break
		}
	}

	exp = stack["base"]
	for {
		found := false
		for i := len(keys) - 1; i > 0; i-- {
			if strings.Contains(exp, keys[i]) {
				exp = strings.Replace(exp, keys[i], stack[keys[i]], -1)
				found = true
			}
		}
		if !found {
			break
		}
	}
	return exp
}

// https://github.com/PaesslerAG/gval
// Driver code
func main() {
	s := ("70.12 and ( 85. or ( 7. and 90. ) ) or ( 70. and 100. )")
	//s := "71. or 80.3"
	exp := ToFuzzyExpression(s)
	fmt.Println(s, exp)
	//exp = "9>4 || 2==10"
	i, err := gval.Evaluate(exp, nil)
	fmt.Println(i, err)
}
