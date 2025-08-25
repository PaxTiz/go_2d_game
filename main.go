package main

import (
	"os"
	"vcernuta/raylib/core/state"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	args := os.Args
	debug := false
	if len(args) == 2 && args[1] == "true" {
		debug = true
	}

	rl.SetTargetFPS(60)

	rl.InitWindow(utils.WINDOW_WIDTH, utils.WINDOW_HEIGHT, "Welcome from Raylib !")
	defer rl.CloseWindow()

	keyboardLayout := utils.InitKeyboardLayoutAzerty()

	state := state.InitState(keyboardLayout, debug)
	defer state.Textures.Unload()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(state.Camera.Camera)

		rl.ClearBackground(rl.Black)

		state.HandleKeyboardEvents()
		state.Draw()

		rl.EndMode2D()
		rl.EndDrawing()
	}
}
