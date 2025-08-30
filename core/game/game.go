package game

import (
	"vcernuta/raylib/core/renderer"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Debug bool

	Textures *utils.Textures

	Player *Player

	Camera *Camera

	Level *World

	KeyboardLayout *utils.KeyboardLayout

	Renderer *renderer.Renderer
}

func InitGame(keyboardLayout utils.KeyboardLayout, debug bool) Game {
	game := Game{}

	textures := utils.InitTextures()
	player := InitPlayer(&game, textures)
	camera := InitCamera(&game, player)
	renderer := renderer.InitRenderer()

	player.Id = renderer.AddObject(player)

	level := InitWorldFromDirectory(&game, textures, "./resources/levels/home")
	for _, tile := range level.Tiles {
		tile.Id = renderer.AddObject(tile)
	}

	game.Debug = debug
	game.Textures = &textures
	game.Player = &player
	game.Level = &level
	game.KeyboardLayout = &keyboardLayout
	game.Camera = &camera
	game.Renderer = &renderer

	return game
}

func (game *Game) HandleKeyboardEvents() {
	delta := rl.GetFrameTime()

	game.Player.HandleKeyboardEvents(delta)
	game.Camera.SyncPositionWithPlayer()
	game.Camera.Zoom()

	game.handleReloadKeypress()
}

func (game Game) Draw() {
	game.Renderer.Render()

	game.debugFPS()
}

func (game Game) debugFPS() {
	if game.Debug {
		rl.ClearColor(rl.White.R, rl.White.G, rl.White.B, 1)
		rl.DrawFPS(4, 4)
	}
}

func (game *Game) handleReloadKeypress() {
	if rl.IsKeyPressed(rl.KeyR) {
		textures := utils.InitTextures()
		game.Textures = &textures

		newPlayer := InitPlayer(game, textures)
		game.Player = &newPlayer

		newCamera := InitCamera(game, newPlayer)
		game.Camera = &newCamera

		level := InitWorldFromDirectory(game, textures, "./resources/levels/home")
		game.Level = &level
	}
}
