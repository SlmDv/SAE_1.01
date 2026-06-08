package main

import (
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
// Commentaires supplémentaires :
// La gestion de l'interface est ajoutée ici. Si l'interface est
// activée dans le fichier config.json et que l'utilisateur presse la
// touche Espace, l'interface s'affiche/se masque. Si l'interface est visible,
// sa fonction UpdateInterface() est appelée à chaque frame.

func (g *game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) && config.General.EnableInterface {
		particles.TheInterface.SwitchInterface()
	}
	if particles.TheInterface.Visible && config.General.EnableInterface {
		particles.TheInterface.UpdateInterface()
	}
	g.system.Update()
	return nil
}
