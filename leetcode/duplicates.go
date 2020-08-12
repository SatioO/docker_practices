package leetcode

// HasDuplicates checks whether the slice has duplicates
func HasDuplicates(numbers []int) bool {
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
