package particles

import (
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestSwitchInterface(t *testing.T) {
	x := &superInterface{Visible: false}
	x.SwitchInterface()

	if !x.Visible {
		t.Fatalf("Erreur : le booléen Visible devrait être true")
	}

	x.SwitchInterface()

	if x.Visible {
		t.Fatalf("Erreur : le booléen Visible devrait être false")
	}
}

func TestInitInterface(t *testing.T) {
	InitInterface()

	if TheInterface == nil {
		t.Fatalf("Erreur : l'interface n'est pas correctement initialisé")
	}

	if TheInterface.Visible != false {
		t.Fatalf("Erreur : La valeur initiale de Visible devrait être false")
	}

	if TheInterface.Config == nil {
		t.Fatalf("Erreur : la variable Config ne devrait pas être nil")
	}
}

func TestIsOnButton(t *testing.T) {
	if !isOnButton(15, 25, 10, 30, 20, 40) {
		t.Fatalf("Erreur : les coordonnées ne sont pas reconnues comme valides")
	}

	if !isOnButton(10, 20, 10, 30, 20, 40) {
		t.Fatalf("Erreur : les coordonnées sont au bord, mais doivent quand même être valides.")
	}

	if isOnButton(99, 99, 10, 30, 20, 40) {
		t.Fatalf("Erreur : les coordonnées sont reconnues comme valides.")
	}
}

// Note : la fonction UpdateInterface n'est pas testée ici car elle dépend de l'état du clavier et de la souris,
// et nous ne parvenons pas à trouver une solution acceptable. Nous prenons le parti de dire que les tests de isOnButton ci-dessus valident le comportement principal.

func TestDrawButton_Error(t *testing.T) {
	screen := ebiten.NewImage(800, 600)

	DrawButton(screen, 50, 80, 80, 40, "Test", color.RGBA{100, 100, 100, 100})
}

func TestDrawInterface_Error(t *testing.T) {
	screen := ebiten.NewImage(800, 600)

	DrawInterface(screen)
}
