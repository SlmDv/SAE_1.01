package particle

import (
	"math/rand/v2"
	"project-particles/config"
)

// increaseOpacity() est une méthode qui incrémente la variable Opacity d'une particule dont un pointeur associé est donné en argument.
// Elle verifie d'abord que la variable Opacity peut encore subir une modification sans dépasser son maximum.
// Elle prend en receiver un pointeur vers une particule.
func (p *Particle) increaseOpacity() {
	if p.Opacity <= 0.9 {
		p.Opacity += 0.1
	}
}

// decreaseOpacity() est une méthode qui decrémente la variable Opacity d'une particule dont un pointeur associé est donné en argument.
// Elle verifie d'abord que la variable Opacity peut encore subir une modification sans dépasser son minimum.
// Elle prend en receiver un pointeur vers une particule.
func (p *Particle) decreaseOpacity() {
	if p.Opacity >= 0.1 {
		p.Opacity -= 0.1
	}
}

// updateConfitionnelle() est une méthode appliquée à une particule, qui met à jour l'état dynamique d'une particule. Elle utilise plusieurs variables stockées dans config.json.
// De plus, elle appelle les fonctions increaseOpacity() et decreaseOpacity() de manière aléatoire afin de opacifier/désopacifier la particule.
// Elle prend en receiver un pointeur vers une particule.
func (p *Particle) UpdateConditionnelle() {
	p.SpeedY += config.General.Gravity / 360
	p.PositionX += p.SpeedX
	p.PositionY += p.SpeedY
	p.Rotation += config.General.ParticleRotationSpeed / 360
	p.LifeTime++
	if rand.Float64() <= 0.5 && config.General.OpacityVariation {
		p.increaseOpacity()
	} else if rand.Float64() > 0.5 && config.General.OpacityVariation {
		p.decreaseOpacity()
	}
}
