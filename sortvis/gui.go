package main

import (
	rl "sortvis/raylib"

	gui "github.com/gen2brain/raylib-go/raygui"
)

func MakeGUI() {
	MakeAlgoChoiceDropdownBox()
	MakeRunButton()
	MakeFpsValuebox()
	MakeArraySizeValuebox()
}

var algoChoiceBoxEditmode bool = false
var currentAlgoChoice int32 = 0
var algoChoicesText string = "shuffle;quicksort;mergesort;bubblesort"
var algos []func() = []func(){ShuffleArray, Quicksort, Mergesort, Bubblesort}

func MakeAlgoChoiceDropdownBox() {
	if gui.DropdownBox(
		ALGOCHOICE_BOX_BOUNDS,
		algoChoicesText,
		&currentAlgoChoice,
		algoChoiceBoxEditmode,
	) {
		algoChoiceBoxEditmode = !algoChoiceBoxEditmode
	}
}

func MakeRunButton() {
	if gui.Button(RUN_BUTTON_BOUNDS, "Run") {
		go algos[currentAlgoChoice]()
	}

}

var fpsValueboxEditmode bool = false
var fpsValueboxValue int32 = int32(DEFAULT_FPS)

func MakeFpsValuebox() {
	if gui.ValueBox(
		FPS_VALUEBOX_BOUNDS,
		"FPS",
		&fpsValueboxValue,
		MIN_FPS,
		MAX_FPS,
		fpsValueboxEditmode,
	) {
		fpsValueboxEditmode = !fpsValueboxEditmode
		rl.SetTargetFPS(int(fpsValueboxValue))
	}
}

var arraySizeValueboxEditmode bool = false
var arraySizeValueboxValue int32 = int32(DEFAULT_ARRAY_SIZE)

func MakeArraySizeValuebox() {
	if gui.ValueBox(
		ARRAYSIZE_VALUEBOX_BOUNDS,
		"Array Size",
		&arraySizeValueboxValue,
		MIN_ARRAY_SIZE,
		MAX_ARRAY_SIZE,
		arraySizeValueboxEditmode,
	) {
		arraySizeValueboxEditmode = !arraySizeValueboxEditmode
		CreateArray(int(arraySizeValueboxValue))
	}
}
