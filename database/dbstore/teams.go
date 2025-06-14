package dbstore

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (s *DBStore) GetTeamByID(teamID string) (*schema.Team, error) {
	return schema.Teams(
		qm.Where("id = ?", teamID),
	).One(s.db)
}

func (s *DBStore) GetUniqueRandomTeams(count int) (schema.TeamSlice, error) {
	initialTeams, err := schema.Teams(
		qm.OrderBy("RANDOM()"),
		qm.Limit(count*3),
	).All(s.db)
	if err != nil {
		return nil, err
	}

	rand.Shuffle(len(initialTeams), func(i, j int) { initialTeams[i], initialTeams[j] = initialTeams[j], initialTeams[i] })

	uniqueTeams := make(schema.TeamSlice, 0, count)
	names := make(map[string]struct{})

	for _, team := range initialTeams {
		if _, exists := names[team.Name]; exists {
			continue
		}

		uniqueTeams = append(uniqueTeams, team)
		names[team.Name] = struct{}{}

		if len(uniqueTeams) == count {
			break
		}
	}
	if len(uniqueTeams) < count {
		return uniqueTeams, fmt.Errorf("only found %d unique teams out of requested %d", len(uniqueTeams), count)
	}

	return uniqueTeams, nil
}

func (s *DBStore) AddTeam(team *schema.Team) error {
	return team.Insert(s.db, boil.Infer())
}

func (s *DBStore) UpdateTeam(team *schema.Team) (int64, error) {
	return team.Update(s.db, boil.Whitelist("name", "short_name", "country_id", "region_id", "crest_url", "abbr", "ortec_selection_id"))
}

func (s *DBStore) TeamExistsByID(teamID string) (bool, error) {
	return schema.TeamExists(s.db, teamID)
}

func (s *DBStore) TeamExistsByName(name string) (bool, *schema.Team, error) {
	team, err := schema.Teams(qm.Where("name = ?", name)).One(s.db)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil, nil
		}
		return false, nil, err
	}
	return true, team, nil
}

func (s *DBStore) ValidateTeam(team *schema.Team) error {
	exists, _, err := s.TeamExistsByName(team.Name)
	if err != nil {
		return fmt.Errorf("error checking team name existence: %w", err)
	}
	if exists {
		return fmt.Errorf("team name '%s' already exists", team.Name)
	}
	return nil
}

func (s *DBStore) GetOrCreateTeamByName(teamName string) (*schema.Team, error) {
	team, err := schema.Teams(qm.Where("name=?", teamName)).One(s.db)
	if err == nil {
		return team, nil // Team found
	}

	// Assuming no team found, create a new one.
	newTeam := &schema.Team{
		Name: teamName,
		// Set other necessary team fields here.
	}
	err = newTeam.Insert(s.db, boil.Infer())
	if err != nil {
		return nil, fmt.Errorf("failed to create new team '%s': %w", teamName, err)
	}
	return newTeam, nil
}

func (s *DBStore) GetTeamByName(teamName string) (*schema.Team, error) {
	return schema.Teams(qm.Where("name=?", teamName)).One(s.db)
}

func (s *DBStore) GetUniqueRandomTeamExcluding(usedTeams map[string]struct{}) (*schema.Team, error) {
	var teamIDs []string // Using a slice of strings to hold team IDs

	// Initialize the base query
	query := `SELECT team_id FROM nft_bucket`

	var args []interface{}
	if len(usedTeams) > 0 {
		var placeholders []string
		i := 1 // Start indexing placeholders from 1 for databases like PostgreSQL
		for id := range usedTeams {
			placeholders = append(placeholders, fmt.Sprintf("$%d", i))
			args = append(args, id)
			i++
		}

		// Append the condition to exclude used teams
		query += fmt.Sprintf(" WHERE team_id NOT IN (%s)", strings.Join(placeholders, ","))
	} else {
		// Get all distinct team IDs from nft_bucket
		query = `SELECT DISTINCT team_id FROM nft_bucket`
	}

	// Execute the query
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Scan each row and append the team_id to teamIDs
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("failed to scan team_id: %w", err)
		}
		teamIDs = append(teamIDs, id)
	}

	// Check for errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	if len(teamIDs) == 0 {
		return nil, fmt.Errorf("no unique teams available")
	}

	// Select a random team from the slice
	randomIndex := rand.Intn(len(teamIDs))
	selectedTeamID := teamIDs[randomIndex]

	// Use the team ID to get the full team schema
	teamSchema, err := s.GetTeamByID(selectedTeamID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve team with ID %s: %w", selectedTeamID, err)
	}

	return teamSchema, nil
}
