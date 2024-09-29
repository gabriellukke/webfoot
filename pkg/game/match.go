package game

import (
	"fmt"
	"math/rand"
	"time"
)

func SimulateMatch(team1, team2 Team) (team1Score, team2Score int) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	team1Score = r.Intn(team1.AvgSkill() / 10)
	team2Score = r.Intn(team2.AvgSkill() / 10)

	fmt.Printf("Match Result: %s %d - %d %s\n", team1.Name, team1Score, team2Score, team2.Name)

	return team1Score, team2Score
}
