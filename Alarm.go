package main

import "fmt"

//main function
func main() {
	var input int
	_input := GetUserInput()
	input = _input
	DoTheLoop(input)
}

//Gets the users input
func GetUserInput() int {
	var input int
	fmt.Println("How many times should the alarm go off?")
	fmt.Scan(&input)
	return input
}

//Uses the input from the user and displays the text
func DoTheLoop(_input int) {
	for i := 1; i <= _input; i++ {
		fmt.Println("Alarm ", i, "!")
	}
}
