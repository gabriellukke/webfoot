package main

import (
	"fmt"
	"webfoot/pkg/game"
)

func main() {
	team1 := game.NewTeam("Team A", []game.Player{
		game.NewPlayer("Player 1", 75, 80),
		game.NewPlayer("Player 2", 85, 90),
		game.NewPlayer("Player 3", 80, 85),
	})

	team2 := game.NewTeam("Team B", []game.Player{
		game.NewPlayer("Player 4", 70, 75),
		game.NewPlayer("Player 5", 90, 85),
		game.NewPlayer("Player 6", 78, 88),
	})

	fmt.Println("Welcome to Football Simulation!")
	fmt.Printf("Team %s is ready with %d players.\n", team1.Name, len(team1.Players))
	fmt.Printf("Team %s is ready with %d players.\n", team2.Name, len(team2.Players))

	game.SimulateMatch(team1, team2)
}
