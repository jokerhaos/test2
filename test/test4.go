package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 用golang写一个方法,输入 4 个数字和一个期望值,通过加减乘除,算出期望值
func main() {
	input := []int{2, 2, 5, 4}
	target := 24
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
				// if err != nil {
				// 	// 如果表达式计算出错，继续下一个组合
				// 	continue
				// }
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
	// tokens = tokens[:len(tokens)-2]
	// 初始化结果为第一个数字
	result, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0
	}

	// 遍历切片中的运算符和数字，依次计算表达式
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
		case "*":
			result *= operand
		case "/":
			if operand == 0 {
				return 0
			}
			result /= operand
		}
	}
	fmt.Printf("%s ====> %#v ====> %f \r\n", expression, tokens, result)
	return result
}
