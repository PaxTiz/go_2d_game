package state

import (
	"vcernuta/raylib/core/camera"
	"vcernuta/raylib/core/entities"
	"vcernuta/raylib/core/levels"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type State struct {
	Debug bool

	Textures *utils.Textures

	Player *entities.Player

	Camera *camera.Camera

	Level *levels.Level

	KeyboardLayout *utils.KeyboardLayout
}

func InitState(keyboardLayout utils.KeyboardLayout, debug bool) State {
	textures := utils.InitTextures()
	player := entities.InitPlayer(textures)
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

func (state State) HandleKeyboardEvents() {
	delta := rl.GetFrameTime()

	state.Player.HandleKeyboardEvents(delta, *state.KeyboardLayout, state.Debug)
	state.Camera.SyncPositionWithPlayer(*state.Player)
}

func (state State) Draw() {
	state.Level.Draw(state.Debug)
	state.Player.Draw(state.Debug)
}
