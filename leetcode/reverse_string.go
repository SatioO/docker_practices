package leetcode

import (
	"strings"
)

// ReverseString defines a logic to reverse a given string
func ReverseString(str string) string {
	arr := strings.Split(str, "")
	i := 0
	j := len(arr) - 1

	for i < j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp

		i++
		j--
	}

	return strings.Join(arr, "")
}
