package main

import "fmt"

func main() {
	var number1, number2 float64
	var operator string

	fmt.Println("Welcome To CLI calculator")
	fmt.Println("Input Only Numbers")

	fmt.Print("Enter the first number : ")
	fmt.Scanln(&number1)

	fmt.Print("Enter the second number : ")
	fmt.Scanln(&number2)

	fmt.Print("Enter the operator ( + - * /) : ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%v %s %v = %v", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%v %s %v = %v", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%v %s %v = %v", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0 {
			fmt.Println("Divid  by Zero Situation ")
		} else {
			fmt.Printf("%v %s %v = %v", number1, operator, number2, number1/number2)

		}
	default:
		fmt.Println("Invalid operator")
	}
	fmt.Println()
}
