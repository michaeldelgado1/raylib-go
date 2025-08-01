package main

import (
	"github.com/michaeldelgado1/raylib-go/raygui"
	rl "github.com/michaeldelgado1/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 600, "raylib-go - icons example")
	defer rl.CloseWindow()

	raygui.LoadIcons("default_icons_with_255.rgi", false)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		raygui.DrawIcon(raygui.ICON_255, 100, 100, 8, rl.Gray)
		rl.EndDrawing()
	}
}
