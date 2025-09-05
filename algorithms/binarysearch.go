package algorithms

// Add any imports here
import (
	"math"
)

// Return the index of the item, or -1 if it's not in the list
func BinarySearch(haystack []int, needle int) int {
	var min int = 0
	var max int = len(haystack)
	for min < max {
		var mid int = min + int(math.Floor(float64(max-min)/2))
		if haystack[mid] == needle {
			return mid
		} else if needle < haystack[mid] {
			max = mid
		} else {
			min = mid + 1
		}
	}
	return -1
}
