package entities

import (
	"log"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Entity Entity
}

func InitPlayer(textures utils.Textures) Player {
	return Player{
		Entity: Entity{
			Texture:             textures.PlayerSpritesheet,
			Size:                rl.Vector2{X: 16, Y: 32},
			Position:            rl.Vector2{X: utils.WINDOW_WIDTH/2 - 16*2, Y: utils.WINDOW_HEIGHT/2 - 32*2},
			SpritesheetPosition: rl.Vector2{X: 0, Y: 0},
		},
	}
}

// TODO: Character animation
func (player *Player) HandleKeyboardEvents(delta float32, keyboardLayout utils.KeyboardLayout, debug bool) {
	speed := float32(150)

	movement := player.Entity.Position

	movement.X -= handleMovement(keyboardLayout.PlayerLeft, speed, delta, debug)
	movement.X += handleMovement(keyboardLayout.PlayerRight, speed, delta, debug)
	movement.Y -= handleMovement(keyboardLayout.PlayerTop, speed, delta, debug)
	movement.Y += handleMovement(keyboardLayout.PlayerBottom, speed, delta, debug)

	player.Entity.Position = movement
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
	}
}

func handleMovement(key int32, speed float32, delta float32, debug bool) float32 {
	if rl.IsKeyDown(key) {
		if debug {
			log.Printf("Key pressed : %s (move to top of %v)", utils.UnicodePointToLetter(key), speed*delta)
		}

		return speed * delta
	}

	return 0
}
