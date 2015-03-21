package sorting

func MergeSort(slice []int) {
	aux := make([]int, len(slice)/2+1)
	if len(slice) < 2 {
		return
	}

	mid := len(slice) / 2
	MergeSort(slice[:mid])
	MergeSort(slice[mid:])
	if slice[mid-1] < slice[mid] {
		return // why?
	}

	copy(aux, slice[:mid])
	left, right := 0, mid
	for i := 0; ; i++ {
		if aux[left] <= slice[right] {
			slice[i] = aux[left]
			left += 1
			if left == mid {
				break
			}
		} else {
			slice[i] = slice[right]
			right += 1
			if right == len(slice) {
				copy(slice[i+1:], aux[left:mid])
				break
			}
		}
	}
}

func MergeSort2(slice []int) {
	aux := make([]int, len(slice))
	sort(slice, aux, 0, len(slice)-1)
}

func sort(slice, aux []int, low, high int) {
	if high <= low {
		return
	}
	mid := low + (high-low)/2
	sort(slice, aux, low, mid)
	sort(slice, aux, mid+1, high)
	merge(slice, aux, low, mid, high)
}

func merge(slice, aux []int, low, mid, high int) {
	i, j := low, mid+1
	copy(aux, slice)

	for k := low; k <= high; k += 1 {
		if i > mid {
			slice[k] = aux[j]
			j += 1
		} else if j > high {
			slice[k] = aux[i]
			i += 1
		} else if aux[j] < aux[i] {
			slice[k] = aux[j]
			j += 1
		} else {
			slice[k] = aux[i]
			i += 1
		}
	}
}
