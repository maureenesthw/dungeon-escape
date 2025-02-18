package main

func (g *Game) StepOnTrap() {
	g.HealthPoints -= 20
	PrintStrings([]string{"You stepped on a trap! You lost 20 HP!"})
}
