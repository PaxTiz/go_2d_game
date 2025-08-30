package loader

import (
	"fmt"
	"iter"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type RawTile struct {
	X       int
	Y       int
	Layer   int
	Solid   int
	Texture string
}

func InitRawTileFromString(line string, baseTile *RawTile) RawTile {
	baseX := 0
	baseY := 0
	baseLayer := 0
	if baseTile != nil {
		baseX = baseTile.X
		baseY = baseTile.Y
		baseLayer = baseTile.Layer
	}

	parts := lo.Map(strings.Split(line, " "), func(item string, _ int) string {
		return strings.TrimSpace(item)
	})

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatalf("Could not parse %s to integer", parts[0])
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("Could not parse %s to integer", parts[1])
	}

	solid, err := strconv.Atoi(parts[3])
	if err != nil {
		log.Fatalf("Could not parse %s to integer", parts[3])
	}

	layer, err := strconv.Atoi(parts[4])
	if err != nil {
		log.Fatalf("Could not parse %s to integer", parts[4])
	}

	return RawTile{
		X:       baseX + x,
		Y:       baseY + y,
		Layer:   baseLayer + layer,
		Solid:   solid,
		Texture: parts[2],
	}
}

func LoadLevelAsArray(directory string) []RawTile {
	return loadFromPath(nil, directory, "level.txt")
}

func loadFromPath(baseTile *RawTile, directory string, subpath ...string) []RawTile {
	fullPath := []string{directory}
	for _, s := range subpath {
		fullPath = append(fullPath, s)
	}

	relativePath := path.Join(fullPath...)
	file, err := os.ReadFile(relativePath)
	if err != nil {
		panic(err)
	}

	content := strings.TrimSpace(string(file))
	lines := strings.SplitSeq(content, "\n")

	return parseLines(baseTile, directory, lines)
}

func parseLines(baseTile *RawTile, directory string, lines iter.Seq[string]) []RawTile {
	tiles := []RawTile{}
	for line := range lines {
		// Skip empty lines
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		// Check is line is a comment in the file
		if strings.HasPrefix(line, "#") {
			continue
		}

		tile := InitRawTileFromString(line, baseTile)
		if model, ok := strings.CutPrefix(tile.Texture, "model#"); ok {
			children := loadFromPath(&tile, directory, "models", fmt.Sprintf("%s.txt", model))
			for _, child := range children {
				tiles = append(tiles, child)
			}
		} else {
			tiles = append(tiles, tile)
		}

	}

	return tiles
}
