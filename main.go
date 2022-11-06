package main

import (
	"fmt"
	"strings"
)

func SliceToString(slice []string) string {
	ret := ""
	sep := ""
	for x := range slice {
		ret = ret + sep + slice[x]
		sep = " "
	}
	return ret
}

func ToPrefixSingleLogic(exp string) (string, bool) {
	s := strings.Split(exp, " ")
	for x := range s {
		switch s[x] {
		case "and":
			fmt.Println(exp, "=> ", "math.Min( "+SliceToString(s[:x])+" , "+SliceToString(s[x+1:])+" )")
			return "math.Min( " + SliceToString(s[:x]) + " , " + SliceToString(s[x+1:]) + " )", true
		case "or":
			fmt.Println(exp, "=> ", "math.Max( "+SliceToString(s[:x])+" , "+SliceToString(s[x+1:])+" )")
			return "math.Max( " + SliceToString(s[:x]) + " , " + SliceToString(s[x+1:]) + " )", true
		}
	}
	return exp, false
}
func ToPrefix(exp string) string {
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

		}
	}

	for {
		var found bool
		for key := range stack {
			var f bool
			stack[key], f = ToPrefixSingleLogic(stack[key])
			found = found || f
		}
		if !found {
			break
		}

	}
	exp = stack["base"]
	fmt.Println(exp)
	for {
		s := strings.Split(exp, " ")
		found := false
		for x := range s {
			v, ok := stack[s[x]]
			if ok {
				exp = strings.Replace(exp, s[x], v, -1)
				found = true
			}
		}
		if !found {
			break
		}
		fmt.Println(exp)
	}
	return exp
}

// Driver code
func main() {
	s := ("x and (w or (a and b)) or (x and y)")

	fmt.Println(s, ToPrefix(s))

	//math.Max(math.Min(x, math.Max(w, math.Min(a, b)), math.Min(x, y)))

}
