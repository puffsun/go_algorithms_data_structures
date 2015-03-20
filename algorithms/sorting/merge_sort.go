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
