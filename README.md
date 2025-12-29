# CS Match Predictor

A rudimentary Go-based application for predicting CS:GO match results based on team statistics.

## Features

- **Team Management**: Load team data from JSON files
- **Match Prediction**: Predict match outcomes based on team ratings and win rates
- **Confidence Levels**: High, Medium, or Low confidence based on prediction strength
- **Elo-like Algorithm**: Uses a formula similar to Elo ratings for win probability calculation

## How It Works

The prediction algorithm considers:
1. **Team Rating**: A numerical rating representing overall team strength
2. **Win Rate**: Calculated from wins, losses, and draws
3. **Elo-like Formula**: Calculates win probability using the formula:
   ```
   P(A) = 1 / (1 + 10^((R_B - R_A) / 400))
   ```

## Running the Application

```bash
go run main.go
```

## Adding Teams

Edit `data/teams.json` to add or modify teams. Each team requires:
- `id`: Unique identifier
- `name`: Team name
- `region`: Geographic region
- `wins`: Number of wins
- `losses`: Number of losses
- `draws`: Number of draws
- `rating`: Team rating (higher is better)
- `world_rank`: Current world ranking

## Example Output

```
CS Match Predictor
==================

Available teams:
  - Natus Vincere
  - G2 Esports
  - FaZe Clan
  - Team Vitality
  - Astralis

Predicting match: Natus Vincere vs G2 Esports
--------------------------------
Predicted Winner: Team 1
Confidence: Medium

Win Chances:
  Natus Vincere: 60.07%
  G2 Esports: 39.93%
```

## Future Improvements

- Interactive CLI for selecting teams
- Historical match data integration
- Machine learning-based predictions
- Web API interface
- Real-time data fetching from HLTV or other sources
- Player statistics consideration
- Map-specific predictions
