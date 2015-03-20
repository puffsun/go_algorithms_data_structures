package sorting

func InsertionSort(slice []int) {
	if slice == nil || len(slice) == 0 {
		return
	}

	for i := 1; i < len(slice); i++ {
		val := slice[i]
		j := i - 1
		for j >= 0 && slice[j] > val {
			slice[j+1] = slice[j]
		}
		slice[j+1] = val
	}
}
