package main

import "strconv"

func Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	} else if len(numbers) == 1 {
		number, _ := strconv.Atoi(numbers)
		return number
	} else {
		firstNumber, _ := strconv.Atoi(numbers[:1])
		secondNumber, _ := strconv.Atoi(numbers[2:])
		return firstNumber + secondNumber
	}
}
