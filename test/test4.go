package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 用golang写一个方法，输入4个数字和一个期望值，通过加减乘除，算出期望值
func main() {
	input := []int{2, 2, 5, 4}
	target := 21
	calculate(input, target)
}

func calculate(input []int, target int) (bool, string) {
	// 定义运算符组合
	operators := []string{"+", "-", "*", "/"}
	num1, num2, num3, num4 := input[0], input[1], input[2], input[3]
	// 遍历所有可能的运算符组合
	for _, op1 := range operators {
		for _, op2 := range operators {
			for _, op3 := range operators {
				// 构建表达式
				expression := fmt.Sprintf("%d %s %d %s %d %s %d", num1, op1, num2, op2, num3, op3, num4)
				// 计算表达式结果
				result := eval(expression)
				// 判断计算结果是否等于期望值
				if result == float64(target) {
					fmt.Println(expression, "=", target)
					return true, expression
				}
			}
		}
	}

	return false, ""
}

func eval(expression string) float64 {
	// 将表达式字符串按空格分隔成切片
	tokens := strings.Split(expression, " ")
	// 优化：先计算乘法和除法
	for i := 1; i < len(tokens); i += 2 {
		operator := tokens[i]
		if operator == "*" || operator == "/" {
			operand1, _ := strconv.ParseFloat(tokens[i-1], 64)
			operand2, _ := strconv.ParseFloat(tokens[i+1], 64)
			var result float64
			switch operator {
			case "*":
				result = operand1 * operand2
			case "/":
				if operand2 == 0 {
					return 0
				}
				result = operand1 / operand2
			}
			// 替换乘法和除法的部分为计算结果
			tokens[i-1] = strconv.FormatFloat(result, 'f', -1, 64)
			// 移除乘法和除法的部分
			tokens = append(tokens[:i], tokens[i+2:]...)
			i -= 2 // 调整索引，继续循环
		}
	}

	// 初始化结果为第一个数字
	result, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0
	}

	// 遍历剩余的运算符和数字，依次计算表达式
	for i := 1; i < len(tokens); i += 2 {
		operator := tokens[i]
		operand, err := strconv.ParseFloat(tokens[i+1], 64)
		if err != nil {
			return 0
		}
		// 根据运算符进行相应的计算
		switch operator {
		case "+":
			result += operand
		case "-":
			result -= operand
		}
	}

	fmt.Printf("%s ====> %#v ====> %f \r\n", expression, tokens, result)
	return result
}
