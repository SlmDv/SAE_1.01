package particles

import (
	"project-particles/config"
	"testing"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

var context = audio.NewContext(44100)

func TestNewSoundManager_Base(t *testing.T) {
	sound := NewSoundManager(context)

	if sound == nil {
		t.Fatal("Erreur : sound ne devrait pas être nil")
	}

	if sound.Context != context {
		t.Fatal("Erreur : la fonction ne recupère pas l'audio fourni")
	}
}

func TestNewSoundManager_LeftSoundEnabled(t *testing.T) {
	config.General.EnableLeftClickSound = true
	config.General.EnableRightClickSound = false

	sound2 := NewSoundManager(context)

	if sound2.LeftSound == nil {
		t.Fatal("Erreur : le son du clic gauche devrait être activé")
	}
	if sound2.RightSound != nil {
		t.Fatal("Erreur : le son du clic droit devrait être désactivé")
	}
}

func TestNewSoundManager_RightSoundEnabled(t *testing.T) {
	config.General.EnableLeftClickSound = false
	config.General.EnableRightClickSound = true

	sound3 := NewSoundManager(context)

	if sound3.RightSound == nil {
		t.Fatal("Erreur : le son du clic droit devrait être activé")
	}
	if sound3.LeftSound != nil {
		t.Fatal("Erreur : le son du clic gauche devrait être désactivé")
	}
}
