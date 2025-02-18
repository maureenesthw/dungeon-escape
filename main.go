package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	gameSession := StartNewGame()
	gameSession.Play()
}
