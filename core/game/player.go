package game

import (
	"log"
	"vcernuta/raylib/core/renderer"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Id                  int
	Game                *Game
	Texture             rl.Texture2D
	Size                rl.Vector2
	SpritesheetPosition rl.Vector2
	Position            rl.Vector2
	CollisionRect       rl.Rectangle
}

func InitPlayer(game *Game, textures utils.Textures) Player {
	return Player{
		Game:                game,
		Texture:             textures.PlayerSpritesheet,
		Size:                rl.Vector2{X: 16, Y: 32},
		Position:            rl.Vector2{X: utils.WINDOW_WIDTH/2 - 16, Y: utils.WINDOW_HEIGHT/4 - 32*2},
		SpritesheetPosition: rl.Vector2{X: 0, Y: 0},
		CollisionRect: rl.Rectangle{
			X:      utils.WINDOW_WIDTH/2 - 16,
			Y:      (utils.WINDOW_HEIGHT/2 - 16*2),
			Width:  32,
			Height: 32,
		},
	}
}

// TODO: Snap player position to nearest tile if position if not perfectly aligned
// TODO: Character animation
func (player *Player) HandleKeyboardEvents(delta float32) {
	kl := player.Game.KeyboardLayout
	debug := player.Game.Debug

	speed := float32(150)
	displacement := speed * delta

	movement := player.Position

	movement.X -= _handleMovement(kl.PlayerLeft, displacement, debug)
	movement.X += _handleMovement(kl.PlayerRight, displacement, debug)
	movement.Y -= _handleMovement(kl.PlayerTop, displacement, debug)
	movement.Y += _handleMovement(kl.PlayerBottom, displacement, debug)

	if player.CheckCollisionsHorizontally(movement) {
		player.Position.X = movement.X
		player.CollisionRect.X = movement.X
	}
	if player.CheckCollisionsVertically(movement) {
		player.Position.Y = movement.Y
		player.CollisionRect.Y = movement.Y + player.Size.Y
	}

	player.Game.Renderer.UpdateObject(player.Id, player)
}

func (player Player) Render() {
	source := rl.Rectangle{
		X:      player.SpritesheetPosition.X,
		Y:      player.SpritesheetPosition.Y,
		Width:  player.Size.X,
		Height: player.Size.Y,
	}
	destination := rl.Rectangle{
		X:      player.Position.X,
		Y:      player.Position.Y,
		Width:  player.Size.X * utils.TEXTURE_SCALING,
		Height: player.Size.Y * utils.TEXTURE_SCALING,
	}
	origin := rl.Vector2{X: 0, Y: 0}

	rl.DrawTexturePro(player.Texture, source, destination, origin, 0, rl.White)

	if player.Game.Debug {
		rl.DrawRectangleLines(int32(destination.X), int32(destination.Y), int32(destination.Width), int32(destination.Height), rl.Blue)
		rl.DrawRectangleRec(player.CollisionRect, rl.ColorAlpha(rl.Blue, 0.4))
	}
}

func (player Player) CheckCollisionsHorizontally(position rl.Vector2) bool {
	level := player.Game.Level
	tiles := level.FindSolidTilesMatchingDirection(rl.NewVector2(position.X, player.Position.Y+32))
	return len(tiles) == 0
}

func (player Player) CheckCollisionsVertically(position rl.Vector2) bool {
	level := player.Game.Level
	tiles := level.FindSolidTilesMatchingDirection(rl.NewVector2(player.Position.X, position.Y+32))
	return len(tiles) == 0
}

func (player Player) AsRect() rl.Rectangle {
	return rl.NewRectangle(
		player.Position.X,
		player.Position.Y,
		player.Size.X,
		player.Size.Y,
	)
}

func (player Player) GetEntityData() renderer.EntityData {
	return renderer.EntityData{
		Position: player.Position,
		Size:     player.Size,
		Layer:    10,
	}
}

func _handleMovement(key int32, displacement float32, debug bool) float32 {
	if rl.IsKeyDown(key) {
		if debug {
			log.Printf("Key pressed : %s (move to top of %v)", utils.UnicodePointToLetter(key), displacement)
		}

		return displacement
	}

	return 0
}
