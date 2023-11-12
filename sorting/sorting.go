package sorting

func BubbleSort(values []int) {
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values)-1-i; j++ {
			currentValue := values[j]
			valueToCheck := values[j+1]
			if currentValue > valueToCheck {
				swap(values, j, j+1)
			}
		}
	}
}

func swap(values []int, index1, index2 int) {
	temp := values[index1]
	values[index1] = values[index2]
	values[index2] = temp
}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

func quickSort(values []int, low, high int) {
	if low >= high {
		return
	}

	pivotIndex := partition(values, low, high)
	quickSort(values, low, pivotIndex-1)
	quickSort(values, pivotIndex+1, high)
}

func partition(values []int, low, high int) int {
	pivot := values[high]
	index := low - 1

	for i := low; i < high; i++ {
		if values[i] <= pivot {
			index++
			swap(values, i, index)
		}
	}

	index++
	values[high] = values[index]
	values[index] = pivot

	return index
}
