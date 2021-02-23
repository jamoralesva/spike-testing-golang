package main

import (
	"strconv"
	"strings"
)

func Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	} else {
		splitNumbers := strings.Split(numbers, ",")
		result := 0
		for i := 0; i < len(splitNumbers); i++ {
			number, _ := strconv.Atoi(splitNumbers[i])
			result += number
		}
		return result
	}
}
