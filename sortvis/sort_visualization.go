package main

import (
	"image/color"
	rl "sortvis/raylib"
)

var arr []int
var currentIndex chan int

func main() {
	CreateArray(DEFAULT_ARRAY_SIZE)
	currentIndex = make(chan int)

	InitWindow()
	RenderLoop()
}

func InitWindow() {
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, WINDOW_NAME)
	rl.SetTargetFPS(DEFAULT_FPS)
}

func RenderLoop() {
	defer rl.CloseWindow()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(BG_COLOR)

		RenderArray()
		RenderCurrentItem()
		MakeGUI()

		rl.EndDrawing()
	}
}

func RenderArray() {
	for i := 0; i < len(arr); i++ {
		RenderOneItem(i, arr[i], ITEM_COLOR)
	}
}

func RenderCurrentItem() {
	select {
	case i := <-currentIndex:
		RenderOneItem(i, arr[i], CURRENT_ITEM_COLOR)
	default:
		RenderOneItem(0, arr[0], CURRENT_ITEM_COLOR)
	}
}

func RenderOneItem(index, value int, color color.RGBA) {
	WidthScale := (WINDOW_WIDTH / len(arr))
	HeightScale := (WINDOW_HEIGHT / len(arr))

	posX := index * WidthScale
	posY := WINDOW_HEIGHT - (value * HeightScale)
	itemWidth := WidthScale - ITEM_GAP
	itemHeight := value * HeightScale

	rl.DrawRectangle(posX, posY, itemWidth, itemHeight, color)
}
