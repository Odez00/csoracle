package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Odez00/csoracle/models"
	"github.com/Odez00/csoracle/predictor"
)

func main() {
	fmt.Println("CS Match Predictor")
	fmt.Println("==================")

	// Load teams from JSON file
	teams, err := loadTeams("data/teams.json")
	if err != nil {
		log.Fatalf("Error loading teams: %v", err)
	}

	if len(teams) == 0 {
		fmt.Println("No teams found. Please add teams to data/teams.json")
		return
	}

	// Update team ELO ratings
	models.UpdateElo(&teams)

	// Create predictor
	pred := predictor.New(teams)

	// List available teams
	fmt.Println("\nAvailable teams:")
	teamNames := pred.ListTeams()
	for _, name := range teamNames {
		fmt.Printf("  - %s\n", name)
	}

	// Example prediction
	if len(teams) >= 2 {
		team1Name := teams[0].Name
		team2Name := teams[3].Name

		fmt.Printf("\nPredicting match: %s vs %s\n", team1Name, team2Name)
		fmt.Println("--------------------------------")

		result, err := pred.PredictMatch(team1Name, team2Name)
		if err != nil {
			log.Fatalf("Error predicting match: %v", err)
		}

		// Display prediction
		fmt.Printf("Predicted Winner: %s\n", result.PredictedWinner)
		fmt.Printf("Confidence: %s\n", result.Confidence)
		fmt.Printf("\nWin Chances:\n")
		fmt.Printf("  %s: %.2f%%\n", result.Team1Name, result.Team1WinChance)
		fmt.Printf("  %s: %.2f%%\n", result.Team2Name, result.Team2WinChance)
	}
}

// loadTeams loads team data from a JSON file
func loadTeams(filepath string) ([]models.Team, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// Handle empty file
	if len(data) == 0 {
		return []models.Team{}, nil
	}

	var teams []models.Team
	err = json.Unmarshal(data, &teams)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return teams, nil
}
