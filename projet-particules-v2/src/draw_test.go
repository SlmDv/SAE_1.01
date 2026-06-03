package main

import (
	"project-particles/config"
	"project-particles/particle"
	"testing"
)

func TestNotInRange(t *testing.T) {
	if !notInRange(0, 1, 10) {
		t.Fatal("Erreur : 0 devrait être hors intervalle")
	}
	if !notInRange(11, 1, 10) {
		t.Fatal("Erreur : 11 devrait être hors intervalle")
	}
	if notInRange(5, 1, 10) {
		t.Fatal("Erreur : 5 ne devrait pas être hors intervalle")
	}

}

func TestNeedToKill(t *testing.T) {
	config.Get("config.json")
	config.General.WindowSizeX = 100
	config.General.WindowSizeY = 100
	config.General.DrawOblivionMargin = 0

	p := &particle.Particle{
		PositionX: 50,
		PositionY: 50,
		ScaleX:    1,
		ScaleY:    1,
	}
	if needToKill(p) {
		t.Fatal("Erreur : la particule devrait toujours être dessinée.")
	}

	p = &particle.Particle{
		PositionX: -10,
		PositionY: 50,
		ScaleX:    1,
		ScaleY:    1,
	}
	if !needToKill(p) {
		t.Fatal("Erreur : la particule ne devrait plus être dessinée.")
	}

	p = &particle.Particle{
		PositionX: 100,
		PositionY: 100,
		ScaleX:    1,
		ScaleY:    1,
	}
	if needToKill(p) {
		t.Fatal("Erreur : la particule est en bordure, mais devrait toujours être dessinée")
	}
}
