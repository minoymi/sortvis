package main

import "slices"

var lastSwapIndex int

func Bubblesort() {
	lastSwapIndex = len(arr)
	for !slices.IsSorted(arr) {
		bubblesortOnePass(lastSwapIndex)
	}
}

func bubblesortOnePass(lastSwapInLastRun int) {
	for i := 1; i < lastSwapInLastRun; i++ {
		bubblesortOneSwap(i)
	}
}

func bubblesortOneSwap(i int) {
	a, b := arr[i-1], arr[i]
	if a > b {
		arr[i] = a
		arr[i-1] = b
		lastSwapIndex = i
	}
	currentIndex <- i
}
