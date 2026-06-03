package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"project-particles/particle"
)

// RandomSpeed() est une fonction qui sert simplement à renvoyer un float64 compris entre 1.5 et 6.
// Elle ne prend pas d'arguments.
func RandomSpeed() float64 {
	return 1.5 + rand.Float64()*(4.5)
}

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
// Elle prend en argument un pointeur vers un SoundManager, qui sera utilisé
// pour jouer les sons lors des clics de souris. Elle retourne un System.
// L'import de container/list et math/rand est requis.
func NewSystem(sound *SoundManager) System {
	L := list.New()
	var spawnPointX float64
	var spawnPointY float64
	for i := 0; i < config.General.InitNumParticles; i++ {

		p := particle.NewParticle()

		if config.General.RandomSpawn {
			spawnPointX = float64(rand.Intn(config.General.WindowSizeX))
			spawnPointY = float64(rand.Intn(config.General.WindowSizeY))
		} else {
			spawnPointX = float64(config.General.SpawnX)
			spawnPointY = float64(config.General.SpawnY)
		}

		p = particle.Particle{
			PositionX: spawnPointX,
			PositionY: spawnPointY,
			ScaleX:    scale(),
			ScaleY:    scale(),
			ColorRed:  1, ColorGreen: 1, ColorBlue: 1,
			Opacity:  1,
			SpeedX:   rand.Float64()*15 - 7.5,
			SpeedY:   rand.Float64()*15 - 7.5,
			Rotation: rand.Float64(),
			LifeTime: 0,
		}

		L.PushBack(&p)
	}

	return System{Content: L, Sound: sound}
}
