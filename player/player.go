package player

import (
	"fmt"
	"os/exec"
	"strings"
)

func Play(params string) {
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
		play(sounds[bestIndex])
	}
}

func Stop() {
	cmd := exec.Command("pkill", "player")
	cmd.Run()
}

func play(sound Sound) {
	Stop()
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
