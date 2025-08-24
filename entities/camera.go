package entities

import (
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera struct {
	Camera rl.Camera2D
}

func InitCamera(player Player) Camera {
	return Camera{
		Camera: rl.Camera2D{
			Offset:   rl.Vector2{X: (utils.WINDOW_WIDTH / 2) - player.Entity.Size.X*2, Y: (utils.WINDOW_HEIGHT / 2) - player.Entity.Size.Y*2},
			Target:   player.Entity.Position,
			Rotation: 0,
			Zoom:     1,
		},
	}
}

func (camera *Camera) SyncPositionWithPlayer(player Player) {
	camera.Camera.Target = player.Entity.Position
}
