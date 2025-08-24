package utils

import (
	"os"
	"path"
	"regexp"
	"slices"
	"strings"

	"github.com/samber/lo"
)

func LoadLevelAsArray(directory string) [][][]string {
	items, err := os.ReadDir(directory)
	if err != nil {
		panic(err)
	}

	// Map over directory file to find only files that starts with `layer_`
	files := lo.Map(lo.Filter(items, func(item os.DirEntry, index int) bool {
		matches, err := regexp.MatchString(LEVEL_FILENAME_REGEX, item.Name())
		if err != nil {
			panic(err)
		}

		return !item.IsDir() && matches
	}), func(item os.DirEntry, index int) string {
		return item.Name()
	})

	slices.Sort(files)

	levels := [][][]string{}

	for _, file := range files {
		relativePath := path.Join(directory, file)
		file, err := os.ReadFile(relativePath)
		if err != nil {
			panic(err)
		}

		content := strings.TrimSpace(string(file))
		lines := strings.SplitSeq(content, "\n")

		level := [][]string{}
		for line := range lines {
			names := []string{}

			rawNames := strings.SplitSeq(line, " ")
			for name := range rawNames {
				names = append(names, strings.TrimSpace(name))
			}

			level = append(level, names)
		}

		levels = append(levels, level)
	}

	return levels
}
