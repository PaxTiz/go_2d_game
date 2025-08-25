package levels

import (
	"vcernuta/raylib/core/entities"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Level struct {
	Layers [][]entities.Entity
}

func InitLevelFromDirectory(textures utils.Textures, path string) Level {
	level := utils.LoadLevelAsArray(path)
	spreadsheetPositions := utils.TexturesSpreadsheetPositions()

	layers := [][]entities.Entity{}
	for _, layer := range level {
		lines := []entities.Entity{}

		for x, line := range layer {
			for y, element := range line {
				if element == "__empty__" {
					continue
				}

				position := spreadsheetPositions[element]
				entity := entities.InitEntityFromTexture(
					textures.MapSpritesheet,
					y*(16*utils.TEXTURE_SCALING),
					x*(16*utils.TEXTURE_SCALING),
					int(position.X),
					int(position.Y),
					16,
					16,
				)

				lines = append(lines, entity)
			}
		}

		layers = append(layers, lines)
	}

	return Level{Layers: layers}
}

// Draw each layer of the level in order to permit texture overloading
func (level Level) Draw(debug bool) {
	for _, layer := range level.Layers {
		for _, element := range layer {
			source := rl.Rectangle{
				X:      element.SpritesheetPosition.X,
				Y:      element.SpritesheetPosition.Y,
				Width:  element.Size.X,
				Height: element.Size.Y,
			}
			destination := rl.Rectangle{
				X:      element.Position.X,
				Y:      element.Position.Y,
				Width:  element.Size.X * utils.TEXTURE_SCALING,
				Height: element.Size.Y * utils.TEXTURE_SCALING,
			}
			origin := rl.Vector2{X: 0, Y: 0}

			rl.DrawTexturePro(element.Texture, source, destination, origin, 0, rl.White)

			if debug {
				rl.DrawRectangleLines(int32(destination.X), int32(destination.Y), int32(destination.Width), int32(destination.Height), rl.Red)
			}
		}
	}
}
