module examples

go 1.22.0

toolchain go1.24.4

replace github.com/michaeldelgado1/raylib-go/raylib => ../raylib

replace github.com/michaeldelgado1/raylib-go/raygui => ../raygui

replace github.com/michaeldelgado1/raylib-go/easings => ../easings

replace github.com/michaeldelgado1/raylib-go/physics => ../physics

require (
	github.com/jakecoffman/cp v1.2.1
	github.com/michaeldelgado1/raylib-go/easings v0.0.0-00010101000000-000000000000
	github.com/michaeldelgado1/raylib-go/physics v0.0.0-00010101000000-000000000000
	github.com/michaeldelgado1/raylib-go/raygui v0.0.0-00010101000000-000000000000
	github.com/michaeldelgado1/raylib-go/raylib v0.0.0-20250801060855-4d6b70c638c9
	github.com/neguse/go-box2d-lite v0.0.0-20170921151050-5d8ed9b7272b
)

require (
	github.com/ebitengine/purego v0.8.1 // indirect
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f // indirect
	golang.org/x/sys v0.27.0 // indirect
)
