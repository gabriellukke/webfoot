package game

type Team struct {
	Name    string
	Players []Player
}

func NewTeam(name string, players []Player) Team {
	return Team{Name: name, Players: players}
}

func (t Team) AvgSkill() int {
	totalSkill := 0
	for _, player := range t.Players {
		totalSkill += player.Skill
	}
	return totalSkill / len(t.Players)
}
