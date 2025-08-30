package game

import (
	"vcernuta/raylib/core/loader"
	"vcernuta/raylib/core/renderer"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	Id                  int
	Game                *Game
	Texture             rl.Texture2D
	Size                rl.Vector2
	SpritesheetPosition rl.Vector2
	Position            rl.Vector2
	Layer               int
	Solid               bool
}

func InitTileFromRawTile(game *Game, tile loader.RawTile, textures utils.Textures) Tile {
	spreadsheetPositions := utils.TexturesSpreadsheetPositions()

	position := spreadsheetPositions[tile.Texture]
	return Tile{
		Game:                game,
		Texture:             textures.MapSpritesheet,
		Position:            rl.NewVector2(float32(tile.X*(16*utils.TEXTURE_SCALING)), float32(tile.Y*(16*utils.TEXTURE_SCALING))),
		SpritesheetPosition: rl.NewVector2(position.X, position.Y),
		Size:                rl.NewVector2(16, 16),
		Layer:               tile.Layer,
		Solid:               tile.Solid == 1,
	}
}

func (tile Tile) Render() {
	source := rl.Rectangle{
		X:      tile.SpritesheetPosition.X,
		Y:      tile.SpritesheetPosition.Y,
		Width:  tile.Size.X,
		Height: tile.Size.Y,
	}
	destination := rl.Rectangle{
		X:      tile.Position.X,
		Y:      tile.Position.Y,
		Width:  tile.Size.X * utils.TEXTURE_SCALING,
		Height: tile.Size.Y * utils.TEXTURE_SCALING,
	}
	origin := rl.Vector2{X: 0, Y: 0}

	rl.DrawTexturePro(tile.Texture, source, destination, origin, 0, rl.White)

	if tile.Game.Debug {
		rl.DrawRectangleLines(int32(destination.X), int32(destination.Y), int32(destination.Width), int32(destination.Height), rl.Red)

		if tile.Solid {
			rl.DrawRectangleRec(destination, rl.ColorAlpha(rl.Red, 0.4))
		}
	}
}

func (tile Tile) AsRect() rl.Rectangle {
	return rl.NewRectangle(
		tile.Position.X,
		tile.Position.Y,
		tile.Size.X*utils.TEXTURE_SCALING,
		tile.Size.Y*utils.TEXTURE_SCALING,
	)
}

func (tile Tile) GetEntityData() renderer.EntityData {
	return renderer.EntityData{
		Position: tile.Position,
		Size:     tile.Size,
		Layer:    tile.Layer,
	}
}
