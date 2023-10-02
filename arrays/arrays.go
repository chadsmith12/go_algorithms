package search

func LinearSearch(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}

func BinarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack)

	for low < high {
		midPoint := low + (high-low)/2
		value := haystack[midPoint]

		if value == needle {
			return true
		} else if value > needle {
			high = midPoint
		} else {
			low = midPoint + 1
		}
	}

	return false
}
