package sorting

// Usually bucket sort is a sorting algorithm for integer (or some other data types)
// that takes advantage of the fact that integers have a finite number of bits.
//
// Below implementation assume that 0..m (m included)is the range of integers in the array,
// we can also change it a little to pass the lower range into this method.
//
// Care should taken that you do not make a off by 1 error.
func BucketSort(slice []int, m int) {
	if slice == nil || len(slice) == 0 {
		return
	}
	buckets := make([]int, m+1)
	for i := 0; i < len(slice); i++ {
		buckets[slice[i]] += 1
	}

	index := 0
	for i := 0; i <= m; i++ {
		for j := 0; j < buckets[i]; j++ {
			slice[index] = i
			index += 1
		}
	}
}
