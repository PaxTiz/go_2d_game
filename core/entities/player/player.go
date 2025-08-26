package player

import (
	"log"
	"vcernuta/raylib/core/entities"
	"vcernuta/raylib/core/levels"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Entity entities.Entity

	CollisionRect rl.Rectangle
}

func InitPlayer(textures utils.Textures) Player {
	return Player{
		Entity: entities.Entity{
			Texture:             textures.PlayerSpritesheet,
			Size:                rl.Vector2{X: 16, Y: 32},
			Position:            rl.Vector2{X: utils.WINDOW_WIDTH/2 - 16*2, Y: utils.WINDOW_HEIGHT/2 - 32*2},
			SpritesheetPosition: rl.Vector2{X: 0, Y: 0},
		},
		CollisionRect: rl.Rectangle{
			X:      utils.WINDOW_WIDTH/2 - 16*2,
			Y:      (utils.WINDOW_HEIGHT/2 - 16*2),
			Width:  32,
			Height: 32,
		},
	}
}

// TODO: Snap player position to nearest tile if position if not perfectly aligned
// TODO: Character animation
func (player *Player) HandleKeyboardEvents(delta float32, level levels.Level, keyboardLayout utils.KeyboardLayout, debug bool) {
	speed := float32(150)
	displacement := speed * delta

	movement := player.Entity.Position

	movement.X -= _handleMovement(keyboardLayout.PlayerLeft, displacement, debug)
	movement.X += _handleMovement(keyboardLayout.PlayerRight, displacement, debug)
	movement.Y -= _handleMovement(keyboardLayout.PlayerTop, displacement, debug)
	movement.Y += _handleMovement(keyboardLayout.PlayerBottom, displacement, debug)

	if player.CheckCollisions(level, movement) {
		player.Entity.Position = movement
		player.CollisionRect.X = movement.X
		player.CollisionRect.Y = movement.Y + player.Entity.Size.Y
	}
}

func (player Player) Draw(debug bool) {
	source := rl.Rectangle{
		X:      player.Entity.SpritesheetPosition.X,
		Y:      player.Entity.SpritesheetPosition.Y,
		Width:  player.Entity.Size.X,
		Height: player.Entity.Size.Y,
	}
	destination := rl.Rectangle{
		X:      player.Entity.Position.X,
		Y:      player.Entity.Position.Y,
		Width:  player.Entity.Size.X * utils.TEXTURE_SCALING,
		Height: player.Entity.Size.Y * utils.TEXTURE_SCALING,
	}
	origin := rl.Vector2{X: 0, Y: 0}

	rl.DrawTexturePro(player.Entity.Texture, source, destination, origin, 0, rl.White)

	if debug {
		rl.DrawRectangleLines(int32(destination.X), int32(destination.Y), int32(destination.Width), int32(destination.Height), rl.Blue)
		rl.DrawRectangleRec(player.CollisionRect, rl.ColorAlpha(rl.Blue, 0.4))
	}
}

func (player Player) CheckCollisions(level levels.Level, position rl.Vector2) bool {
	tiles := level.FindSolidTilesMatchingDirection(rl.NewVector2(position.X, position.Y+32))
	return len(tiles) == 0
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
