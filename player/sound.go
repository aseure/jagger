package player

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

type Sound struct {
	Extension string
	Keywords  []string
	Path      string
}

func listSounds(directory string) []Sound {
	fileInfos, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("No sound files found")
		return nil
	}

	var sounds []Sound
	for _, f := range fileInfos {
		name := f.Name()
		sounds = append(sounds, Sound{
			Extension: strings.TrimSpace(filepath.Ext(name)),
			Keywords:  getKeywords(name),
			Path:      path.Join(directory, name),
		})
	}

	return sounds
}

func getKeywords(name string) []string {
	extension := filepath.Ext(name)
	name = strings.TrimSuffix(name, extension)
	return strings.Split(name, "-")
}
