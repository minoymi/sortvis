package main

func Quicksort() {
	quicksortNoWrapper(0, len(arr)-1)
}

func quicksortNoWrapper(low, high int) {
	if low > high {
		return
	}
	partitionIndex := quicksortPartition(low, high)

	quicksortNoWrapper(low, partitionIndex-1)
	quicksortNoWrapper(partitionIndex+1, high)

}

func quicksortPartition(low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			currentIndex <- i
		}
	}
	return i
}
