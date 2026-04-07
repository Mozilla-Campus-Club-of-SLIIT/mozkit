package engine

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/internal/components"
	"github.com/Mozilla-Campus-Club-of-SLIIT/mozkit/scripts"
)

// func loadScripts() {
// 	//? Load scripts from the embedded filesystem: scripts/scripts.go
// 	fs.WalkDir(scripts.Scripts, ".", func(path string, d fs.DirEntry, err error) error {
// 		if !d.IsDir() && strings.HasSuffix(d.Name(), ".toml") {
// 			fmt.Println("Found:", path)
// 		}
// 		return nil
// 	})
// }

type Describe struct {
	Collection []struct {
		Name        string `toml:"name"`
		Description string `toml:"description"`
		Directory   string `toml:"directory"`
	} `toml:"collection"`
}

func ScriptList() []components.Item {
	content, err := scripts.Scripts.ReadFile("describe.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Uh oh! Mozkit couldn't find the 'describe.toml' blueprint...\nMoxy, did you forget to include it again?!\nError: %v\n", err)
		os.Exit(1)
	}

	var describe Describe
	metadata, err := toml.Decode(string(content), &describe)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Yikes! Mozkit crashed while reading 'describe.toml'.\nMoxy, it's 'LeviOsa', not 'LevioSA'! Check your TOML syntax again.\nError: %v\n", err)
		os.Exit(1)
	}

	if undecoded := metadata.Undecoded(); len(undecoded) > 0 {
		log.Printf("Hmm... Mozkit noticed some strange extra items in 'describe.toml'. Moxy, what are these doing here?: %v\n", undecoded)
	}

	var items []components.Item

	for i, col := range describe.Collection {
		if col.Name == "" || col.Description == "" || col.Directory == "" {
			fmt.Fprintf(os.Stderr,
				"Oh no! Mozkit found a broken collection in 'describe.toml'!\n"+
					"There is a big difference between rare and raw... Moxy, IT IS RAW!!\n"+
					"Check Collection #%d. Here is all I see:\n"+
					"-> name: '%s', description: '%s', directory: '%s'\n",
				i+1, col.Name, col.Description, col.Directory,
			)
			os.Exit(1)
		}

		items = append(items, components.Item{
			TitleStr: col.Name,
			DescStr:  col.Description,
		})
	}

	return items
}
