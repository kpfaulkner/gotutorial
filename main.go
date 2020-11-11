package main

import "fmt"

func adder(firstNum int, secondNum int) int {

	var answer int

	answer = firstNum + secondNum

	return answer

}

func main() {

	/*
		var num1 int = 1
		var num2 int = 2

		var total = num1 + num2

		fmt.Printf("num1 %d plus num2 %d equals %d\n", num1, num2, total)
	*/

	/*
		var name string
		name = "ferret"
		fmt.Printf("name is %s\n", name)
	*/

	/*
	     fmt.Printf("enter text :\n")
	     var input string

	     fmt.Scanln(&input)
	   	fmt.Printf("you wrote %s\n", input)
	*/

	/*
		var myString string
		var num int

		// convert string to number.
		myString = "123"
		num, _ = strconv.Atoi(myString)
		fmt.Printf("string is %s : number is %d\n", myString, num)

		// FAIL convert string to number.
		myString = "123kdk" // this is NOT a number.
		num, _ = strconv.Atoi(myString)
		fmt.Printf("string is %s : number is %d\n", myString, num)
	*/

	var firstNum int
	var secondNum int
	var answer int

	firstNum = 1
	secondNum = 2
	answer = adder(firstNum, secondNum)

	fmt.Printf("%d plus %d equals %d\n", firstNum, secondNum, answer)
}
