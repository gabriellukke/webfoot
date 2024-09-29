package game_test

import (
	"testing"
	"webfoot/pkg/game"
)

func TestSimulateMatch(t *testing.T) {
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

	// Simulate a match
	team1Score, team2Score := game.SimulateMatch(team1, team2)

	// Check if the scores are within a reasonable range
	if team1Score < 0 || team2Score < 0 {
		t.Errorf("Scores cannot be negative, got: %d - %d", team1Score, team2Score)
	}

	// Check if a draw is possible
	if team1Score == team2Score {
		t.Logf("The match ended in a draw: %d - %d", team1Score, team2Score)
	}

	// Check for team 1 win
	if team1Score > team2Score {
		t.Logf("Team 1 wins: %d - %d", team1Score, team2Score)
	}

	// Check for team 2 win
	if team2Score > team1Score {
		t.Logf("Team 2 wins: %d - %d", team1Score, team2Score)
	}
}

func TestSimulateMultipleMatches(t *testing.T) {
	team1 := game.NewTeam("Team A", []game.Player{
		game.NewPlayer("Player 1", 80, 80),
		game.NewPlayer("Player 2", 85, 90),
		game.NewPlayer("Player 3", 75, 85),
	})

	team2 := game.NewTeam("Team B", []game.Player{
		game.NewPlayer("Player 4", 70, 75),
		game.NewPlayer("Player 5", 65, 85),
		game.NewPlayer("Player 6", 60, 88),
	})

	for i := 0; i < 10; i++ {
		team1Score, team2Score := game.SimulateMatch(team1, team2)

		// Ensure the scores are non-negative
		if team1Score < 0 || team2Score < 0 {
			t.Errorf("Scores cannot be negative, got: %d - %d", team1Score, team2Score)
		}

		// Test that results are within a reasonable range, based on the AvgSkill
		if team1Score > 8 || team2Score > 8 {
			t.Errorf("Scores are too high, got: %d - %d", team1Score, team2Score)
		}
	}
}
