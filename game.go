package main

import (
	"fmt"
	"math/rand"
)

type Game struct {
	Moves        int
	HealthPoints int
	KeyObtained  bool
	IsWin        bool
}

func StartNewGame() Game {
	return Game{
		Moves:        InitialMoves,
		HealthPoints: InitialHealthPoints,
		KeyObtained:  false,
		IsWin:        false,
	}
}

func (g *Game) Play() {
	PrintStrings([]string{
		"Welcome to Dungeon Escape!",
		fmt.Sprintf("You have %dHP and %d moves to escape.", InitialHealthPoints, InitialMoves),
		"\n",
	})

	for g.Moves > 0 {
		chosenDoor := ""
		for chosenDoor != "L" && chosenDoor != "R" && chosenDoor != "F" {
			chosenDoor = GetUserInput("Choose a door: (L)eft, (R)ight, (F)orward")
		}

		g.Move()
		if g.HealthPoints <= 0 {
			PrintStrings([]string{"You lost all your HP! Game over!"})
			break
		}
		if g.IsWin {
			PrintStrings([]string{fmt.Sprintf("You escaped the dungeon in %d moves! You win!", InitialMoves-g.Moves)})
			break
		}

		PrintStrings([]string{"", fmt.Sprintf("Moves left: %d | HP: %d", g.Moves, g.HealthPoints), "----------------------------------------", "\n"})
	}

	if g.Moves == 0 {
		PrintStrings([]string{"You ran out of moves! Game over!"})
	}
}

func (g *Game) Move() {
	g.Moves--

	chance := rand.Intn(99)
	switch {
	case chance < 35:
		// 35% chance of monster encounter
		g.MonsterEncounter()
	case chance < 55:
		// 20% chance of encountering trap
		g.StepOnTrap()
	case chance < 70:
		// 15% chance of finding exit door
		g.FindExitDoor()
	default:
		// 30% chance of finding a key
		g.FindAKey()
	}
}
