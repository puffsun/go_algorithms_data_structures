package sorting

func QuickSortWithAnonymousFunc(slice []int) {
	if slice == nil || len(slice) == 0 {
		return
	}

	var sortRecurse func(left, right int)
	var partition func(left, right, pivot int) int

	partition = func(left, right, pivot int) int {
		v := slice[pivot]
		right -= 1
		slice[pivot], slice[right] = slice[right], slice[pivot]

		for i := left; i < right; i += 1 {
			if slice[i] <= v {
				slice[i], slice[left] = slice[left], slice[i]
				left += 1
			}
		}
		slice[left], slice[right] = slice[right], slice[left]
		return left
	}

	sortRecurse = func(left, right int) {
		if left < right {
			// pick the first element as pivot
			pivot := left
			pivot = partition(left, right, pivot)
			sortRecurse(left, pivot)
			sortRecurse(pivot+1, right)
		}
	}
	sortRecurse(0, len(slice))
}

func QuickSort(slice []int) {
	if slice == nil || len(slice) == 0 {
		return
	}

	quickSort(slice, 0, len(slice))
}

func partition(slice []int, left, right, pivot int) int {
	v := slice[pivot]
	right -= 1
	slice[pivot], slice[right] = slice[right], slice[pivot]

	for i := left; i < right; i += 1 {
		if slice[i] <= v {
			slice[i], slice[left] = slice[left], slice[i]
			left += 1
		}
	}
	slice[left], slice[right] = slice[right], slice[left]
	return left
}

func quickSort(slice []int, left, right int) {
	if left < right {
		// pick the first element as pivot
		pivot := left
		pivot = partition(slice, left, right, pivot)
		quickSort(slice, left, pivot)
		quickSort(slice, pivot+1, right)
	}
}
