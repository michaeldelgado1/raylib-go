package main

import (
	"fmt"

	gui "github.com/michaeldelgado1/raylib-go/raygui"
	rl "github.com/michaeldelgado1/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raygui - button")

	rl.SetTargetFPS(60)

	var button bool

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		button = gui.Button(rl.NewRectangle(50, 150, 100, 40), "Click")
		if button {
			fmt.Println("Clicked on button")
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
