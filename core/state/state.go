package state

import (
	"vcernuta/raylib/core/camera"
	"vcernuta/raylib/core/entities/player"
	"vcernuta/raylib/core/levels"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type State struct {
	Debug bool

	Textures *utils.Textures

	Player *player.Player

	Camera *camera.Camera

	Level *levels.Level

	KeyboardLayout *utils.KeyboardLayout
}

func InitState(keyboardLayout utils.KeyboardLayout, debug bool) State {
	textures := utils.InitTextures()
	player := player.InitPlayer(textures)
	camera := camera.InitCamera(player)

	level := levels.InitLevelFromDirectory(textures, "./resources/levels/home")

	return State{
		Debug:          debug,
		Textures:       &textures,
		Player:         &player,
		Level:          &level,
		KeyboardLayout: &keyboardLayout,
		Camera:         &camera,
	}
}

func (state *State) HandleKeyboardEvents() {
	delta := rl.GetFrameTime()

	state.Player.HandleKeyboardEvents(delta, *state.Level, *state.KeyboardLayout, state.Debug)
	state.Camera.SyncPositionWithPlayer(*state.Player)
	state.Camera.Zoom(*state.KeyboardLayout, state.Debug)

	state.handleReloadKeypress()
}

func (state State) Draw() {
	state.Level.Draw(state.Debug)
	state.Player.Draw(state.Debug)

	state.debugFPS()
}

func (state State) debugFPS() {
	if state.Debug {
		rl.ClearColor(rl.White.R, rl.White.G, rl.White.B, 1)
		rl.DrawFPS(4, 4)
	}
}

func (state *State) handleReloadKeypress() {
	if rl.IsKeyPressed(rl.KeyR) {
		textures := utils.InitTextures()
		state.Textures = &textures

		newPlayer := player.InitPlayer(textures)
		state.Player = &newPlayer

		newCamera := camera.InitCamera(newPlayer)
		state.Camera = &newCamera

		level := levels.InitLevelFromDirectory(textures, "./resources/levels/home")
		state.Level = &level
	}
}
