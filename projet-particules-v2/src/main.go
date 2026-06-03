package main

import (
	"log"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

// main est la fonction principale du projet. Elle commence par lire le fichier
// de configuration, puis elle charge en mémoire l'image d'une particule. Elle
// initialise ensuite la fenêtre d'affichage, puis elle crée un système de
// particules encapsulé dans un "game" et appelle la fonction RunGame qui se
// charge de faire les mise-à-jour (Update) et affichages (Draw) de manière
// régulière.
// Commentaire supplémentaire :
// ELle initialise également le gestionnaire ebiten des sons en créant un audio.Context.
// De plus, elle initialise l'interface dynamique du programme.

func main() {

	config.Get("config.json")
	assets.Get()

	audioContext := audio.NewContext(44100)
	sound := particles.NewSoundManager(audioContext)

	particles.InitInterface()

	ebiten.SetWindowTitle(config.General.WindowTitle)
	ebiten.SetWindowSize(config.General.WindowSizeX, config.General.WindowSizeY)

	g := game{system: particles.NewSystem(sound)}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}
}
