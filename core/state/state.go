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
}

func (state State) Draw() {
	state.Level.Draw(state.Debug)
	state.Player.Draw(state.Debug)

	state.DebugFPS()
}

func (state State) DebugFPS() {
	if state.Debug {
		rl.ClearColor(rl.White.R, rl.White.G, rl.White.B, 1)
		rl.DrawFPS(4, 4)
	}
}
