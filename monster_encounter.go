package main

import (
	"math/rand"
)

func (g *Game) MonsterEncounter() {
	choice := ""
	for choice != "F" && choice != "R" {
		choice = GetUserInput("Do you want to (F)ight or (R)un away? ")
	}
	if choice == "F" {
		g.FightMonster()
	} else if choice == "R" {
		g.RunAway()
	}
}

func (g *Game) FightMonster() {
	chance := rand.Intn(99)
	if chance < 50 {
		g.HealthPoints -= 30
		PrintStrings([]string{"You were hit by the monster! You lost 30 HP!"})
	} else {
		PrintStrings([]string{"You defeated the monster! No damage taken."})
	}
}

func (g *Game) RunAway() {
	chance := rand.Intn(99)
	if chance < 30 {
		g.HealthPoints -= 15
		PrintStrings([]string{"You failed to run away! You lost 15 HP!"})
	}
	PrintStrings([]string{"You successfully ran away! No damage taken."})
}
