package raylib_adapter

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type raylib_adapter struct{}

func InitWindow(width int, height int, title string) {
	rl.InitWindow(int32(width), int32(height), title)
}

func SetTargetFPS(fps int) {
	rl.SetTargetFPS(int32(fps))
}

func CloseWindow() {
	rl.CloseWindow()
}

func WindowShouldClose() bool {
	return rl.WindowShouldClose()
}

func BeginDrawing() {
	rl.BeginDrawing()
}

func EndDrawing() {
	rl.EndDrawing()
}

func ClearBackground(col color.RGBA) {
	rl.ClearBackground(col)
}

func DrawRectangle(posX int, posY int, width int, height int, col color.RGBA) {
	rl.DrawRectangle(int32(posX), int32(posY), int32(width), int32(height), col)
}
