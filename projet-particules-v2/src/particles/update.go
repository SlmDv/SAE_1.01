package particles

import (
	"container/list"
	"math"
	"math/rand"
	"project-particles/config"
	"project-particles/particle"

	"github.com/hajimehoshi/ebiten/v2"
)

var CurrentAngle float64 = 0
var counter float64 = 0

// NextCoeffRotation() est une fonction qui renvoie un angle incrémenté à chaque appel par la valeur VitesseRotationMoulin, stockée dans config.json.
// L'angle se réinitialise lorsqu'il atteint 2*Pi, permettant de réinitialiser le motif circulaire que cette fonction contribuera à dessiner.
// l'import de math est requis.
func NextCoeffRotation() float64 {

	CurrentAngle += float64(config.General.RotorSpeed) / 360
	if CurrentAngle >= 2*math.Pi {
		CurrentAngle = 0
	}
	return CurrentAngle
}

// RandomSpawnSpeed() est une fonction qui genère et renvoie une vitesse aléatoire comprise entre -5 et 5.
// Utilisée pour donner une direction aléatoire aux particules lorsque la variable RandomSpawn de la config est true.
// Elle ne prend pas d'arguments en entrée.
// l'import de math/rand est requis.
func RandomSpawnSpeed() float64 {
	return -5 + rand.Float64()*(5-(-5))
}

// RotationSpeed() est une fonction qui prends trois valeurs et en résulte deux coefficients.
// Ces coefficients seront pris dans le calcul d'une orientation circulaire des vitesses des particules.
// L'import de math est requis.
func RotationSpeed(a, b float64, angle float64) (float64, float64) {
	cosA := math.Cos(angle)
	sinA := math.Sin(angle)
	return a*cosA - b*sinA, a*sinA + b*cosA
}

// chromaticCircle() est une fonction qui renvoie des coefficients compris entre 0 et 1, avec trois cas en fonction de la valeur de l'angle donné en paramètre.
// Ces coefficients seront utilisées pour generer un dégradé de couleur, chaque particule a une couleur qui dépend de la couleur donnée au précédent appel.
// L'import de math est requis.
func chromaticCircle(angle float64) float64 {
	if config.General.EnableColors {
		angle = math.Mod((angle+180), 360) - 180
		if math.Abs(angle) <= 60 {
			return 1 - math.Abs(angle)/60
		} else if 120 < math.Abs(angle) && math.Abs(angle) <= 180 {
			return (math.Abs(angle) - 120) / 60
		} else {
			return 0
		}
	} else {
		return 1
	}
}

// SpawnRate() est une fonction qui permet de faire spawner un nombre (SpawnRate) de particules a chaque frame, et gère un compteur interne pour prévoir des spawnrates non-entiers.
// Elle prend en argument un pointeur vers un System, ainsi qu'un angle float64.
func SpawnRate(s *System, angle float64) {
	rate := config.General.SpawnRate + counter
	whole := int(rate)
	counter = rate - float64(whole)

	for i := 0; i < whole; i++ {
		s.spawnParticle(angle)
	}
}

// generatePosition() est une fonction qui renvoie des coordonnées de génération.
// Selon la valeur du booléen RandomSpawn stocké dans config.json, les coordonées seront :
// -aléatoires au sein de la fenêtre.
// -en un point précis de la fenêtre.
// Dans le deuxième cas, l'application d'un coefficient modélisé par math.Cos(angle)*20 permet que le point d'apparition s'incrive dans un motif circulaire.
// L'import de math est requis.
func generatePosition(angle float64) (x float64, y float64) {
	if config.General.RandomSpawn {
		x = float64(rand.Intn(config.General.WindowSizeX))
		y = float64(rand.Intn(config.General.WindowSizeY))
	} else {
		x = float64(config.General.SpawnX) + math.Cos(angle)*20
		y = float64(config.General.SpawnY) + math.Sin(angle)*20
	}
	return
}

// Scale() est une fonction qui renvoie une valeur float64 correspondant à l'échelle d'une particule.
// Si la variable ScaleVariation dans config.json est supérieure à 0, l'échelle sera une valeur aléatoire entre 2 et ScaleVariation.
// Sinon, l'échelle sera fixée à 0.8.
func scale() float64 {
	var Scale float64
	if config.General.ScaleVariation > 0 {
		Scale = 2 + rand.Float64()*(config.General.ScaleVariation-2)
	} else {
		Scale = 0.8
	}
	return Scale
}

// spawnParticule() est une méthode qui, à chaque appel, genère un nouvel element de type particle.Particle et l'ajoute à la fin de la liste s.Content.
// Cette méthode s'applique à un système de particules, et prend en argument un angle float64.
// Les caractéristiques SpawnX et SpawnY de la particule sont conditionnées à la valeur du booléen RandomSpawn, stocké dans config.json.
// Les caractéristiques de SpeedX et SpeedY de la particule sont conditionnées à la valeur de RotorSpeed, stockée dans config.json.
// L'import de math est requis.
func (s *System) spawnParticle(angle float64) {
	var vx, vy float64

	if !config.General.RandomSpawn && config.General.RotorSpeed != 0 {
		vx, vy = RotationSpeed(RandomSpeed(), RandomSpeed(), angle)
	} else {
		vx = RandomSpawnSpeed()
		vy = RandomSpawnSpeed()
	}

	angleDegre := math.Mod(angle*90/math.Pi+360, 360)
	spawnPointX, spawnPointY := generatePosition(angle)

	generee := particle.Particle{
		PositionX:  spawnPointX,
		PositionY:  spawnPointY,
		ScaleX:     scale(),
		ScaleY:     scale(),
		ColorRed:   chromaticCircle(angleDegre),
		ColorGreen: chromaticCircle(angleDegre - 120),
		ColorBlue:  chromaticCircle(angleDegre + 120),
		Opacity:    0.5,
		SpeedX:     vx,
		SpeedY:     vy,
		Rotation:   rand.Float64(),
		LifeTime:   0,
	}

	s.Content.PushBack(&generee)
}

// timeToDie() est une fonction qui renvoie un booléen selon que la variable LifeTime d'une particule atteint ou dépasse la variable DeathTime stockée dans config.json.
// Elle prend en argument un pointeur vers une Particle.
func timeToDie(p *particle.Particle) bool {
	return p.LifeTime >= config.General.DeathTime
}

// Update met à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
// Commentaires supplémentaires :
// Elle prend en argument un pointeur vers un System. Pour des raisons de performances, la gestion des particules à supprimer se fait à l'aide d'une seconde liste doublement chaînée.
// De plus, elle gère l'implementation des extensions de déplacement du spawner avec le clic gauche, et de spawn additionnel de particules avec le clic droit.
func (s *System) Update() {

	angle := NextCoeffRotation()
	var toRemove *list.List = list.New()
	enabledRemove := config.General.MemoryThanksYou

	for uneParticule := s.Content.Front(); uneParticule != nil; uneParticule = uneParticule.Next() {
		p, ok := uneParticule.Value.(*particle.Particle)
		if timeToDie(p) && config.General.DeathTime > 0 {
			if enabledRemove {
				toRemove.PushFront(uneParticule)
			}
			continue
		}
		if ok {
			p.UpdateConditionnelle()
		}
	}
	SpawnRate(s, angle)
	if enabledRemove {
		for uneCell := toRemove.Front(); uneCell != nil; uneCell = uneCell.Next() {
			s.Content.Remove(uneCell.Value.(*list.Element))
		}
	}

	if !TheInterface.Visible {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && config.General.LeftClickMovesSpawner {
			config.General.SpawnX, config.General.SpawnY = ebiten.CursorPosition()
			if config.General.EnableLeftClickSound && s.Sound.LeftSound != nil {
				s.Sound.LeftSound.Rewind()
				s.Sound.LeftSound.Play()
			}
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) && config.General.RightClickSpawning {
		x, y := ebiten.CursorPosition()
		if config.General.EnableRightClickSound && s.Sound.RightSound != nil {
			s.Sound.RightSound.Rewind()
			s.Sound.RightSound.Play()
		}

		for i := 0; i < config.General.InitNumParticles; i++ {

			p := particle.Particle{
				PositionX: float64(x),
				PositionY: float64(y),
				ScaleX:    scale(),
				ScaleY:    scale(),
				ColorRed:  rand.Float64(), ColorGreen: rand.Float64(), ColorBlue: rand.Float64(),
				Opacity:  0.5,
				SpeedX:   rand.Float64()*15 - 7.5,
				SpeedY:   rand.Float64()*15 - 7.5,
				Rotation: rand.Float64(),
				LifeTime: 0,
			}

			s.Content.PushFront(&p)
		}
	}
}
