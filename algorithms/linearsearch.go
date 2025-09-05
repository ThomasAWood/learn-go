package algorithms

// Return index of the needle, or -1 if it's not in the haystack.
func LinearSearch(haystack []int, needle int) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return i
		}
	}
	return -1
}
