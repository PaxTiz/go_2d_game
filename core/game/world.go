package game

import (
	"vcernuta/raylib/core/loader"
	"vcernuta/raylib/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type World struct {
	Game  *Game
	Tiles []Tile
}

func InitLevelFromDirectory(game *Game, textures utils.Textures, path string) World {
	level := utils.LoadLevelAsArray(path)

	levelTiles := []Tile{}
	for _, line := range level {
		rawTile := loader.InitFromString(line)
		tile := InitTileFromRawTile(game, rawTile, textures)
		levelTiles = append(levelTiles, tile)
	}

	return World{Game: game, Tiles: levelTiles}
}

func (level World) FindSolidTilesMatchingDirection(direction rl.Vector2) []Tile {
	tiles := []Tile{}

	for _, tile := range level.Tiles {
		rect := rl.NewRectangle(direction.X, direction.Y, 32, 32)

		if tile.Solid && rl.CheckCollisionRecs(tile.AsRect(), rect) {
			tiles = append(tiles, tile)
		}
	}

	return tiles
}
