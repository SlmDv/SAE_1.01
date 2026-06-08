package particle

import (
	"math"
	"project-particles/config"
	"reflect"
	"testing"
)

func TestNoArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		1,
		1,
		1,
		1,
		0,
		0,
		0}
	result := NewParticle()
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestOneArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		1,
		1,
		1,
		1,
		0,
		0,
		2}
	result := NewParticle(2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestTwoArg(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		1,
		1,
		1,
		1,
		0,
		0.8,
		2}
	result := NewParticle(0.8, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestThreeArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		1,
		1,
		1,
		1,
		0.9,
		0.8,
		2}
	result := NewParticle(0.9, 0.8, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestFourArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		1,
		1,
		1,
		2,
		0.9,
		0.8,
		2}
	result := NewParticle(2, 0.9, 0.8, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestFiveArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		1,
		1,
		0.5,
		2,
		0.9,
		0.8,
		2}
	result := NewParticle(0.5, 2, 0.9, 0.8, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestSixArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		1,
		0.4,
		0.5,
		2,
		0.9,
		0.8,
		2}
	result := NewParticle(0.4, 0.5, 2, 0.9, 0.8, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestSevenArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		0.3,
		0.4,
		0.5,
		2,
		0.9,
		0.8,
		2}
	result := NewParticle(0.3, 0.4, 0.5, 2, 0.9, 0.8, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestEightArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1.1,
		0.3,
		0.4,
		0.5,
		2,
		0.9,
		0.8,
		2}
	result := NewParticle(1.1, 0.3, 0.4, 0.5, 2, 0.9, 0.8, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestNineArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		2,
		0.9,
		0.8,
		2}
	result := NewParticle(1.2, 1.1, 0.3, 0.4, 0.5, 2, 0.9, 0.8, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestTenArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		math.Pi,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		2,
		0.9,
		0.8,
		2}
	result := NewParticle(math.Pi, 1.2, 1.1, 0.3, 0.4, 0.5, 2, 0.9, 0.8, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestElevenArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY) / 2,
		math.Pi,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		2,
		0.9,
		0.8,
		2}
	result := NewParticle(float64(config.General.WindowSizeY), math.Pi, 1.2, 1.1, 0.3, 0.4, 0.5, 2, 0.9, 0.8, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}
func TestTwelveArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX) / 2,
		float64(config.General.WindowSizeY) / 2,
		math.Pi,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		2,
		0.9,
		0.8,
		2}

	result := NewParticle(float64(config.General.WindowSizeY)/2, float64(config.General.WindowSizeY), math.Pi, 1.2, 1.1, 0.3, 0.4, 0.5, 2, 0.9, 0.8, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}
