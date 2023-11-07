package main

func Mergesort() {
	mergesortNoWrapper(arr, 0)
}

func mergesortNoWrapper(slc []int, start int) {
	if len(slc) < 2 {
		return
	}

	mid := len(slc) / 2

	left := slc[:mid]
	right := slc[mid:]

	mergesortNoWrapper(left, start)
	mergesortNoWrapper(right, start+mid)

	merge(left, right, start)

}

func merge(first, second []int, pos int) {
	sorted := make([]int, 0)

	for len(first) > 0 && len(second) > 0 {
		if first[0] < second[0] {
			sorted = append(sorted, first[0])
			first = first[1:]
		} else {
			sorted = append(sorted, second[0])
			second = second[1:]
		}
	}

	sorted = append(sorted, first...)
	sorted = append(sorted, second...)

	copyToArr(sorted, pos)
}

func copyToArr(sorted []int, pos int) {
	for i := 0; i < len(sorted); i++ {
		arr[pos+i] = sorted[i]
		currentIndex <- pos + i
	}
}
