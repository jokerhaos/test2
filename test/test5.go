// package main

// import (
// 	"fmt"
// )

// func calculate(expression string) int {
// 	stack := []int{}
// 	operators := []rune{}
// 	num := 0

// 	for _, ch := range expression {
// 		if ch >= '0' && ch <= '9' {
// 			// 如果遇到数字字符，将其转换成整数并累积到 num
// 			num = num*10 + int(ch-'0')
// 		} else {
// 			// 如果遇到运算符，先将当前累积的数字压入栈
// 			stack = append(stack, num)
// 			num = 0

// 			// 处理栈中已有的运算符
// 			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(ch) {
// 				applyOperator(&stack, &operators)
// 			}

// 			// 将当前运算符压入栈
// 			operators = append(operators, ch)
// 		}
// 	}

// 	// 处理剩余的运算符
// 	stack = append(stack, num)
// 	for len(operators) > 0 {
// 		applyOperator(&stack, &operators)
// 	}

// 	// 栈顶元素即为最终结果
// 	return stack[0]
// }

// func precedence(operator rune) int {
// 	switch operator {
// 	case '+', '-':
// 		return 1
// 	case '*':
// 		return 2
// 	default:
// 		return 0
// 	}
// }

// func applyOperator(stack *[]int, operators *[]rune) {
// 	operator := (*operators)[len(*operators)-1]
// 	*operators = (*operators)[:len(*operators)-1]

// 	right := (*stack)[len(*stack)-1]
// 	*stack = (*stack)[:len(*stack)-1]

// 	left := (*stack)[len(*stack)-1]
// 	*stack = (*stack)[:len(*stack)-1]

// 	switch operator {
// 	case '+':
// 		*stack = append(*stack, left+right)
// 	case '-':
// 		*stack = append(*stack, left-right)
// 	case '*':
// 		*stack = append(*stack, left*right)
// 	}
// }

// func init() {
// 	expression := "12+4-1*12+"
// 	result := calculate(expression)
// 	fmt.Println("Expression:", expression)
// 	fmt.Println("Result:", result)
// }
