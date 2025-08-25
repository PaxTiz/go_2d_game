package tiles

import (
	"log"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type RawTile struct {
	X           int
	Y           int
	Layer       int
	Transparent int
	Texture     string
}

func InitFromString(line string) RawTile {
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

	transparent, err := strconv.Atoi(parts[3])
	if err != nil {
		log.Fatalf("Could not parse %s to integer", parts[3])
	}

	layer, err := strconv.Atoi(parts[4])
	if err != nil {
		log.Fatalf("Could not parse %s to integer", parts[4])
	}

	return RawTile{
		X:           x,
		Y:           y,
		Layer:       layer,
		Transparent: transparent,
		Texture:     parts[2],
	}
}
