package dbstore

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/itemslabs/clubz-api/database"
	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CreatePlayerBucket creates a new player bucket in the database with detailed error handling.
func (s *DBStore) CreatePlayerBucket(pb *schema.PlayerBucket) (*schema.PlayerBucket, error) {
	if err := pb.Insert(s.db, boil.Infer()); err != nil {
		return nil, fmt.Errorf("failed to insert player bucket: %w", err)
	}
	return pb, nil
}

// GetPlayerBucketByID retrieves a player bucket by its ID with detailed error handling.
func (s *DBStore) GetPlayerBucketByID(id string) (*schema.PlayerBucket, error) {
	pb, err := schema.FindPlayerBucket(s.db, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find player bucket by ID %s: %w", id, err)
	}
	return pb, nil
}

// GetAllPlayerBuckets retrieves all player buckets from the database with detailed error handling.
func (s *DBStore) GetAllPlayerBuckets() (schema.PlayerBucketSlice, error) {
	pbs, err := schema.PlayerBuckets().All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all player buckets: %w", err)
	}
	return pbs, nil
}

// UpdatePlayerBucket updates an existing player bucket with detailed error handling.
func (s *DBStore) UpdatePlayerBucket(pb *schema.PlayerBucket) error {
	_, err := pb.Update(s.db, boil.Whitelist("first_name", "last_name", "team_id", "updated_at"))
	if err != nil {
		return fmt.Errorf("failed to update player bucket: %w", err)
	}
	return nil
}

// DeletePlayerBucket deletes a player bucket by its ID with detailed error handling.
func (s *DBStore) DeletePlayerBucket(id string) error {
	pb, err := schema.FindPlayerBucket(s.db, id)
	if err != nil {
		return fmt.Errorf("failed to find player bucket for deletion: %w", err)
	}
	_, err = pb.Delete(s.db)
	if err != nil {
		return fmt.Errorf("failed to delete player bucket: %w", err)
	}
	return nil
}

// GetPlayerBucketsByTeamID retrieves all player buckets for a given team ID with detailed error handling.
func (s *DBStore) GetPlayerBucketsByTeamID(teamID string) (schema.PlayerBucketSlice, error) {
	pbs, err := schema.PlayerBuckets(
		qm.Where("team_id = ?", teamID),
	).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve player buckets by team ID %s: %w", teamID, err)
	}
	return pbs, nil
}

// AssignPlayerToTeam updates or creates a new player bucket with a specific team assignment with detailed error handling.
func (s *DBStore) AssignPlayerToTeam(playerID, teamID string) (*schema.PlayerBucket, error) {
	pb := &schema.PlayerBucket{
		ID:     playerID,
		TeamID: teamID,
	}
	if err := pb.Upsert(s.db, true, []string{"id"}, boil.Whitelist("team_id"), boil.Infer()); err != nil {
		return nil, fmt.Errorf("failed to upsert player bucket: %w", err)
	}
	return pb, nil
}

// CountPlayerBuckets returns the total number of player buckets with detailed error handling.
func (s *DBStore) CountPlayerBuckets() (int64, error) {
	count, err := schema.PlayerBuckets().Count(s.db)
	if err != nil {
		return 0, fmt.Errorf("failed to count player buckets: %w", err)
	}
	return count, nil
}

// CountPlayersByTeam counts the number of players in each team.
func (s *DBStore) CountPlayersByTeam(ctx context.Context) (map[string]int, error) {
	type result struct {
		TeamID string
		Count  int
	}
	var results []result

	// Adjusted call to Bind with context.Context as the first argument
	err := schema.PlayerBuckets(
		qm.Select(schema.PlayerBucketColumns.TeamID, "COUNT(*)"),
		qm.GroupBy(schema.PlayerBucketColumns.TeamID),
	).Bind(ctx, s.db, &results) // Now providing context as the first argument
	if err != nil {
		return nil, fmt.Errorf("failed to count players by team: %w", err)
	}

	counts := make(map[string]int)
	for _, r := range results {
		counts[r.TeamID] = r.Count
	}
	return counts, nil
}

// GetPlayersByStatus retrieves all players with a specific achievement status.
func (s *DBStore) GetPlayersByStatus(status string) (schema.PlayerBucketSlice, error) {
	var mods []qm.QueryMod
	switch status {
	case "Bronze":
		mods = append(mods, schema.PlayerBucketWhere.Bronze.EQ(null.BoolFrom(true)))
	case "Silver":
		mods = append(mods, schema.PlayerBucketWhere.Silver.EQ(null.BoolFrom(true)))
	case "Gold":
		mods = append(mods, schema.PlayerBucketWhere.Gold.EQ(null.BoolFrom(true)))
	case "Platinum":
		mods = append(mods, schema.PlayerBucketWhere.Platinum.EQ(null.BoolFrom(true)))
	case "Diamond":
		mods = append(mods, schema.PlayerBucketWhere.Diamond.EQ(null.BoolFrom(true)))
	default:
		return nil, fmt.Errorf("invalid status: %s", status)
	}

	pbs, err := schema.PlayerBuckets(mods...).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve player buckets by %s status: %w", status, err)
	}
	return pbs, nil
}

// GetPlayersByTeamAndStatus retrieves all players from a specific team with a specific achievement status.
func (s *DBStore) GetPlayersByTeamAndStatus(teamID, status string) (schema.PlayerBucketSlice, error) {
	var mods []qm.QueryMod
	mods = append(mods, qm.Where("team_id = ?", teamID))
	switch status {
	case "Bronze":
		mods = append(mods, schema.PlayerBucketWhere.Bronze.EQ(null.BoolFrom(true)))
	case "Silver":
		mods = append(mods, schema.PlayerBucketWhere.Silver.EQ(null.BoolFrom(true)))
	case "Gold":
		mods = append(mods, schema.PlayerBucketWhere.Gold.EQ(null.BoolFrom(true)))
	case "Platinum":
		mods = append(mods, schema.PlayerBucketWhere.Platinum.EQ(null.BoolFrom(true)))
	case "Diamond":
		mods = append(mods, schema.PlayerBucketWhere.Diamond.EQ(null.BoolFrom(true)))
	default:
		return nil, fmt.Errorf("invalid status: %s", status)
	}

	pbs, err := schema.PlayerBuckets(mods...).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve player buckets by team ID %s and %s status: %w", teamID, status, err)
	}
	return pbs, nil
}

func (s *DBStore) CalculateRarityPercentages() (*database.RarityPercentages, error) {
	// Retrieve the sum of all rarity caps
	var totalCap int
	var rarityCaps struct {
		Common    int
		Uncommon  int
		Rare      int
		UltraRare int
		Legendary int
	}

	err := s.db.QueryRow(`
        SELECT
            SUM(common) AS common,
            SUM(uncommon) AS uncommon,
            SUM(rare) AS rare,
            SUM(ultra_rare) AS ultra_rare,
            SUM(legendary) AS legendary
        FROM player_bucket
    `).Scan(&rarityCaps.Common, &rarityCaps.Uncommon, &rarityCaps.Rare, &rarityCaps.UltraRare, &rarityCaps.Legendary)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate total caps: %w", err)
	}

	totalCap = rarityCaps.Common + rarityCaps.Uncommon + rarityCaps.Rare + rarityCaps.UltraRare + rarityCaps.Legendary

	// Calculate percentages
	rp := &database.RarityPercentages{
		Common:    float64(rarityCaps.Common) / float64(totalCap) * 100,
		Uncommon:  float64(rarityCaps.Uncommon) / float64(totalCap) * 100,
		Rare:      float64(rarityCaps.Rare) / float64(totalCap) * 100,
		UltraRare: float64(rarityCaps.UltraRare) / float64(totalCap) * 100,
		Legendary: float64(rarityCaps.Legendary) / float64(totalCap) * 100,
	}

	return rp, nil
}

func (s *DBStore) GetPlayersByRarity(rarity string) (schema.PlayerBucketSlice, error) {
	var mods []qm.QueryMod

	switch rarity {
	case "common":
		mods = append(mods, qm.Where("common > 0"))
	case "uncommon":
		mods = append(mods, qm.Where("uncommon > 0"))
	case "rare":
		mods = append(mods, qm.Where("rare > 0"))
	case "ultra_rare":
		mods = append(mods, qm.Where("ultra_rare > 0"))
	case "legendary":
		mods = append(mods, qm.Where("legendary > 0"))
	default:
		return nil, fmt.Errorf("invalid rarity: %s", rarity)
	}

	return schema.PlayerBuckets(mods...).All(s.db)
}

// GetRandomPlayerByTeamAndRarity retrieves a random player from a specific team with a specified rarity and a cap greater than 0.
func (s *DBStore) GetRandomPlayerByTeamAndRarity(teamID string, rarity string) (*schema.PlayerBucket, error) {
	var mods []qm.QueryMod
	mods = append(mods, qm.Where("team_id = ?", teamID))

	// Adjusting for boolean rarity flags
	switch rarity {
	case "common":
		mods = append(mods, qm.Where("common = TRUE"))
	case "uncommon":
		mods = append(mods, qm.Where("uncommon = TRUE"))
	case "rare":
		mods = append(mods, qm.Where("rare = TRUE"))
	case "ultra_rare":
		mods = append(mods, qm.Where("ultra_rare = TRUE"))
	case "legendary":
		mods = append(mods, qm.Where("legendary = TRUE"))
	default:
		return nil, fmt.Errorf("invalid rarity: %s", rarity)
	}

	// Retrieve all matching players
	players, err := schema.PlayerBuckets(mods...).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve players by team ID %s and %s rarity: %w", teamID, rarity, err)
	}

	if len(players) == 0 {
		return nil, fmt.Errorf("no players found for team ID %s with %s rarity", teamID, rarity)
	}

	// Select a random player from the list
	randomIndex := rand.Intn(len(players))
	return players[randomIndex], nil
}

// GetRandomPlayerBucketByTeamID retrieves a random player bucket for a given team ID.
func (s *DBStore) GetRandomPlayerBucketByTeamID(teamID string) (*schema.PlayerBucket, error) {
	// Retrieve all player buckets for the specified team ID
	players, err := schema.PlayerBuckets(
		qm.Where("team_id = ?", teamID),
	).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve player buckets by team ID %s: %w", teamID, err)
	}

	if len(players) == 0 {
		return nil, fmt.Errorf("no players found for team ID %s", teamID)
	}

	// Select a random player from the list
	randomIndex := rand.Intn(len(players))
	return players[randomIndex], nil
}
