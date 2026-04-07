package engine

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/components"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/scripts"
)

type Collection struct {
	Collection []components.Item `toml:"collection"`
}

func validateItem(item components.Item, location string) {
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

func undecodedWarning(location string, undecoded []toml.Key) {
	if len(undecoded) > 0 {
		log.Printf("Hmm... Mozkit noticed some strange extra items in '%s'. Moxy, what are these doing here?: %v\n", location, undecoded)
	}
}

func CollectionList() []components.Item {
	const fileName = "collection.toml"
	content, err := scripts.Scripts.ReadFile(fileName)
	readError(fileName, err)

	var collection Collection
	metadata, err := toml.Decode(string(content), &collection)
	tomlDecodeError(fileName, err)
	undecodedWarning(fileName, metadata.Undecoded())

	for i, col := range collection.Collection {
		validateItem(col, fmt.Sprintf("Collection #%d", i+1))
	}

	return collection.Collection
}

func ScriptList(directory string) []components.Item {
	var items []components.Item

	entries, err := scripts.Scripts.ReadDir(directory)
	readError(directory, err)

	for _, entry := range entries {
		if !entry.IsDir() && len(entry.Name()) > 5 && entry.Name()[len(entry.Name())-5:] == ".toml" {

			path := directory + "/" + entry.Name()
			content, err := scripts.Scripts.ReadFile(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Yikes! Mozkit couldn't open '%s' even though it saw it sitting there!\nError: %v\n", path, err)
				os.Exit(1)
			}

			var item components.Item
			_, err = toml.Decode(string(content), &item)
			tomlDecodeError(path, err)

			item.Target = path

			validateItem(item, path)

			items = append(items, item)
		}
	}

	return items
}
