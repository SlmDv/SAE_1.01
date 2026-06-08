package particle

import (
	"math"
	"project-particles/config"
	"testing"
)

func TestIncreaseOpacity(t *testing.T) {
	p := &Particle{Opacity: 0.5}

	p.increaseOpacity()
	if math.Abs(p.Opacity-0.6) > 0.001 {
		t.Fatalf("Attendu: Opacity = 0.6, fourni : %f", p.Opacity)
	}
}

func TestDecreaseOpacity(t *testing.T) {
	p := &Particle{Opacity: 0.5}

	p.decreaseOpacity()
	if math.Abs(p.Opacity-0.4) > 0.001 {
		t.Fatalf("Attendu: Opacity = 0.4, fourni : %f", p.Opacity)
	}
}

func TestUpdateConditionnelle_PositionUpdate(t *testing.T) {
	config.Get("../config.json")
	config.General.RandomSpawn = false
	expectedSpeedY := 3.0 + config.General.Gravity/360
	expectedPositionY := 20.0 + expectedSpeedY

	p := &Particle{
		PositionX: 10.0,
		PositionY: 20.0,
		SpeedX:    2.0,
		SpeedY:    3.0,
		Rotation:  0.5,
	}

	p.UpdateConditionnelle()

	if math.Abs(p.PositionX-12.0) > 0.001 {
		t.Fatalf("Attendu PositionX : 12.0, fourni : %f", p.PositionX)
	}
	if math.Abs(p.PositionY-expectedPositionY) > 0.001 {
		t.Fatalf("Attendu PositionY : %f, fourni : %f", expectedPositionY, p.PositionY)
	}
}

func TestUpdateConditionnelle_RotationUpdate(t *testing.T) {
	config.Get("../config.json")
	initialRotation := 0.5
	p := &Particle{
		Rotation: initialRotation,
	}

	p.UpdateConditionnelle()

	if p.Rotation <= initialRotation {
		t.Fatal("Attendu : rotation augmentée, fourni :", p.Rotation)
	}
}

func TestUpdateConditionnelle_GravityWhenNotRandomSpawn(t *testing.T) {
	config.Get("../config.json")
	config.General.RandomSpawn = false
	config.General.Gravity = 360.0
	initialSpeedY := 0.0
	p := &Particle{
		SpeedY: initialSpeedY,
	}

	p.UpdateConditionnelle()

	if p.SpeedY <= initialSpeedY {
		t.Fatalf("Attendu : SpeedY augmentée par la gravité, fourni : vitesse initiale : %f, vitesse d'arrivée : %f", initialSpeedY, p.SpeedY)
	}
}

func TestUpdateConditionnelle_NoGravityWhenRandomSpawn(t *testing.T) {
	config.Get("../config.json")
	config.General.RandomSpawn = true
	config.General.Gravity = 0
	initialSpeedY := 5.0
	p := &Particle{
		SpeedY: initialSpeedY,
		SpeedX: 2.0,
	}

	p.UpdateConditionnelle()

	if math.Abs(p.SpeedY-initialSpeedY) > 0.001 {
		t.Fatalf("Attendu : SpeedY inchangée (pas de gravité), fourni : %f", p.SpeedY)
	}
}

func TestUpdateConditionnelle_LifeTimeIncrement(t *testing.T) {
	config.Get("../config.json")
	p := &Particle{
		LifeTime: 999,
	}

	p.UpdateConditionnelle()

	if p.LifeTime != 1000 {
		t.Fatalf("Erreur : la méthode n'incrémente pas correctement la durée de vie de la particule")
	}

}
