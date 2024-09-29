package game_test

import (
	"testing"
	"webfoot/pkg/game"
)

func TestNewPlayer(t *testing.T) {
	player := game.NewPlayer("Player 1", 80, 90)
	if player.Name != "Player 1" {
		t.Errorf("Expected player name to be 'Player 1', got %s", player.Name)
	}
}
