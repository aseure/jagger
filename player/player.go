package player

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func Play(params string, response chan<- string) {
	sounds := listSounds("sounds")
	if len(sounds) == 0 {
		return
	}

	var bestScore float64 = 0
	var bestIndex int = 0
	for i, sound := range sounds {
		score := computeScore(sound.Keywords, strings.Split(params, " "))
		if score > bestScore {
			bestScore = score
			bestIndex = i
		}
	}

	if bestScore > 0 {
		bestSound := sounds[bestIndex]
		response <- bestSound.Name
		play(bestSound)
	}
}

func Stop(response chan<- string) {
	if response != nil {
		response <- ""
	}
	cmd := exec.Command("pkill", "player")
	cmd.Run()
}

func List(response chan<- string) {
	var list []string
	sounds := listSounds("sounds")

	for _, s := range sounds {
		list = append(list, s.Name)
	}

	b, err := json.Marshal(list)
	if err != nil {
		response <- ""
	} else {
		response <- string(b)
	}
}

func play(sound Sound) {
	Stop(nil)
	var program string

	switch sound.Extension {
	case ".mp3", ".wav", ".mov":
		program = "mplayer"
	default:
		fmt.Println("Extension unsupported: ", sound.Extension)
		return
	}

	fmt.Println("[PLAY]", sound.Path)
	cmd := exec.Command(program, sound.Path)
	cmd.Run()
}
