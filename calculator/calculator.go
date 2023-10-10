package calculator

import (
	"fmt"
)

func Calculator() {

	var operation int

	fmt.Println("Welcome to the calculator!")
	fmt.Println("Enter two numbers: ")
	firstNumber, secondNumber := getNumbers()
	for operation != 6 {
		fmt.Println("Pick an operation: ")
		fmt.Println("1) Add")
		fmt.Println("2) Subtract")
		fmt.Println("3) Multiply")
		fmt.Println("4) Divide")
		fmt.Println("5) New Numbers")
		fmt.Println("6) Exit")
		fmt.Scanln(&operation)
		userInput(operation, firstNumber, secondNumber)
	}
}

func getNumbers() (int, int) {
	var firstNumber int
	var secondNumber int
	fmt.Println("Please enter the first number: ")
	fmt.Scanln(&firstNumber)
	fmt.Println("Please enter the second number: ")
	fmt.Scanln(&secondNumber)
	return firstNumber, secondNumber
}

func userInput(operation int, firstNumber int, secondNumber int) (int, int) {
	switch operation {
	case 1:
		result := Add(firstNumber, secondNumber)
		fmt.Println(fmt.Sprintf("%d", firstNumber) + "+" +
			fmt.Sprintf("%d", secondNumber) + "=" +
			fmt.Sprintf("%d", result))
	case 2:
		result := Subtract(firstNumber, secondNumber)
		fmt.Println(fmt.Sprintf("%d", firstNumber) + "-" +
			fmt.Sprintf("%d", secondNumber) + "=" +
			fmt.Sprintf("%d", result))
	case 3:
		result := Multiply(firstNumber, secondNumber)
		fmt.Println(fmt.Sprintf("%d", firstNumber) + "*" +
			fmt.Sprintf("%d", secondNumber) + "=" +
			fmt.Sprintf("%d", result))
	case 4:
		result := Divide(firstNumber, secondNumber)
		fmt.Println(fmt.Sprintf("%d", firstNumber) + "/" +
			fmt.Sprintf("%d", secondNumber) + "=" +
			fmt.Sprintf("%d", result))
	case 5:
		firstNumber, secondNumber = getNumbers()
	case 6:
		fmt.Println("Goodbye!")
	default:
		fmt.Println("Invalid operation")
	}
	return firstNumber, secondNumber
}

func Divide(firstNumber int, secondNumber int) int {
	return firstNumber / secondNumber
}

func Multiply(firstNumber int, secondNumber int) int {
	return firstNumber * secondNumber
}

func Subtract(firstNumber int, secondNumber int) int {
	return firstNumber - secondNumber
}

func Add(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}
