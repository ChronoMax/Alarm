package main

import "fmt"

// Give input
func main() {
	var input int
	_input := GetUserInput()
	input = _input
	DoTheLoop(input)
}

func GetUserInput() int {
	var input int
	fmt.Println("How many times should the alarm go off?")
	fmt.Scan(&input)
	return input
}

func DoTheLoop(_input int) {
	for i := 1; i <= _input; i++ {
		fmt.Println("Alarm ", i, "!")
	}
}
