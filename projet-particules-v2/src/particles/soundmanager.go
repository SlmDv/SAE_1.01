package particles

import (
	"os"
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

// SoundManager{} est une structure qui possède trois éléments fournis par ebiten/audio.
// Elle contient le "contexte audio" global ainsi que les lecteurs associés aux sons
// configurés pour les clics gauche et droit. Les champs LeftSound et RightSound
// peuvent être nil si les options dans config.json sont désactivées ou si le chargement des fichiers audio a échoué.

type SoundManager struct {
	Context    *audio.Context
	LeftSound  *audio.Player
	RightSound *audio.Player
}

// NewSoundManager() est une fonction qui initialise le gestionnaire audio d'ebiten en fonction des paramètres dans config.json.
// Elle prend en argument un context, et retourne un élément SoundManager contenant le contexte passé en argument, y compris si l'argument passé est nil.
func NewSoundManager(context *audio.Context) *SoundManager {
	sound := &SoundManager{
		Context: context,
	}

	if config.General.EnableLeftClickSound {
		file, _ := os.Open("assets/sounds/LeftClick.mp3")
		decode, _ := mp3.DecodeWithSampleRate(44100, file)
		sound.LeftSound, _ = context.NewPlayer(decode)
	}

	if config.General.EnableRightClickSound {
		file, _ := os.Open("assets/sounds/RightClick.mp3")
		decode, _ := mp3.DecodeWithSampleRate(44100, file)
		sound.RightSound, _ = context.NewPlayer(decode)
	}

	return sound
}
