package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	Texture             rl.Texture2D
	Size                rl.Vector2
	SpritesheetPosition rl.Vector2
	Position            rl.Vector2
}

func InitEntity(texturePath string, xPos int, yPos int, xSpreadPos int, ySpreadPos int, width int, height int) Entity {
	texture := rl.LoadTexture(texturePath)

	return Entity{
		Texture:             texture,
		Size:                rl.Vector2{X: float32(width), Y: float32(height)},
		Position:            rl.Vector2{X: float32(xPos), Y: float32(yPos)},
		SpritesheetPosition: rl.Vector2{X: float32(xSpreadPos), Y: float32(ySpreadPos)},
	}
}

func InitEntityFromTexture(texture rl.Texture2D, xPos int, yPos int, xSpreadPos int, ySpreadPos int, width int, height int) Entity {
	return Entity{
		Texture:             texture,
		Size:                rl.Vector2{X: float32(width), Y: float32(height)},
		Position:            rl.Vector2{X: float32(xPos), Y: float32(yPos)},
		SpritesheetPosition: rl.Vector2{X: float32(xSpreadPos), Y: float32(ySpreadPos)},
	}
}
