package particle

import (
	"project-particles/config"
)

/*
Renvoie une particule définie par

	PositionX, PositionY            float64 valeur par défaut WindowSizeX, WindowSizeY
	Rotation                        float64 valeur par défaut 0
	ScaleX, ScaleY                  float64 valeur par défaut 1,1
	ColorRed, ColorGreen, ColorBlue float64 valeur par défaut 1,1,1
	Opacity 						float64 valeur par défaut 1

La particule peut ne pas être à l'écran
*/

func NewParticle(optionals ...float64) Particle {
	//définie les valeurs par défaut des paramètres
	params := []float64{float64(config.General.WindowSizeX) / 2, float64(config.General.WindowSizeY) / 2, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0}
	lp := len(params)
	lo := len(optionals)
	for i := 0; i < lo; i++ {
		params[i+lp-lo] = optionals[i]
	}

	return Particle{
		PositionX:  params[0],
		PositionY:  params[1],
		Rotation:   params[2],
		ScaleX:     params[3],
		ScaleY:     params[4],
		ColorRed:   params[5],
		ColorGreen: params[6],
		ColorBlue:  params[7],
		Opacity:    params[8],
		SpeedX:     params[9],
		SpeedY:     params[10],
		LifeTime:   params[11],
	}
}
