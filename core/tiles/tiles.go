package tiles

import (
	"vcernuta/raylib/core/entities"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	Entity      entities.Entity
	Layer       int
	Transparent bool
}

func InitFromRawTile(tile RawTile, textures utils.Textures) Tile {
	spreadsheetPositions := utils.TexturesSpreadsheetPositions()

	position := spreadsheetPositions[tile.Texture]
	entity := entities.InitEntityFromTexture(
		textures.MapSpritesheet,
		tile.X*(16*utils.TEXTURE_SCALING),
		tile.Y*(16*utils.TEXTURE_SCALING),
		int(position.X),
		int(position.Y),
		16,
		16,
	)

	return Tile{
		Entity:      entity,
		Layer:       tile.Layer,
		Transparent: tile.Transparent == 0,
	}
}

func (tile Tile) Draw(debug bool) {
	source := rl.Rectangle{
		X:      tile.Entity.SpritesheetPosition.X,
		Y:      tile.Entity.SpritesheetPosition.Y,
		Width:  tile.Entity.Size.X,
		Height: tile.Entity.Size.Y,
	}
	destination := rl.Rectangle{
		X:      tile.Entity.Position.X,
		Y:      tile.Entity.Position.Y,
		Width:  tile.Entity.Size.X * utils.TEXTURE_SCALING,
		Height: tile.Entity.Size.Y * utils.TEXTURE_SCALING,
	}
	origin := rl.Vector2{X: 0, Y: 0}

	rl.DrawTexturePro(tile.Entity.Texture, source, destination, origin, 0, rl.White)

	if debug {
		rl.DrawRectangleLines(int32(destination.X), int32(destination.Y), int32(destination.Width), int32(destination.Height), rl.Red)
	}
}
