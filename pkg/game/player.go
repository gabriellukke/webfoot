package game

type Player struct {
	Name    string
	Skill   int // Skill level, from 1 to 100
	Stamina int // Stamina level, from 1 to 100
}

func NewPlayer(name string, skill, stamina int) Player {
	return Player{Name: name, Skill: skill, Stamina: stamina}
}
