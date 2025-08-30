package game

import (
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera struct {
	Game   *Game
	Camera rl.Camera2D
}

func InitCamera(game *Game, player Player) Camera {
	return Camera{
		Game: game,
		Camera: rl.Camera2D{
			// As the camera is placed at the same position as the player, we need to offset the camera for it to be positionned at the center of the window
			Offset: rl.Vector2{X: (utils.WINDOW_WIDTH / 2) - player.Size.X*2, Y: (utils.WINDOW_HEIGHT / 2) - player.Size.Y*2},
			// Initially, the camera is at the same position as the player
			Target:   player.Position,
			Rotation: 0,
			Zoom:     1,
		},
	}
}

func (camera *Camera) SyncPositionWithPlayer() {
	camera.Camera.Target = camera.Game.Player.Position
}

func (camera *Camera) Zoom() {
	if camera.Game.Debug {
		kl := camera.Game.KeyboardLayout

		if camera.Camera.Zoom < 5 && rl.IsKeyDown(kl.Modifier) && rl.IsKeyDown(kl.Plus) {
			camera.Camera.Zoom += 0.1
		} else if camera.Camera.Zoom > 0.5 && rl.IsKeyDown(kl.Modifier) && rl.IsKeyDown(kl.Minus) {
			camera.Camera.Zoom -= 0.1
		}
	}
}
