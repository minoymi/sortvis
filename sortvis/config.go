package main

import rl "github.com/gen2brain/raylib-go/raylib"

var WINDOW_WIDTH int = 1200
var WINDOW_HEIGHT int = 800
var WINDOW_NAME string = "Sorting Algorithm Visualization"

var DEFAULT_FPS int = 60
var MIN_FPS int = 15
var MAX_FPS int = 4000

var BG_COLOR rl.Color = rl.Black
var ITEM_COLOR rl.Color = rl.White
var CURRENT_ITEM_COLOR rl.Color = rl.Orange
var ITEM_GAP int = 1

var DEFAULT_ARRAY_SIZE int = 400
var MIN_ARRAY_SIZE int = 5
var MAX_ARRAY_SIZE int = 800

var ALGOCHOICE_BOX_BOUNDS = rl.Rectangle{X: 0, Y: 0, Width: 100, Height: 30}
var RUN_BUTTON_BOUNDS = rl.Rectangle{X: 100, Y: 0, Width: 30, Height: 30}
var FPS_VALUEBOX_BOUNDS = rl.Rectangle{X: 160, Y: 0, Width: 30, Height: 30}
var ARRAYSIZE_VALUEBOX_BOUNDS = rl.Rectangle{X: 250, Y: 0, Width: 30, Height: 30}
