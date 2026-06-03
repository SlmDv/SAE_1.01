package particles

import (
	"project-particles/config"
	"project-particles/particle"
	"testing"
)

func TestRandomSpeed_Limits(t *testing.T) {
	res := RandomSpeed()
	if res >= 6 || res < 1.5 {

		t.Fatal("Attendu : entre 1.5 et 6", " fourni :", res)
	}
}

func TestRandomSpeed_Randomness(t *testing.T) {
	x1 := RandomSpeed()
	x2 := RandomSpeed()
	x3 := RandomSpeed()
	if x1 == x2 && x2 == x3 {
		t.Fatal("Attendu : trois nombres différents", " fourni :", x1, ",", x2, "et", x3)
	}
}

func TestNewSystem_InitNumParticles(t *testing.T) {
	config.Get("../config.json")
	system := NewSystem(nil)

	if system.Content.Len() != config.General.InitNumParticles {
		t.Fatalf("Attendu : %d particules, fourni : %d",
			config.General.InitNumParticles, system.Content.Len())
	}
}

func TestNewSystem_FixedSpawn(t *testing.T) {
	config.Get("../config.json")
	config.General.RandomSpawn = false
	system := NewSystem(nil)

	for e := system.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*particle.Particle)
		if p.PositionX != float64(config.General.SpawnX) || p.PositionY != float64(config.General.SpawnY) {
			t.Errorf("Position attendue : (%d, %d), fournie : (%f, %f)",
				config.General.SpawnX, config.General.SpawnY, p.PositionX, p.PositionY)
		}
	}
}

func TestNewSystem_RandomSpawn(t *testing.T) {
	config.Get("../config.json")
	config.General.RandomSpawn = true
	system := NewSystem(nil)

	for e := system.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*particle.Particle)
		if p.PositionX < 0 || p.PositionX >= float64(config.General.WindowSizeX) {
			t.Errorf("PositionX hors limites : %f (attendu entre 0 et %d)",
				p.PositionX, config.General.WindowSizeX)
		}
		if p.PositionY < 0 || p.PositionY >= float64(config.General.WindowSizeY) {
			t.Errorf("PositionY hors limites : %f (attendu entre 0 et %d)",
				p.PositionY, config.General.WindowSizeY)
		}
	}
}

func TestNewSystem_ParticleProperties(t *testing.T) {
	config.Get("../config.json")
	system := NewSystem(nil)

	for e := system.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*particle.Particle)

		if p.ScaleX < 0.75 || p.ScaleX > 1.5 {
			t.Errorf("ScaleX hors limites : %f (attendu entre 0.75 et 1.5)", p.ScaleX)
		}
		if p.ScaleY < 0.75 || p.ScaleY > 1.5 {
			t.Errorf("ScaleY hors limites : %f (attendu entre 0.75 et 1.5)", p.ScaleY)
		}
		if p.Opacity != 1.0 {
			t.Errorf("Opacity attendu : 1.0, fourni : %f", p.Opacity)
		}
		if p.Rotation < 0 || p.Rotation > 1 {
			t.Errorf("Rotation hors limites : %f (attendu entre 0 et 1)", p.Rotation)
		}
	}
}
