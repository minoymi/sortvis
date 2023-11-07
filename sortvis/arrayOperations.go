package main

import "math/rand"

func CreateArray(size int) {
	arr = []int{}
	for i := 1; i <= size; i++ {
		arr = append(arr, i)
	}
}

func ShuffleArray() {
	swap := func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
		currentIndex <- int(i)
	}
	rand.Shuffle(len(arr), swap)
}
