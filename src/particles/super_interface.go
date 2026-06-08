package particles

import (
	"image/color"
	"project-particles/config"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// superInterface{} est une structure délimitant les caractéritstiques de l'interface utilisateur qui s'affichera pendant le jeu, et permettra d'influer sur les valeurs du config.json
// Cette structure se compose d'un booléen Visible, et d'un pointeur vers la structure config.
type superInterface struct {
	Visible bool
	Config  *config.Config
}

var TheInterface *superInterface

// SwitchInterface() est une méthode appliquée à une instance d'un SuperInterface.
// Lorsqu'elle est appelée, elle change la valeur du booléen Visible, et détermine donc si l'interface s'affiche ou non.
// Prend en receiver une instance d'un SuperInterface défini ci-dessus.
func (a *superInterface) SwitchInterface() {
	a.Visible = !a.Visible
}

// InitInterface() est une fonction qui initialise l'interface de jeu.
// Elle crée simplement une instance de superInterface, assignée à la variable TheInterface.
func InitInterface() {
	TheInterface = &superInterface{
		Visible: false,
		Config:  &config.General,
	}
}

// isOnButton() est une fonction qui renvoie un booléen selon que les coordonnées x et y passées en argument se trouvent à l'intérieur des limites définies par xmin, xmax, ymin et ymax.
func isOnButton(x, y int, xmin, xmax, ymin, ymax int) bool {
	return x >= xmin && x <= xmax && y >= ymin && y <= ymax
}

// UpdateInterface() est une méthode appliquée à une instance de SuperInterface.
// Elle detecte si un clic gauche est entré au clavier, si c'est le cas, selon les coordonnées du clic, le bouton qui s'y trouve s'active et modifie config.json.
// Prend en receiver une instance d'un SuperInterface défini ci-dessus.
func (a *superInterface) UpdateInterface() {
	x, y := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if isOnButton(x, y, 50, 130, 80, 120) {
			a.Config.RandomSpawn = !a.Config.RandomSpawn
		}
		if isOnButton(x, y, 150, 250, 80, 120) {
			a.Config.OpacityVariation = !a.Config.OpacityVariation
		}
		if isOnButton(x, y, 270, 385, 80, 120) {
			a.Config.RightClickSpawning = !a.Config.RightClickSpawning
		}
		if isOnButton(x, y, 405, 535, 80, 120) {
			a.Config.LeftClickMovesSpawner = !a.Config.LeftClickMovesSpawner
		}
		if isOnButton(x, y, 555, 680, 80, 120) {
			a.Config.EnableLeftClickSound = !a.Config.EnableLeftClickSound
		}
		if isOnButton(x, y, 225, 355, 140, 180) {
			a.Config.EnableRightClickSound = !a.Config.EnableRightClickSound
		}
		if isOnButton(x, y, 425, 515, 140, 180) {
			a.Config.MemoryThanksYou = !a.Config.MemoryThanksYou
		}

		if isOnButton(x, y, 70, 90, 500, 520) {
			a.Config.SpawnRate += 25
		}
		if isOnButton(x, y, 100, 120, 500, 520) {
			if a.Config.SpawnRate > 25 {
				a.Config.SpawnRate -= 25
			} else {
				a.Config.SpawnRate = 0
			}
		}

		if isOnButton(x, y, 220, 240, 500, 520) {
			a.Config.ScaleVariation += 0.5
		}
		if isOnButton(x, y, 250, 270, 500, 520) {
			a.Config.ScaleVariation -= 0.5
		}

		if isOnButton(x, y, 370, 390, 500, 520) {
			a.Config.RotorSpeed += 5
		}

		if isOnButton(x, y, 400, 420, 500, 520) {
			if a.Config.RotorSpeed <= 5 {
				a.Config.RotorSpeed = 0
			} else {
				a.Config.RotorSpeed -= 5
			}
		}

		if isOnButton(x, y, 520, 540, 500, 520) {
			a.Config.Gravity += 5
		}
		if isOnButton(x, y, 550, 570, 500, 520) {
			a.Config.Gravity -= 5
		}

		if isOnButton(x, y, 670, 690, 500, 520) {
			a.Config.DeathTime += 50
		}

		if isOnButton(x, y, 700, 720, 500, 520) {
			if a.Config.DeathTime <= 50 {
				a.Config.DeathTime = 1
			} else {
				a.Config.DeathTime -= 50
			}
		}
	}
}

// DrawButton() est une fonction qui dessine un bouton sur un screen donné en argument. Le bouton a pour longueur et hauteur w et h, et se positionne aux coordonnées x et y données en argument.
// Le bouton prend un nom et une couleur, définis par les deux derniers arguments.
func DrawButton(screen *ebiten.Image, x, y, w, h float64, nom string, rgb color.Color) {
	img := ebiten.NewImage(int(w), int(h))
	img.Fill(rgb)
	params := &ebiten.DrawImageOptions{}
	params.GeoM.Translate(x, y)
	screen.DrawImage(img, params)
	if w != 20 {
		ebitenutil.DebugPrintAt(screen, nom, int(x+2), int(y+10))
	} else {
		ebitenutil.DebugPrintAt(screen, nom, int(x+6), int(y+3))
	}
}

// DrawInterface() est une fonction qui dessine l'interface sur le screen donné en argument.
// Pour chaque bouton souhaité sur cet interface, elle appelle DrawButton(). Enfin, elle affiche les intitulés de ces boutons et les valeurs actuelles des variables du config.json.
// Elle prend en argument un pointeur vers une image ebiten.
func DrawInterface(screen *ebiten.Image) {
	red := color.RGBA{200, 100, 100, 200}
	blue := color.RGBA{100, 100, 200, 200}

	DrawButton(screen, 50, 80, 80, 40, "RandomSpawn", red)
	DrawButton(screen, 150, 80, 100, 40, "OpacityVariation", red)
	DrawButton(screen, 270, 80, 115, 40, "RightClickSpawning", red)
	DrawButton(screen, 405, 80, 130, 40, "LeftClickMovesSpawner", red)
	DrawButton(screen, 555, 80, 125, 40, "EnableLeftClickSound", red)
	DrawButton(screen, 225, 140, 130, 40, "EnableRightClickSound", red)
	DrawButton(screen, 425, 140, 98, 40, "MemoryThanksYou", red)

	DrawButton(screen, 70, 500, 20, 20, "+", blue)
	DrawButton(screen, 100, 500, 20, 20, "-", blue)
	DrawButton(screen, 220, 500, 20, 20, "+", blue)
	DrawButton(screen, 250, 500, 20, 20, "-", blue)
	DrawButton(screen, 370, 500, 20, 20, "+", blue)
	DrawButton(screen, 400, 500, 20, 20, "-", blue)
	DrawButton(screen, 520, 500, 20, 20, "+", blue)
	DrawButton(screen, 550, 500, 20, 20, "-", blue)
	DrawButton(screen, 670, 500, 20, 20, "+", blue)
	DrawButton(screen, 700, 500, 20, 20, "-", blue)

	ebitenutil.DebugPrintAt(screen, "SpawnRate: "+strconv.FormatFloat(config.General.SpawnRate, 'f', 2, 64), 70, 480)
	ebitenutil.DebugPrintAt(screen, "ScaleVariation: "+strconv.FormatFloat(config.General.ScaleVariation, 'f', 2, 64), 220, 480)
	ebitenutil.DebugPrintAt(screen, "RotorSpeed: "+strconv.FormatUint(uint64(config.General.RotorSpeed), 10), 370, 480)
	ebitenutil.DebugPrintAt(screen, "Gravity: "+strconv.FormatFloat(config.General.Gravity, 'f', 2, 64), 520, 480)
	ebitenutil.DebugPrintAt(screen, "DeathTime: "+strconv.FormatFloat(config.General.DeathTime, 'f', 2, 64), 670, 480)

	ebitenutil.DebugPrintAt(screen, strconv.FormatBool(config.General.RandomSpawn), 60, 100)
	ebitenutil.DebugPrintAt(screen, strconv.FormatBool(config.General.OpacityVariation), 170, 100)
	ebitenutil.DebugPrintAt(screen, strconv.FormatBool(config.General.RightClickSpawning), 290, 100)
	ebitenutil.DebugPrintAt(screen, strconv.FormatBool(config.General.LeftClickMovesSpawner), 430, 100)
	ebitenutil.DebugPrintAt(screen, strconv.FormatBool(config.General.EnableLeftClickSound), 580, 100)
	ebitenutil.DebugPrintAt(screen, strconv.FormatBool(config.General.EnableRightClickSound), 235, 160)
	ebitenutil.DebugPrintAt(screen, strconv.FormatBool(config.General.MemoryThanksYou), 435, 160)
}
