package sorting

func BubbleSort(slice []int) {
	if slice == nil || len(slice) == 0 {
		return
	}

	for count := len(slice) - 1; ; count-- {
		swap := false
		for i := 1; i <= count; i++ {
			if slice[i-1] > slice[i] {
				slice[i-1], slice[i] = slice[i], slice[i-1]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}
