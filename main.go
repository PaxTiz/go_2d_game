package main

import (
	"os"
	"vcernuta/raylib/core/game"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	args := os.Args
	debug := false
	if len(args) == 2 && args[1] == "true" {
		debug = true
	}

	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.InitWindow(utils.WINDOW_WIDTH, utils.WINDOW_HEIGHT, "Welcome from Raylib !")
	defer rl.CloseWindow()

	keyboardLayout := utils.InitKeyboardLayoutAzerty()

	game := game.InitGame(keyboardLayout, debug)
	defer game.Textures.Unload()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(game.Camera.Camera)

		rl.ClearBackground(rl.Black)

		game.HandleKeyboardEvents()
		game.Draw()

		rl.EndMode2D()
		rl.EndDrawing()
	}
}
