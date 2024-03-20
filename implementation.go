package Lab_2

import (
	"strings"
)

func isOperator(c string) bool {
	return strings.ContainsAny(c, "+-*/^")
}

func reverseSlice(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func PrefixToPostfix(prefixExpr string) string {
	expr := reverseSlice(strings.Split(prefixExpr, " "))
	var stack []string

	for _, token := range expr {
		if isOperator(token) {
			if len(stack) < 2 {
				return "Error! Invalid expression!"
			}
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, op1+" "+op2+" "+token)
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "Error! Invalid expression!"
	}

	return stack[0]
}
