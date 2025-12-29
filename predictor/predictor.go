package predictor

import (
	"fmt"
	"math"

	"github.com/Odez00/csoracle/models"
)

// MatchPrediction contains the prediction result
type MatchPrediction struct {
	Team1Name       string  `json:"team1_name"`
	Team2Name       string  `json:"team2_name"`
	Team1WinChance  float64 `json:"team1_win_chance"`
	Team2WinChance  float64 `json:"team2_win_chance"`
	PredictedWinner string  `json:"predicted_winner"`
	Confidence      string  `json:"confidence"`
}

// Predictor handles match predictions
type Predictor struct {
	teams []models.Team
}

// New creates a new Predictor with the given teams
func New(teams []models.Team) *Predictor {
	return &Predictor{
		teams: teams,
	}
}

// PredictMatch predicts the outcome of a match between two teams
func (p *Predictor) PredictMatch(team1Name, team2Name string) (*MatchPrediction, error) {
	team1, err := p.findTeam(team1Name)
	if err != nil {
		return nil, fmt.Errorf("team not found: %s", team1Name)
	}

	team2, err := p.findTeam(team2Name)
	if err != nil {
		return nil, fmt.Errorf("team not found: %s", team2Name)
	}

	// Calculate win chances using a simple rating-based formula
	team1Chance, team2Chance := p.calculateWinChances(team1, team2)

	// Determine predicted winner and confidence
	predictedWinner, confidence := p.determinePrediction(team1Chance, team2Chance)

	return &MatchPrediction{
		Team1Name:       team1Name,
		Team2Name:       team2Name,
		Team1WinChance:  team1Chance,
		Team2WinChance:  team2Chance,
		PredictedWinner: predictedWinner,
		Confidence:      confidence,
	}, nil
}

// findTeam searches for a team by name
func (p *Predictor) findTeam(name string) (*models.Team, error) {
	for i := range p.teams {
		if p.teams[i].Name == name {
			return &p.teams[i], nil
		}
	}
	return nil, fmt.Errorf("team not found")
}

// calculateWinChances calculates win chances based on team ratings and win rates
func (p *Predictor) calculateWinChances(team1, team2 *models.Team) (float64, float64) {
	// Use a combination of rating and win rate
	team1Score := team1.Rating + (team1.WinRate() / 10.0)
	team2Score := team2.Rating + (team2.WinRate() / 10.0)

	// Calculate expected win probability using Elo-like formula
	// P(A) = 1 / (1 + 10^((R_B - R_A) / 400))
	team1WinProb := 1.0 / (1.0 + math.Pow(10.0, (team2Score-team1Score)/400.0))
	team2WinProb := 1.0 - team1WinProb

	return team1WinProb * 100.0, team2WinProb * 100.0
}

// determinePrediction determines the predicted winner and confidence level
func (p *Predictor) determinePrediction(team1Chance, team2Chance float64) (string, string) {
	if team1Chance > team2Chance {
		return "Team 1", p.getConfidenceLevel(team1Chance)
	} else if team2Chance > team1Chance {
		return "Team 2", p.getConfidenceLevel(team2Chance)
	}
	return "Draw", "Low"
}

// getConfidenceLevel returns a confidence level based on win chance
func (p *Predictor) getConfidenceLevel(winChance float64) string {
	if winChance >= 70.0 {
		return "High"
	} else if winChance >= 55.0 {
		return "Medium"
	}
	return "Low"
}

// ListTeams returns all available teams
func (p *Predictor) ListTeams() []string {
	names := make([]string, len(p.teams))
	for i, t := range p.teams {
		names[i] = t.Name
	}
	return names
}
