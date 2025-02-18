package main

func (g *Game) FindExitDoor() {
	if g.KeyObtained {
		input := GetUserInput("You found the exit door! Do you want to (E)scape?")
		if input == "E" {
			g.IsWin = true
		}
	} else {
		PrintStrings([]string{"The door is locked! Keep searching!"})
	}
}
