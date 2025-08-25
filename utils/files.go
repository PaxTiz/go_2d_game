package utils

import (
	"os"
	"path"
	"strings"
)

func LoadLevelAsArray(directory string) []string {
	relativePath := path.Join(directory, "level.txt")
	file, err := os.ReadFile(relativePath)
	if err != nil {
		panic(err)
	}

	content := strings.TrimSpace(string(file))
	lines := strings.SplitSeq(content, "\n")

	tiles := []string{}
	for line := range lines {
		// Skip empty lines
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		// Check is line is a comment in the file
		if strings.HasPrefix(line, "#") {
			continue
		}

		tiles = append(tiles, line)
	}

	return tiles
}
