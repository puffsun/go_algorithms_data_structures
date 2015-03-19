package searching

func BinarySearch(slice []int, element int) int {
	if len(slice) == 0 {
		return -1
	}

	begin, end := 0, len(slice)
	mid := ((end - begin) >> 1) + begin
	if slice[mid] == element {
		return mid
	} else if slice[mid] < element {
		return BinarySearch(slice[mid+1:end], element)
	} else {
		return BinarySearch(slice[0:mid], element)
	}
}
