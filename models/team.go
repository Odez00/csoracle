package models

import "encoding/json"

// Team represents a CS:GO team with their stats
type Team struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Wins      int     `json:"wins"`
	Losses    int     `json:"losses"`
	Rating    float64 `json:"rating"`
	WorldRank int     `json:"world_rank"`
}

// WinRate returns the team's win rate as a percentage
func (t *Team) WinRate() float64 {
	total := t.Wins + t.Losses
	if total == 0 {
		return 0.0
	}
	return (float64(t.Wins) / float64(total)) * 100.0
}

// TotalMatches returns the total number of matches played
func (t *Team) TotalMatches() int {
	return t.Wins + t.Losses
}

// UnmarshalTeams creates a slice of Team structs from JSON data
func UnmarshalTeams(data []byte) ([]Team, error) {
	var teams []Team
	err := json.Unmarshal(data, &teams)
	return teams, err
}

// UpdateElo updates teams elo based on teams wins and losses
func UpdateElo(teams *[]Team) {
	for i := 0; i < len(*teams); i++ {
		eloPlus := 12.0 * (*teams)[i].Wins
		(*teams)[i].Rating += float64(eloPlus)
		eloMinus := 12.0 * float64((*teams)[i].Losses)
		(*teams)[i].Rating -= float64(eloMinus)
	}
}
