package sorting

func ShellSort(slice []int) {
	if slice == nil || len(slice) == 0 {
		return
	}

	increment := len(slice) / 2
	for increment > 0 {
		for i := increment; i < len(slice); i++ {
			j := i
			temp := slice[i]

			for j >= increment && slice[j-increment] > temp {
				slice[j] = slice[j-increment]
				j = j - increment
			}
			slice[j] = temp
		}

		if increment == 2 {
			increment = 1
		} else {
			increment = int(increment * 5 / 11)
		}
	}
}
