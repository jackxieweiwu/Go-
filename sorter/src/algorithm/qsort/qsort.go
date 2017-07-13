package qsort

func qsort(values []int, left, right int) int {
	temp := values[left]

	i, j := left, right

	for i < j {
		for i < j && values[j] >= temp {
			j--
		}

		values[i] = values[j]

		for i < j && values[i] <= temp {
			i++
		}

		values[j] = values[i]
	}

	values[i] = temp

	return i
}

func quicksort(values []int, left, right int) {
	privot := qsort(values, left, right)

	if privot-1 > left {
		quicksort(values, left, privot-1)
	}
	if privot+1 < right {
		quicksort(values, privot+1, right)
	}
}

func QuickSort(values []int) {

	quicksort(values, 0, len(values)-1)

}
