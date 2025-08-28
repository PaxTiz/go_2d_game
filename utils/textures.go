package utils

import rl "github.com/gen2brain/raylib-go/raylib"

type Textures struct {
	MapSpritesheet rl.Texture2D

	PlayerSpritesheet rl.Texture2D
}

func InitTextures() Textures {
	return Textures{
		MapSpritesheet: rl.LoadTexture("./resources/textures/tiles/spritesheets/transparent.png"),

		PlayerSpritesheet: rl.LoadTexture("./resources/textures/characters/spr_ardley.png"),
	}
}

func (textures Textures) Unload() {
	rl.UnloadTexture(textures.MapSpritesheet)
	rl.UnloadTexture(textures.PlayerSpritesheet)
}

func TexturesSpreadsheetPositions() map[string]rl.Vector2 {
	return map[string]rl.Vector2{
		"__empty__": {X: -1, Y: -1},

		// Pool
		"pool_top_left_corner":  {X: float32(SpreadsheetTexturePosition(3)), Y: 0},
		"pool_top_center":       {X: float32(SpreadsheetTexturePosition(4)), Y: 0},
		"pool_top_right_corner": {X: float32(SpreadsheetTexturePosition(5)), Y: 0},

		"pool_side_left":  {X: float32(SpreadsheetTexturePosition(3)), Y: float32(SpreadsheetTexturePosition(2))},
		"pool_full_water": {X: float32(SpreadsheetTexturePosition(4)), Y: float32(SpreadsheetTexturePosition(2))},
		"pool_side_right": {X: float32(SpreadsheetTexturePosition(5)), Y: float32(SpreadsheetTexturePosition(2))},

		"pool_bottom_left_corner":  {X: float32(SpreadsheetTexturePosition(3)), Y: float32(SpreadsheetTexturePosition(3))},
		"pool_bottom_center":       {X: float32(SpreadsheetTexturePosition(4)), Y: float32(SpreadsheetTexturePosition(3))},
		"pool_bottom_right_corner": {X: float32(SpreadsheetTexturePosition(5)), Y: float32(SpreadsheetTexturePosition(3))},

		// Grass
		"grass_1": {X: float32(SpreadsheetTexturePosition(6)), Y: 0},
		"grass_2": {X: float32(SpreadsheetTexturePosition(6)), Y: float32(SpreadsheetTexturePosition(1))},

		// Dirt
		"dirt_1": {X: float32(SpreadsheetTexturePosition(7)), Y: 0},
		"dirt_2": {X: float32(SpreadsheetTexturePosition(7)), Y: float32(SpreadsheetTexturePosition(1))},

		// Green tree
		"tree_green_pic_top":    {X: float32(SpreadsheetTexturePosition(17)), Y: float32(SpreadsheetTexturePosition(11))},
		"tree_green_pic_bottom": {X: float32(SpreadsheetTexturePosition(17)), Y: float32(SpreadsheetTexturePosition(12))},

		// Home walls
		"home_wall_white_bottom_left_corner":  {X: float32(SpreadsheetTexturePosition(15)), Y: float32(SpreadsheetTexturePosition(16))},
		"home_wall_white_bottom_right_corner": {X: float32(SpreadsheetTexturePosition(17)), Y: float32(SpreadsheetTexturePosition(16))},
		"home_wall_white_full":                {X: float32(SpreadsheetTexturePosition(19)), Y: float32(SpreadsheetTexturePosition(16))},

		// Home roofs
		"home_roof_wood_top_left_corner":     {X: float32(SpreadsheetTexturePosition(21)), Y: float32(SpreadsheetTexturePosition(22))},
		"home_roof_wood_top_right_corner":    {X: float32(SpreadsheetTexturePosition(22)), Y: float32(SpreadsheetTexturePosition(22))},
		"home_roof_wood_left":                {X: float32(SpreadsheetTexturePosition(21)), Y: float32(SpreadsheetTexturePosition(23))},
		"home_roof_wood_right":               {X: float32(SpreadsheetTexturePosition(22)), Y: float32(SpreadsheetTexturePosition(23))},
		"home_roof_wood_bottom_left_corner":  {X: float32(SpreadsheetTexturePosition(21)), Y: float32(SpreadsheetTexturePosition(24))},
		"home_roof_wood_bottom_right_corner": {X: float32(SpreadsheetTexturePosition(22)), Y: float32(SpreadsheetTexturePosition(24))},

		// Doors
		"door_rounded_black_background_left":  {X: float32(SpreadsheetTexturePosition(29)), Y: float32(SpreadsheetTexturePosition(9))},
		"door_rounded_black_background_right": {X: float32(SpreadsheetTexturePosition(30)), Y: float32(SpreadsheetTexturePosition(9))},
	}
}

// Convert a human index in the speadsheet to an image position
func SpreadsheetTexturePosition(humanIndex int) int {
	return (humanIndex-1)*16 + (humanIndex - 1)
}
