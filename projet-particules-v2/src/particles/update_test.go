package particles

import (
	"container/list"
	"math"
	"project-particles/config"
	"project-particles/particle"
	"testing"
)

func TestNextCoeffRotation_Increment(t *testing.T) {
	config.Get("../config.json")
	CurrentAngle = 0
	initialValue := NextCoeffRotation()
	secondValue := NextCoeffRotation()

	if secondValue <= initialValue {
		t.Fatal("Attendu : angle croissant, fourni :", initialValue, "puis", secondValue)
	}
}

func TestNextCoeffRotation_WrapAround(t *testing.T) {
	config.Get("../config.json")
	CurrentAngle = 2*math.Pi - 0.01
	result := NextCoeffRotation()

	if result >= 2*math.Pi {
		t.Fatal("Attendu : angle < 2π, fourni :", result)
	}
}

func TestNextCoeffRotation_Range(t *testing.T) {
	config.Get("../config.json")
	CurrentAngle = 0

	for i := 0; i < 100; i++ {
		result := NextCoeffRotation()
		if result < 0 || result >= 2*math.Pi {
			t.Fatalf("Attendu : angle entre 0 et 2π, fourni : %f", result)
		}
	}
}

func TestRandomSpawnSpeed_Range(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := RandomSpawnSpeed()
		if result < -5 || result >= 5 {
			t.Fatalf("Attendu : un flottant entre -5 inclus et 5 exclu, fourni : %f", result)
		}
	}

}

func TestRandomSpawnSpeed_Diversity(t *testing.T) {
	a, b, c := RandomSpawnSpeed(), RandomSpawnSpeed(), RandomSpawnSpeed()
	if a == b && b == c {
		t.Fatalf("Attendu : trois valeurs générées différentes, fourni : %f, %f, %f", a, b, c)
	}

}

func TestRotationSpeed_ZeroAngle(t *testing.T) {
	vx, vy := RotationSpeed(5.0, 3.0, 0)

	if math.Abs(vx-5.0) > 0.001 || math.Abs(vy-3.0) > 0.001 {
		t.Fatal("Attendu : (5.0, 3.0) pour angle 0, fourni :", vx, vy)
	}
}

func TestRotationSpeed_PiOverTwo(t *testing.T) {
	vx, vy := RotationSpeed(1.0, 0.0, math.Pi/2)

	if math.Abs(vx) > 0.001 || math.Abs(vy-1.0) > 0.001 {
		t.Fatal("Attendu : (0, 1.0) pour angle π/2, fourni :", vx, vy)
	}
}

func TestRotationSpeed_Consistency(t *testing.T) {
	a, b := 4.0, 2.0
	angle := math.Pi / 4

	vx1, vy1 := RotationSpeed(a, b, angle)
	vx2, vy2 := RotationSpeed(a, b, angle)

	if math.Abs(vx1-vx2) > 0.001 || math.Abs(vy1-vy2) > 0.001 {
		t.Fatal("Attendu : résultats identiques pour mêmes paramètres, fourni :", vx1, vy1, "vs", vx2, vy2)
	}
}

func TestChromaticCircle_Zero(t *testing.T) {
	result := chromaticCircle(0)

	if math.Abs(result-1.0) > 0.001 {
		t.Fatal("Attendu : 1.0 pour angle 0, fourni :", result)
	}
}

func TestChromaticCircle_Sixty(t *testing.T) {
	result := chromaticCircle(60)

	if math.Abs(result) > 0.001 {
		t.Fatal("Attendu : 0 pour angle 60, fourni :", result)
	}
}

func TestChromaticCircle_Range(t *testing.T) {
	for angle := -180.0; angle <= 180.0; angle += 10 {
		result := chromaticCircle(angle)
		if result < 0 || result > 1 {
			t.Fatalf("Attendu : valeur entre 0 et 1 pour angle %f, fourni : %f", angle, result)
		}
	}
}

func TestChromaticCircle_WrapAround(t *testing.T) {
	result1 := chromaticCircle(180)
	result2 := chromaticCircle(-180)

	if math.Abs(result1-result2) > 0.001 {
		t.Fatal("Attendu : résultats identiques pour 180 et -180, fourni :", result1, result2)
	}
}

func TestSpawnRate_WholeNumber(t *testing.T) {
	config.Get("../config.json")
	config.General.SpawnRate = 5.0
	counter = 0

	s := System{Content: list.New()}
	initialLen := s.Content.Len()
	SpawnRate(&s, 0)

	if s.Content.Len() != initialLen+5 {
		t.Fatalf("Attendu : %d particules ajoutées, fourni : %d", 5, s.Content.Len()-initialLen)
	}
}

func TestSpawnRate_Fractional(t *testing.T) {
	config.Get("../config.json")
	config.General.SpawnRate = 2.5
	counter = 0

	s := System{Content: list.New()}
	initialLen := s.Content.Len()
	SpawnRate(&s, 0)

	if s.Content.Len() != initialLen+2 && s.Content.Len() != initialLen+3 {
		t.Fatalf("Attendu : 2 ou 3 particules ajoutées, fourni : %d", s.Content.Len()-initialLen)
	}
}

func TestSpawnRate_CounterAccumulation(t *testing.T) {
	config.Get("../config.json")
	config.General.SpawnRate = 0.7
	counter = 0

	s := System{Content: list.New()}
	initialLen := s.Content.Len()
	SpawnRate(&s, 0)
	SpawnRate(&s, 0)

	totalAdded := (s.Content.Len() - initialLen)
	if totalAdded < 1 || totalAdded > 2 {
		t.Fatalf("Attendu : 1 ou 2 particules après 2 appels, fourni : %d", totalAdded)
	}
}

func TestGeneratePosition_RandomSpawnValues(t *testing.T) {
	config.Get("../config.json")
	config.General.RandomSpawn = true

	for i := 0; i < 100; i++ {
		x, y := generatePosition(0)
		if (x < 0 || x > float64(config.General.WindowSizeX)) || (y < 0 || y > float64(config.General.WindowSizeY)) {
			t.Fatalf("Attendu : des coordonnées d'apparition entre 0 et les dimensions de la fenêtre, fourni : %f, %f", x, y)
		}
	}

}

func TestGeneratePosition_NoRandomSpawnValues(t *testing.T) {
	config.Get("../config.json")
	config.General.RandomSpawn = false
	angle := 0.0
	for i := 0; i < 100; i++ {
		v, w := generatePosition(angle)
		x, y := generatePosition(angle + float64(1))
		angle++

		if v == x || w == y {
			t.Fatalf("Attendu : une évolution des coordonnées d'apparition selon l'angle donné, mais on retrouve, pour deux angles différents : %f, %f, et : %f %f", v, x, w, y)
		}
	}

}

func TestScale(t *testing.T) {
	config.General.ScaleVariation = 0
	if scale() != 0.8 {
		t.Fatal("Erreur : scale() devrait renvoyer 0.8 quand ScaleVariation <= 0")
	}
	config.General.ScaleVariation = 2
	if scale() != 2 {
		t.Fatal("Erreur : scale() devrait renvoyer exactement 2 quand ScaleVariation = 2")
	}
	config.General.ScaleVariation = 5
	scale := scale()
	if scale < 2 || scale >= 5 {
		t.Fatalf("Erreur : scale() devrait être dans l'intervalle[2, 5), obtenu %f", scale)
	}
}

func TestSpawnParticle_GenerationTest(t *testing.T) {
	config.Get("../config.json")
	s := &System{
		Content: list.New(),
	}
	angle := 0.0

	s.spawnParticle(angle)
	if s.Content.Len() != 1 {
		t.Fatalf("La fonction spawnParticle ne génère pas de nouvel élément")
	}

}

func TestSpawnParticle_TypeTest(t *testing.T) {
	config.Get("../config.json")
	s := &System{
		Content: list.New(),
	}
	angle := 0.0

	s.spawnParticle(angle)
	p := s.Content.Front()
	if p == nil {
		t.Fatalf("Erreur : aucune particule ajoutée")
	}
	if _, ok := p.Value.(*particle.Particle); !ok {
		t.Fatalf("Erreur : l'element ajouté n'est pas un *particle.Particle")
	}

}

func TestSpawnParticle_Position(t *testing.T) {
	config.Get("../config.json")
	s := &System{
		Content: list.New(),
	}
	angle := math.Pi * 2
	x, y := generatePosition(angle)

	s.spawnParticle(angle)
	p := s.Content.Front().Value.(*particle.Particle)
	if p.PositionX != x || p.PositionY != y {
		t.Fatalf("Erreur de coordonnées de position.")
	}
}

func TestSpawnParticle_SpeedWhileRandomSpawnTrue(t *testing.T) {
	config.Get("../config.json")
	s := &System{
		Content: list.New(),
	}
	angle := 0.0
	config.General.RandomSpawn = true

	for i := 0; i < 100; i++ {
		s.spawnParticle(angle)
		p := s.Content.Front().Value.(*particle.Particle)
		if (p.SpeedX < -5 || p.SpeedX >= 5) || (p.SpeedY < -5 || p.SpeedY >= 5) {
			t.Fatalf("Erreur : une ou plusieurs valeurs de vitesse ne correspondent pas aux valeurs attendues selon la configuration : %f, %f", p.SpeedX, p.SpeedY)
		}

	}

}

func TestSpawnParticle_SpeedWhileRandomSpawnFalse(t *testing.T) {
	config.Get("../config.json")
	s := &System{
		Content: list.New(),
	}
	angle := 0.0
	config.General.RandomSpawn = false

	for i := 0; i < 100; i++ {
		s.spawnParticle(angle)
		p := s.Content.Front().Value.(*particle.Particle)
		vx, vy := p.SpeedX, p.SpeedY
		norme := math.Sqrt(vx*vx + vy*vy)
		if norme < 2.12 || norme > 8.49 {
			t.Fatalf("Erreur : une ou plusieurs valeurs de vitesse ne correspondent pas aux valeurs attendues selon la configuration : %f, %f", vx, vy)
		}

	}
}

func TestSpawnParticle_CorrectScale(t *testing.T) {
	config.Get("../config.json")
	s := &System{
		Content: list.New(),
	}
	angle := 0.0

	for i := 0; i < 100; i++ {
		s.spawnParticle(angle)
		p := s.Content.Front().Value.(*particle.Particle)
		if (p.ScaleX < 0.75 || p.ScaleX > 1.5) || (p.ScaleY < 0.75 || p.ScaleY > 1.5) {
			t.Fatalf("Erreur : La taille des particules n'est pas conforme, longueur : %f, hauteur : %f", p.ScaleX, p.ScaleY)
		}
	}

}

func TestTimeToDie_TooSoonTooLate(t *testing.T) {
	config.Get("../config.json")
	s := &System{
		Content: list.New(),
	}
	angle := 0.0

	s.spawnParticle(angle)
	p := s.Content.Front().Value.(*particle.Particle)
	for i := 0; i < int(config.General.DeathTime); i++ {
		timeLeft := int(config.General.DeathTime - p.LifeTime)
		if timeToDie(p) {
			t.Fatalf("Erreur : la particule est considérée comme ''morte'' %d appels avant d'avoir atteint la valeur spécifiée dans la configuration.", timeLeft)
		}
		p.UpdateConditionnelle()

	}
	if !timeToDie(p) {
		t.Fatalf("Erreur: la particule devrait être morte après %d appels", int(config.General.DeathTime))
	}
}

func TestSystemUpdate_ParticlesUpdated(t *testing.T) {
	config.Get("../config.json")
	s := NewSystem(nil)

	if s.Content.Len() == 0 {
		t.Fatal("Erreur : Système vide, impossible de tester")
	}

	firstParticle := s.Content.Front().Value.(*particle.Particle)
	initialX := firstParticle.PositionX
	initialY := firstParticle.PositionY

	s.Update()

	if math.Abs(firstParticle.PositionX-initialX) < 0.001 && math.Abs(firstParticle.PositionY-initialY) < 0.001 {
		t.Fatal("Attendu : une particule aux coordonnées mises à jour, fourni : des particules aux positions sont inchangées")
	}
}

func TestSystemUpdate_NewParticlesSpawned(t *testing.T) {
	config.Get("../config.json")
	s := System{Content: list.New()}
	initialLen := s.Content.Len()

	s.Update()

	if s.Content.Len() <= initialLen {
		t.Fatalf("Attendu : nouvelles particules ajoutées, longueur : %d -> %d", initialLen, s.Content.Len())
	}
}
