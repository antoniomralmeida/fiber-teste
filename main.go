package main

import (
	"fmt"
)

type No struct {
	token       string
	left, right *No
}

func rnp(tokens []string) *No {
	var tree *No

	for _, x := range tokens {
		switch x {
		case "and":
			n := No{token: "min", left: tree}
			tree = &n
		case "or":
			n := No{token: "max", left: tree}
			tree = &n
		case "(":
		case ")":
		default:
			n := No{token: x}
			if tree == nil {
				tree = &n
			} else if tree.left == nil {
				tree.left = &n
			} else {
				tree.right = &n
			}
		}
	}
	return tree
}

func ExpRnp(n *No) string {
	if n == nil {
		return ""
	} else {
		return n.token + "(" + ExpRnp(n.left) + "," + ExpRnp(n.right) + ")"
	}
}

/*
Crie uma pilha
para cada caractere 't' no fluxo de entrada {
se (t é um operando)
saída t
mais se (t é um parênteses certo){
Tokens POP e de saída até que um parêntese esquerdo seja estourado (mas não faça saída)
}
outra coisa {
Tokens POP e de saída até que um de menor prioridade do que t seja encontrado ou um parênteses esquerdos seja encontrado
ou a pilha está vazia
PUSH t
}
Fichas pop e de saída até que a pilha esteja vazia.
*/

// go get -u github.com/vinshop/exp-tree
func main() {
	t := rnp([]string{"x1", "and", "(", "x2", "or", "x3", ")"})
	fmt.PrintLn(ExpRnp(t))
}
