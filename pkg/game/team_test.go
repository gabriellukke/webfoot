package game_test

import (
	"testing"
	"webfoot/pkg/game"
)

func TestNewTeam(t *testing.T) {
	players := []game.Player{
		game.NewPlayer("Player 1", 75, 80),
		game.NewPlayer("Player 2", 85, 90),
	}

	team := game.NewTeam("Team A", players)

	if team.Name != "Team A" {
		t.Errorf("Expected Team Name 'Team A', got '%s'", team.Name)
	}
	if len(team.Players) != 2 {
		t.Errorf("Expected 2 players, got %d", len(team.Players))
	}
}
