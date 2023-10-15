package sorting

func QuickSort(values []int) {
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
