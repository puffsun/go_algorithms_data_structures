package sorting

func SelectionSort(slice []int) {
	if slice == nil || len(slice) == 0 {
		return
	}

	for i := 0; i < len(slice); i++ {
		min := i
		for j := i; j < len(slice); j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}

		if min != i {
			slice[i], slice[min] = slice[min], slice[i]
		}
	}
}
