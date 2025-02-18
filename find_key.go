package main

func (g *Game) FindAKey() {
	g.KeyObtained = true
	PrintStrings([]string{"You found a key! Now you can escape if you find the exit!"})
}
