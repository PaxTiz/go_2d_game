package camera

import (
	"vcernuta/raylib/core/entities"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera struct {
	Camera rl.Camera2D
}

func InitCamera(player entities.Player) Camera {
	return Camera{
		Camera: rl.Camera2D{
			// As the camera is placed at the same position as the player, we need to offset the camera for it to be positionned at the center of the window
			Offset: rl.Vector2{X: (utils.WINDOW_WIDTH / 2) - player.Entity.Size.X*2, Y: (utils.WINDOW_HEIGHT / 2) - player.Entity.Size.Y*2},
			// Initially, the camera is at the same position as the player
			Target:   player.Entity.Position,
			Rotation: 0,
			Zoom:     1,
		},
	}
}

func (camera *Camera) SyncPositionWithPlayer(player entities.Player) {
	camera.Camera.Target = player.Entity.Position
}

func (camera *Camera) Zoom(keyboardLayout utils.KeyboardLayout, debug bool) {
	if debug {
		if camera.Camera.Zoom < 5 && rl.IsKeyDown(keyboardLayout.Modifier) && rl.IsKeyDown(keyboardLayout.Plus) {
			camera.Camera.Zoom += 0.1
		} else if camera.Camera.Zoom > 0.5 && rl.IsKeyDown(keyboardLayout.Modifier) && rl.IsKeyDown(keyboardLayout.Minus) {
			camera.Camera.Zoom -= 0.1
		}
	}
}
