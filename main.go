package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(63)
	screenHeight := int32(69)

	windowPosition := rl.Vector2{X: 500, Y: 200}

	rl.SetConfigFlags(rl.FlagWindowUndecorated | rl.FlagWindowTransparent | rl.FlagWindowAlwaysRun | rl.FlagWindowTopmost)
	rl.InitWindow(screenWidth, screenHeight, "shimeji")

	rl.SetWindowPosition(int(windowPosition.X), int(windowPosition.Y))
	rl.SetTargetFPS(60)

	mousePosition := rl.Vector2{}
	panOffset := mousePosition
	dragWindow := false
	exitWindow := false

	character := rl.LoadImage("defaultCharacter.png")
	texture := rl.LoadTextureFromImage(character)
	rl.UnloadImage(character)

	for !exitWindow && !rl.WindowShouldClose() {
		mousePosition = rl.GetMousePosition()

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) && !dragWindow {
			if rl.CheckCollisionPointRec(mousePosition, rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight))) {
				dragWindow = true
				panOffset = mousePosition
			}
		}

		if dragWindow {
			windowPosition.X += mousePosition.X - panOffset.X
			windowPosition.Y += mousePosition.Y - panOffset.Y
			rl.SetWindowPosition(int(windowPosition.X), int(windowPosition.Y))

			if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				dragWindow = false
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Blank)

		rl.DrawTextureEx(texture, rl.Vector2{X: 0, Y: 0}, 0, 3, rl.White)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
