package main

import (
	"fmt"
	"math"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particle"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// notInRange() est une fonction qui revoie un booléen selon que la valeur de value soit entre les valeurs de min et max.
// Elle prend en argument trois float64.
func notInRange(value, min, max float64) bool {
	return value < min || value > max
}

// needToKill() est une fonction utilisée uniquement dans draw.go qui verifie si l'état d'une particule doit necessairement être dessiné par la fonction Draw.
// Elle prend en argument un pointeur vers une particule, et renvoie un booléen valant true si l'intégralité de cette dernière est en dehors de la zone d'affichage définie par les deux variables dans config.json.
// Necessite l'import de math.
func needToKill(p *particle.Particle) bool {
	ecart := math.Sqrt((p.ScaleX*p.ScaleX)+(p.ScaleY*p.ScaleY)) / 2
	margin := config.General.DrawOblivionMargin + ecart
	return notInRange(p.PositionX, 0-margin, float64(config.General.WindowSizeX)+margin) ||
		notInRange(p.PositionY, 0-margin, float64(config.General.WindowSizeY)+margin)
}

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.
// Commentaires supplémentaires :
// La fonction needToKill() est utilisée pour conditionner l'affichage d'une particule à sa présence dans l'écran.
func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particle.Particle)

		if ok {
			if needToKill(p) && config.General.DrawOblivionMargin > 0 {
				continue
			} else {

				w, h := assets.ParticleImage.Size()
				options := ebiten.DrawImageOptions{}
				options.GeoM.Translate((-float64(w) / 2), (-float64(h) / 2))
				options.GeoM.Rotate(p.Rotation)
				options.GeoM.Scale(p.ScaleX, p.ScaleY)
				options.GeoM.Translate(p.PositionX, p.PositionY)
				options.ColorScale.Scale(float32(p.ColorRed), float32(p.ColorGreen), float32(p.ColorBlue), float32(p.Opacity))
				screen.DrawImage(assets.ParticleImage, &options)
			}
		}
	}

	if config.General.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))
	}

	if particles.TheInterface.Visible {
		particles.DrawInterface(screen)
	} else if config.General.EnableInterface {
		ebitenutil.DebugPrintAt(screen, "Press SPACE to show the interface", 5, 15)
	}

}
