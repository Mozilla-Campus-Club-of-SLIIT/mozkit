package engine

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/scripts"
)

type Item struct {
	TitleStr string `toml:"title"`
	DescStr  string `toml:"description"`
	Target   string `toml:"target"`
}

func (i Item) Title() string       { return i.TitleStr }
func (i Item) Description() string { return i.DescStr }
func (i Item) FilterValue() string { return i.TitleStr }

type Script struct {
	Title       string `toml:"title"`
	Description string `toml:"description"`
	Actions     []struct {
		Description string   `toml:"description"`
		Tool        string   `toml:"tool"`
		Arguments   []string `toml:"arguments"`
	} `toml:"action"`
}

func validateItem(item Item, location string) {
	if item.TitleStr == "" || item.DescStr == "" || item.Target == "" {
		fmt.Fprintf(os.Stderr,
			"Oh no! Mozkit found a broken item in '%s'!\n"+
				"There is a big difference between rare and raw... Moxy, IT IS RAW!!\n"+
				"Missing title, description, or target.\nFound -> title: '%s', description: '%s', target: '%s'\n",
			location, item.TitleStr, item.DescStr, item.Target,
		)
		os.Exit(1)
	}
}

func readError(location string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Uh oh! Mozkit couldn't read '%s'!\nMoxy, did you forget to include it again?!\nError: %v\n", location, err)
		os.Exit(1)
	}
}

func tomlDecodeError(location string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Yikes! Mozkit crashed while reading '%s'.\nMoxy, it's 'LeviOsa', not 'LevioSA'! Check your TOML syntax again.\nError: %v\n", location, err)
		os.Exit(1)
	}
}

func ScriptList() []Item {
	var items []Item

	entries, err := scripts.Scripts.ReadDir(".")
	readError("scripts", err)

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".toml") {
			path := entry.Name()
			content, err := scripts.Scripts.ReadFile(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Yikes! Mozkit couldn't open '%s' even though it saw it sitting there!\nError: %v\n", path, err)
				os.Exit(1)
			}

			var item Item
			_, err = toml.Decode(string(content), &item)
			tomlDecodeError(path, err)

			item.Target = path

			validateItem(item, path)

			items = append(items, item)
		}
	}

	sort.Slice(items, func(i, j int) bool {
		iTitle := strings.ToLower(items[i].TitleStr)
		jTitle := strings.ToLower(items[j].TitleStr)
		if iTitle == jTitle {
			return items[i].TitleStr < items[j].TitleStr
		}
		return iTitle < jTitle
	})

	return items
}

func LoadScript(path string) Script {
	content, err := scripts.Scripts.ReadFile(path)
	readError(path, err)

	var script Script
	_, err = toml.Decode(string(content), &script)
	tomlDecodeError(path, err)

	return script
}
