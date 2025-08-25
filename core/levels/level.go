package levels

import (
	"sort"
	"vcernuta/raylib/core/tiles"
	"vcernuta/raylib/utils"

	"github.com/samber/lo"
)

type Level struct {
	Tiles []tiles.Tile
}

func InitLevelFromDirectory(textures utils.Textures, path string) Level {
	level := utils.LoadLevelAsArray(path)

	levelTiles := []tiles.Tile{}
	for _, line := range level {
		rawTile := tiles.InitFromString(line)
		tile := tiles.InitFromRawTile(rawTile, textures)
		levelTiles = append(levelTiles, tile)
	}

	return Level{Tiles: levelTiles}
}

// Draw each layer of the level in order to permit texture overloading
func (level Level) Draw(debug bool) {
	layers := lo.GroupBy(level.Tiles, func(item tiles.Tile) int {
		return item.Layer
	})

	keys := lo.Keys(layers)
	sort.Ints(keys)

	for _, key := range keys {
		for _, element := range layers[key] {
			element.Draw(debug)
		}
	}
}
