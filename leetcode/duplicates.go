package main

import (
	"fmt"
)

func findDuplicates(numbers []int) bool {
	hash := map[int]int{}

	for i := 0; i < len(numbers); i++ {
		if _, ok := hash[numbers[i]]; !ok {
			hash[numbers[i]] = hash[numbers[i]] + 1
		} else {
			return true
		}
	}

	return false
}

func main() {
	numbers := []int{1, 2, 3, 4, 7}
	fmt.Println(findDuplicates(numbers))
}
