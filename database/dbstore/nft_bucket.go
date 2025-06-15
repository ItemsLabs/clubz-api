package dbstore

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/types"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CreateNFTBucket creates a new NFT bucket in the database with detailed error handling.
func (s *DBStore) CreateNFTBucket(nb *schema.NFTBucket) (*schema.NFTBucket, error) {
	if err := nb.Insert(s.db, boil.Infer()); err != nil {
		return nil, fmt.Errorf("failed to insert NFT bucket: %w", err)
	}
	return nb, nil
}

func (s *DBStore) GetNFTBucketByID(id string) (*schema.NFTBucket, error) {
	nb, err := schema.FindNFTBucket(s.db, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find NFT bucket by ID %s: %w", id, err)
	}
	return nb, nil
}

// GetNFTBucketByID retrieves an NFT bucket by its ID and filters the player info based on their NFT rarity.
func (s *DBStore) GetNFTBucketByIDandRarity(id string, rarity string) (*schema.NFTBucket, error) {
	nb, err := schema.FindNFTBucket(s.db, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find NFT bucket by ID %s: %w", id, err)
	}

	// Filter the player info based on the rarity
	var filteredNFTs []*schema.NFTBucket
	switch rarity {
	case "common":
		filteredNFTs = append(filteredNFTs, filterCommonRarity(nb))
	case "rare":
		filteredNFTs = append(filteredNFTs, filterRareRarity(nb))
	case "legendary":
		filteredNFTs = append(filteredNFTs, filterLegendaryRarity(nb))
	case "ultra_rare":
		filteredNFTs = append(filteredNFTs, filterUltraRareRarity(nb))
	case "uncommon":
		filteredNFTs = append(filteredNFTs, filterUncommonRarity(nb))
	case "Common":
		filteredNFTs = append(filteredNFTs, filterCommonRarity(nb))
	case "Uncommon":
		filteredNFTs = append(filteredNFTs, filterUncommonRarity(nb))
	case "Rare":
		filteredNFTs = append(filteredNFTs, filterRareRarity(nb))
	case "UltraRare":
		filteredNFTs = append(filteredNFTs, filterUltraRareRarity(nb))
	case "Legendary":
		filteredNFTs = append(filteredNFTs, filterLegendaryRarity(nb))
	default:
		return nil, fmt.Errorf("invalid rarity: %s", rarity)
	}

	// Check if filteredNFTs has any elements
	if len(filteredNFTs) == 0 {
		return nil, fmt.Errorf("no NFTs found for ID %s with rarity %s", id, rarity)
	}

	// Return the first (filtered) NFT bucket
	return filteredNFTs[0], nil
}

// filterCommonRarity filters the NFT to only include common rarity information.
func filterCommonRarity(nft *schema.NFTBucket) *schema.NFTBucket {
	return &schema.NFTBucket{
		ID:                 nft.ID,
		Name:               nft.Name,
		TeamID:             nft.TeamID,
		Age:                nft.Age,
		GamePosition:       nft.GamePosition,
		Position:           nft.Position,
		CommonClaiming:     nft.CommonClaiming,
		CommonDefence:      nft.CommonDefence,
		CommonDistribution: nft.CommonDistribution,
		CommonDribbling:    nft.CommonDribbling,
		CommonPassing:      nft.CommonPassing,
		CommonShooting:     nft.CommonShooting,
		CommonStopping:     nft.CommonStopping,
		CommonImage:        nft.CommonImage,
		CommonMetadata:     nft.CommonMetadata,
	}
}

// filterRareRarity filters the NFT to only include rare rarity information.
func filterRareRarity(nft *schema.NFTBucket) *schema.NFTBucket {
	return &schema.NFTBucket{
		ID:               nft.ID,
		Name:             nft.Name,
		TeamID:           nft.TeamID,
		Age:              nft.Age,
		GamePosition:     nft.GamePosition,
		Position:         nft.Position,
		RareClaiming:     nft.RareClaiming,
		RareDefence:      nft.RareDefence,
		RareDistribution: nft.RareDistribution,
		RareDribbling:    nft.RareDribbling,
		RarePassing:      nft.RarePassing,
		RareShooting:     nft.RareShooting,
		RareStopping:     nft.RareStopping,
		RareImage:        nft.RareImage,
		RareMetadata:     nft.RareMetadata,
	}
}

// filterLegendaryRarity filters the NFT to only include legendary rarity information.
func filterLegendaryRarity(nft *schema.NFTBucket) *schema.NFTBucket {
	return &schema.NFTBucket{
		ID:                    nft.ID,
		Name:                  nft.Name,
		TeamID:                nft.TeamID,
		Age:                   nft.Age,
		GamePosition:          nft.GamePosition,
		Position:              nft.Position,
		LegendaryClaiming:     nft.LegendaryClaiming,
		LegendaryDefence:      nft.LegendaryDefence,
		LegendaryDistribution: nft.LegendaryDistribution,
		LegendaryDribbling:    nft.LegendaryDribbling,
		LegendaryPassing:      nft.LegendaryPassing,
		LegendaryShooting:     nft.LegendaryShooting,
		LegendaryStopping:     nft.LegendaryStopping,
		LegendaryImage:        nft.LegendaryImage,
		LegendaryMetadata:     nft.LegendaryMetadata,
	}
}

// filterUltraRareRarity filters the NFT to only include ultra rare rarity information.
func filterUltraRareRarity(nft *schema.NFTBucket) *schema.NFTBucket {
	return &schema.NFTBucket{
		ID:                    nft.ID,
		Name:                  nft.Name,
		TeamID:                nft.TeamID,
		Age:                   nft.Age,
		GamePosition:          nft.GamePosition,
		Position:              nft.Position,
		UltraRareClaiming:     nft.UltraRareClaiming,
		UltraRareDefence:      nft.UltraRareDefence,
		UltraRareDistribution: nft.UltraRareDistribution,
		UltraRareDribbling:    nft.UltraRareDribbling,
		UltraRarePassing:      nft.UltraRarePassing,
		UltraRareShooting:     nft.UltraRareShooting,
		UltraRareStopping:     nft.UltraRareStopping,
		UltraRareImage:        nft.UltraRareImage,
		UltraRareMetadata:     nft.UltraRareMetadata,
	}
}

// filterUncommonRarity filters the NFT to only include uncommon rarity information.
func filterUncommonRarity(nft *schema.NFTBucket) *schema.NFTBucket {
	return &schema.NFTBucket{
		ID:                   nft.ID,
		Name:                 nft.Name,
		TeamID:               nft.TeamID,
		Age:                  nft.Age,
		GamePosition:         nft.GamePosition,
		Position:             nft.Position,
		UncommonClaiming:     nft.UncommonClaiming,
		UncommonDefence:      nft.UncommonDefence,
		UncommonDistribution: nft.UncommonDistribution,
		UncommonDribbling:    nft.UncommonDribbling,
		UncommonPassing:      nft.UncommonPassing,
		UncommonShooting:     nft.UncommonShooting,
		UncommonStopping:     nft.UncommonStopping,
		UncommonImage:        nft.UncommonImage,
		UncommonMetadata:     nft.UncommonMetadata,
	}
}

// GetNFTBucketByID retrieves an NFT bucket by its ID and filters the player info based on their NFT rarity.
func (s *DBStore) GetNFTBucketByIDandRarityFiltered(id string, rarity string) (*types.GenericRarityNFT, error) {
	nb, err := schema.FindNFTBucket(s.db, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find NFT bucket by ID %s: %w", id, err)
	}

	return PopulateGenericRarityNFT(nb, rarity), nil
}

// PopulateGenericRarityNFT fills the generic struct based on the rarity.
func PopulateGenericRarityNFT(nft *schema.NFTBucket, rarity string) *types.GenericRarityNFT {
	genericNFT := &types.GenericRarityNFT{
		ID:           nft.ID,
		Name:         nft.Name,
		TeamID:       nft.TeamID,
		Age:          nft.Age,
		GamePosition: nft.GamePosition,
		Position:     nft.Position,
		Nationality:  nft.Nationality.String,
		PlayersGroup: null.StringFrom(nft.PlayersGroup.String),
		OptaID:       nft.OptaID,
	}

	switch rarity {
	case "common":
		genericNFT.Claiming = nft.CommonClaiming
		genericNFT.Defence = nft.CommonDefence
		genericNFT.Distribution = nft.CommonDistribution
		genericNFT.Dribbling = nft.CommonDribbling
		genericNFT.Passing = nft.CommonPassing
		genericNFT.Shooting = nft.CommonShooting
		genericNFT.Stopping = nft.CommonStopping
		genericNFT.Image = nft.CommonImage
		genericNFT.Metadata = nft.CommonMetadata
		genericNFT.Limit = nft.CommonLimit.Int
	case "rare":
		genericNFT.Claiming = nft.RareClaiming
		genericNFT.Defence = nft.RareDefence
		genericNFT.Distribution = nft.RareDistribution
		genericNFT.Dribbling = nft.RareDribbling
		genericNFT.Passing = nft.RarePassing
		genericNFT.Shooting = nft.RareShooting
		genericNFT.Stopping = nft.RareStopping
		genericNFT.Image = nft.RareImage
		genericNFT.Metadata = nft.RareMetadata
		genericNFT.Limit = nft.RareLimit.Int
	case "legendary":
		genericNFT.Claiming = nft.LegendaryClaiming
		genericNFT.Defence = nft.LegendaryDefence
		genericNFT.Distribution = nft.LegendaryDistribution
		genericNFT.Dribbling = nft.LegendaryDribbling
		genericNFT.Passing = nft.LegendaryPassing
		genericNFT.Shooting = nft.LegendaryShooting
		genericNFT.Stopping = nft.LegendaryStopping
		genericNFT.Image = nft.LegendaryImage
		genericNFT.Metadata = nft.LegendaryMetadata
		genericNFT.Limit = nft.LegendaryLimit.Int
	case "ultra_rare":
		genericNFT.Claiming = nft.UltraRareClaiming
		genericNFT.Defence = nft.UltraRareDefence
		genericNFT.Distribution = nft.UltraRareDistribution
		genericNFT.Dribbling = nft.UltraRareDribbling
		genericNFT.Passing = nft.UltraRarePassing
		genericNFT.Shooting = nft.UltraRareShooting
		genericNFT.Stopping = nft.UltraRareStopping
		genericNFT.Image = nft.UltraRareImage
		genericNFT.Metadata = nft.UltraRareMetadata
		genericNFT.Limit = nft.UltraRareLimit.Int
	case "uncommon":
		genericNFT.Claiming = nft.UncommonClaiming
		genericNFT.Defence = nft.UncommonDefence
		genericNFT.Distribution = nft.UncommonDistribution
		genericNFT.Dribbling = nft.UncommonDribbling
		genericNFT.Passing = nft.UncommonPassing
		genericNFT.Shooting = nft.UncommonShooting
		genericNFT.Stopping = nft.UncommonStopping
		genericNFT.Image = nft.UncommonImage
		genericNFT.Metadata = nft.UncommonMetadata
		genericNFT.Limit = nft.UncommonLimit.Int
	case "Common":
		genericNFT.Claiming = nft.CommonClaiming
		genericNFT.Defence = nft.CommonDefence
		genericNFT.Distribution = nft.CommonDistribution
		genericNFT.Dribbling = nft.CommonDribbling
		genericNFT.Passing = nft.CommonPassing
		genericNFT.Shooting = nft.CommonShooting
		genericNFT.Stopping = nft.CommonStopping
		genericNFT.Image = nft.CommonImage
		genericNFT.Metadata = nft.CommonMetadata
		genericNFT.Limit = nft.CommonLimit.Int
	case "Rare":
		genericNFT.Claiming = nft.RareClaiming
		genericNFT.Defence = nft.RareDefence
		genericNFT.Distribution = nft.RareDistribution
		genericNFT.Dribbling = nft.RareDribbling
		genericNFT.Passing = nft.RarePassing
		genericNFT.Shooting = nft.RareShooting
		genericNFT.Stopping = nft.RareStopping
		genericNFT.Image = nft.RareImage
		genericNFT.Metadata = nft.RareMetadata
		genericNFT.Limit = nft.RareLimit.Int
	case "Legendary":
		genericNFT.Claiming = nft.LegendaryClaiming
		genericNFT.Defence = nft.LegendaryDefence
		genericNFT.Distribution = nft.LegendaryDistribution
		genericNFT.Dribbling = nft.LegendaryDribbling
		genericNFT.Passing = nft.LegendaryPassing
		genericNFT.Shooting = nft.LegendaryShooting
		genericNFT.Stopping = nft.LegendaryStopping
		genericNFT.Image = nft.LegendaryImage
		genericNFT.Metadata = nft.LegendaryMetadata
		genericNFT.Limit = nft.LegendaryLimit.Int
	case "UltraRare":
		genericNFT.Claiming = nft.UltraRareClaiming
		genericNFT.Defence = nft.UltraRareDefence
		genericNFT.Distribution = nft.UltraRareDistribution
		genericNFT.Dribbling = nft.UltraRareDribbling
		genericNFT.Passing = nft.UltraRarePassing
		genericNFT.Shooting = nft.UltraRareShooting
		genericNFT.Stopping = nft.UltraRareStopping
		genericNFT.Image = nft.UltraRareImage
		genericNFT.Metadata = nft.UltraRareMetadata
		genericNFT.Limit = nft.UltraRareLimit.Int
	case "Uncommon":
		genericNFT.Claiming = nft.UncommonClaiming
		genericNFT.Defence = nft.UncommonDefence
		genericNFT.Distribution = nft.UncommonDistribution
		genericNFT.Dribbling = nft.UncommonDribbling
		genericNFT.Passing = nft.UncommonPassing
		genericNFT.Shooting = nft.UncommonShooting
		genericNFT.Stopping = nft.UncommonStopping
		genericNFT.Image = nft.UncommonImage
		genericNFT.Metadata = nft.UncommonMetadata
		genericNFT.Limit = nft.UncommonLimit.Int
	}

	return genericNFT
}

// GetAllNFTBuckets retrieves all NFT buckets from the database with detailed error handling.
func (s *DBStore) GetAllNFTBuckets() (schema.NFTBucketSlice, error) {
	nbs, err := schema.NFTBuckets().All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all NFT buckets: %w", err)
	}
	return nbs, nil
}

// UpdateNFTBucket updates an existing NFT bucket with detailed error handling.
func (s *DBStore) UpdateNFTBucket(nb *schema.NFTBucket) error {
	_, err := nb.Update(s.db, boil.Whitelist("name", "team_id", "age", "game_position", "position", "star_rating", "updated_at"))
	if err != nil {
		return fmt.Errorf("failed to update NFT bucket: %w", err)
	}
	return nil
}

// DeleteNFTBucket deletes an NFT bucket by its ID with detailed error handling.
func (s *DBStore) DeleteNFTBucket(id string) error {
	nb, err := schema.FindNFTBucket(s.db, id)
	if err != nil {
		return fmt.Errorf("failed to find NFT bucket for deletion: %w", err)
	}
	_, err = nb.Delete(s.db)
	if err != nil {
		return fmt.Errorf("failed to delete NFT bucket: %w", err)
	}
	return nil
}

// GetNFTBucketsByTeamID retrieves all NFT buckets for a given team ID with detailed error handling.
func (s *DBStore) GetNFTBucketsByTeamID(teamID string) (schema.NFTBucketSlice, error) {
	nbs, err := schema.NFTBuckets(
		qm.Where("team_id = ?", teamID),
	).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFT buckets by team ID %s: %w", teamID, err)
	}
	return nbs, nil
}

// CountNFTBuckets returns the total number of NFT buckets with detailed error handling.
func (s *DBStore) CountNFTBuckets() (int64, error) {
	count, err := schema.NFTBuckets().Count(s.db)
	if err != nil {
		return 0, fmt.Errorf("failed to count NFT buckets: %w", err)
	}
	return count, nil
}

// GetNFTsByRarity retrieves all NFTs with a specific rarity.
func (s *DBStore) GetNFTsByRarity(rarity string) (schema.NFTBucketSlice, error) {
	var mods []qm.QueryMod

	switch rarity {
	case "common":
		mods = append(mods, qm.Where("common_claiming IS NOT NULL"))
	case "uncommon":
		mods = append(mods, qm.Where("uncommon_claiming IS NOT NULL"))
	case "rare":
		mods = append(mods, qm.Where("rare_claiming IS NOT NULL"))
	case "ultra_rare":
		mods = append(mods, qm.Where("ultra_rare_claiming IS NOT NULL"))
	case "legendary":
		mods = append(mods, qm.Where("legendary_claiming IS NOT NULL"))
	default:
		return nil, fmt.Errorf("invalid rarity: %s", rarity)
	}

	nbs, err := schema.NFTBuckets(mods...).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFT buckets by %s rarity: %w", rarity, err)
	}
	return nbs, nil
}

// GetNFTsByTeamAndRarity retrieves all NFTs from a specific team with a specific rarity.
func (s *DBStore) GetNFTsByTeamAndRarity(teamID, rarity string) (schema.NFTBucketSlice, error) {
	var mods []qm.QueryMod
	mods = append(mods, qm.Where("team_id = ?", teamID))
	switch rarity {
	case "common":
		mods = append(mods, qm.Where("common_claiming IS NOT NULL"))
	case "uncommon":
		mods = append(mods, qm.Where("uncommon_claiming IS NOT NULL"))
	case "rare":
		mods = append(mods, qm.Where("rare_claiming IS NOT NULL"))
	case "ultra_rare":
		mods = append(mods, qm.Where("ultra_rare_claiming IS NOT NULL"))
	case "legendary":
		mods = append(mods, qm.Where("legendary_claiming IS NOT NULL"))
	default:
		return nil, fmt.Errorf("invalid rarity: %s", rarity)
	}

	nbs, err := schema.NFTBuckets(mods...).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFT buckets by team ID %s and %s rarity: %w", teamID, rarity, err)
	}
	return nbs, nil
}

// GetRandomNFTByTeamAndRarity retrieves a random NFT from a specific team with a specified rarity.
func (s *DBStore) GetRandomNFTByTeamAndRarity(teamID string, rarity string) (*schema.NFTBucket, error) {
	var mods []qm.QueryMod
	mods = append(mods, qm.Where("team_id = ?", teamID))

	switch rarity {
	case "common":
		mods = append(mods, qm.Where("common_claiming IS NOT NULL"))
	case "uncommon":
		mods = append(mods, qm.Where("uncommon_claiming IS NOT NULL"))
	case "rare":
		mods = append(mods, qm.Where("rare_claiming IS NOT NULL"))
	case "ultra_rare":
		mods = append(mods, qm.Where("ultra_rare_claiming IS NOT NULL"))
	case "legendary":
		mods = append(mods, qm.Where("legendary_claiming IS NOT NULL"))
	default:
		return nil, fmt.Errorf("invalid rarity: %s", rarity)
	}

	// Retrieve all matching NFTs
	nfts, err := schema.NFTBuckets(mods...).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFTs by team ID %s and %s rarity: %w", teamID, rarity, err)
	}

	if len(nfts) == 0 {
		return nil, fmt.Errorf("no NFTs found for team ID %s with %s rarity", teamID, rarity)
	}

	// Select a random NFT from the list
	randomIndex := rand.Intn(len(nfts))
	return nfts[randomIndex], nil
}

// GetRandomNFTBucketByTeamID retrieves a random NFT bucket for a given team ID.
func (s *DBStore) GetRandomNFTBucketByTeamID(teamID string) (*schema.NFTBucket, error) {
	// Retrieve all NFT buckets for the specified team ID
	nfts, err := schema.NFTBuckets(
		qm.Where("team_id = ?", teamID),
	).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFT buckets by team ID %s: %w", teamID, err)
	}

	if len(nfts) == 0 {
		return nil, fmt.Errorf("no NFTs found for team ID %s", teamID)
	}

	// Select a random NFT from the list
	randomIndex := rand.Intn(len(nfts))
	return nfts[randomIndex], nil
}

// GetNFTBucketsByPosition retrieves all NFT buckets with a specific position.
func (s *DBStore) GetNFTBucketsByPosition(position string) (schema.NFTBucketSlice, error) {
	nbs, err := schema.NFTBuckets(
		qm.Where("position = ?", position),
	).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFT buckets by position %s: %w", position, err)
	}
	return nbs, nil
}

// GetNFTBucketsByGamePosition retrieves all NFT buckets with a specific game position.
func (s *DBStore) GetNFTBucketsByGamePosition(gamePosition string) (schema.NFTBucketSlice, error) {
	nbs, err := schema.NFTBuckets(
		qm.Where("game_position = ?", gamePosition),
	).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFT buckets by game position %s: %w", gamePosition, err)
	}
	return nbs, nil
}

// GetNFTBucketsByPositionAndTeam retrieves all NFT buckets for a specific team and position.
func (s *DBStore) GetNFTBucketsByPositionAndTeam(teamID, position string) (schema.NFTBucketSlice, error) {
	nbs, err := schema.NFTBuckets(
		qm.Where("team_id = ? AND position = ?", teamID, position),
	).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFT buckets by team ID %s and position %s: %w", teamID, position, err)
	}
	return nbs, nil
}

// GetNFTBucketsByTeamAndNationality retrieves all NFT buckets for a specific team and nationality.
func (s *DBStore) GetNFTBucketsByTeamAndNationality(teamID, nationality string) (schema.NFTBucketSlice, error) {
	nbs, err := schema.NFTBuckets(
		qm.Where("team_id = ? AND nationality = ?", teamID, nationality),
	).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFT buckets by team ID %s and nationality %s: %w", teamID, nationality, err)
	}
	return nbs, nil
}

// GetNFTBucketsByNationality retrieves all NFT buckets with a specific nationality.
func (s *DBStore) GetNFTBucketsByNationality(nationality string) (schema.NFTBucketSlice, error) {
	nbs, err := schema.NFTBuckets(
		qm.Where("nationality = ?", nationality),
	).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFT buckets by nationality %s: %w", nationality, err)
	}
	return nbs, nil
}

// GetNFTBucketsByAgeRange retrieves all NFT buckets within a specific age range.
func (s *DBStore) GetNFTBucketsByAgeRange(minAge, maxAge int) (schema.NFTBucketSlice, error) {
	nbs, err := schema.NFTBuckets(
		qm.Where("age BETWEEN ? AND ?", minAge, maxAge),
	).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFT buckets by age range %d-%d: %w", minAge, maxAge, err)
	}
	return nbs, nil
}

// GetNFTBucketsByRarityAndPosition retrieves all NFT buckets with a specific rarity and position.
func (s *DBStore) GetNFTBucketsByRarityAndPosition(rarity, position string) (schema.NFTBucketSlice, error) {
	var mods []qm.QueryMod

	switch rarity {
	case "common":
		mods = append(mods, qm.Where("common_claiming IS NOT NULL"))
	case "uncommon":
		mods = append(mods, qm.Where("uncommon_claiming IS NOT NULL"))
	case "rare":
		mods = append(mods, qm.Where("rare_claiming IS NOT NULL"))
	case "ultra_rare":
		mods = append(mods, qm.Where("ultra_rare_claiming IS NOT NULL"))
	case "legendary":
		mods = append(mods, qm.Where("legendary_claiming IS NOT NULL"))
	default:
		return nil, fmt.Errorf("invalid rarity: %s", rarity)
	}

	mods = append(mods, qm.Where("position = ?", position))

	nbs, err := schema.NFTBuckets(mods...).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFT buckets by rarity %s and position %s: %w", rarity, position, err)
	}
	return nbs, nil
}

// CountNFTBucketsByRarity returns the total number of NFT buckets for each rarity.
func (s *DBStore) CountNFTBucketsByRarity() (map[string]int64, error) {
	counts := make(map[string]int64)

	var err error
	counts["common"], err = schema.NFTBuckets(qm.Where("common_claiming IS NOT NULL")).Count(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to count common NFT buckets: %w", err)
	}

	counts["uncommon"], err = schema.NFTBuckets(qm.Where("uncommon_claiming IS NOT NULL")).Count(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to count uncommon NFT buckets: %w", err)
	}

	counts["rare"], err = schema.NFTBuckets(qm.Where("rare_claiming IS NOT NULL")).Count(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to count rare NFT buckets: %w", err)
	}

	counts["ultra_rare"], err = schema.NFTBuckets(qm.Where("ultra_rare_claiming IS NOT NULL")).Count(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to count ultra rare NFT buckets: %w", err)
	}

	counts["legendary"], err = schema.NFTBuckets(qm.Where("legendary_claiming IS NOT NULL")).Count(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to count legendary NFT buckets: %w", err)
	}

	return counts, nil
}

// AssignNFTToTeam updates or creates a new NFT bucket with a specific team assignment with detailed error handling.
func (s *DBStore) AssignNFTToTeam(nftID, teamID string) (*schema.NFTBucket, error) {
	nb := &schema.NFTBucket{
		ID:     nftID,
		TeamID: null.StringFrom(teamID),
	}
	if err := nb.Upsert(s.db, true, []string{"id"}, boil.Whitelist("team_id"), boil.Infer()); err != nil {
		return nil, fmt.Errorf("failed to upsert NFT bucket: %w", err)
	}
	return nb, nil
}

// GetRandomNFTBucket retrieves a random NFT bucket.
func (s *DBStore) GetRandomNFTBucket() (*schema.NFTBucket, error) {
	nbs, err := schema.NFTBuckets().All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all NFT buckets: %w", err)
	}

	if len(nbs) == 0 {
		return nil, fmt.Errorf("no NFT buckets found")
	}

	randomIndex := rand.Intn(len(nbs))
	return nbs[randomIndex], nil
}

// GetNFTBucketByName retrieves an NFT bucket by its name.
func (s *DBStore) GetNFTBucketByName(name string) (*schema.NFTBucket, error) {
	nb, err := schema.NFTBuckets(qm.Where("name = ?", name)).One(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to find NFT bucket by name %s: %w", name, err)
	}
	return nb, nil
}

func (s *DBStore) GetRandomNFTBucketByTeamAndRarityAndStarRating(teamID string, rarity string, starRating int, cardpackCollection string) (*schema.NFTBucket, error) {
	var mods []qm.QueryMod
	mods = append(mods, qm.Where("team_id = ?", teamID), qm.Where("star_rating = ?", starRating), qm.Where("players_group = ?", cardpackCollection))

	switch rarity {
	case "common":
		mods = append(mods, qm.Where("common_limit > 0"))
	case "uncommon":
		mods = append(mods, qm.Where("uncommon_limit > 0"))
	case "rare":
		mods = append(mods, qm.Where("rare_limit > 0"))
	case "ultra_rare":
		mods = append(mods, qm.Where("ultra_rare_limit > 0"))
	case "legendary":
		mods = append(mods, qm.Where("legendary_limit > 0"))
	default:
		return nil, fmt.Errorf("invalid rarity: %s", rarity)
	}

	nfts, err := schema.NFTBuckets(mods...).All(s.db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve NFTs by team ID %s and %s rarity: %w", teamID, rarity, err)
	}

	if len(nfts) == 0 {
		return nil, fmt.Errorf("no NFTs found for team ID %s with rarity %s and star rating %d", teamID, rarity, starRating)
	}

	rand.Shuffle(len(nfts), func(i, j int) { nfts[i], nfts[j] = nfts[j], nfts[i] })
	return nfts[0], nil
}

func (s *DBStore) DeductFromRarityLimit(nftID string, rarity string, amount int) error {
	nb, err := schema.FindNFTBucket(s.db, nftID)
	if err != nil {
		return fmt.Errorf("failed to find NFT bucket by ID %s: %w", nftID, err)
	}

	// Deduct from the appropriate rarity limit
	switch rarity {
	case "common":
		if !nb.CommonLimit.Valid {
			return fmt.Errorf("common limit not set for NFT bucket ID %s", nftID)
		}
		if nb.CommonLimit.Int < amount {
			return fmt.Errorf("insufficient common limit to deduct: required %d, available %d", amount, nb.CommonLimit.Int)
		}
		nb.CommonLimit.Int -= amount
		log.Printf("Deducted %d from common limit, remaining: %d", amount, nb.CommonLimit.Int)

	case "legendary":
		if !nb.LegendaryLimit.Valid {
			return fmt.Errorf("legendary limit not set for NFT bucket ID %s", nftID)
		}
		if nb.LegendaryLimit.Int < amount {
			return fmt.Errorf("insufficient legendary limit to deduct: required %d, available %d", amount, nb.LegendaryLimit.Int)
		}
		nb.LegendaryLimit.Int -= amount
		log.Printf("Deducted %d from legendary limit, remaining: %d", amount, nb.LegendaryLimit.Int)

	case "rare":
		if !nb.RareLimit.Valid {
			return fmt.Errorf("rare limit not set for NFT bucket ID %s", nftID)
		}
		if nb.RareLimit.Int < amount {
			return fmt.Errorf("insufficient rare limit to deduct: required %d, available %d", amount, nb.RareLimit.Int)
		}
		nb.RareLimit.Int -= amount
		log.Printf("Deducted %d from rare limit, remaining: %d", amount, nb.RareLimit.Int)

	case "ultra_rare":
		if !nb.UltraRareLimit.Valid {
			return fmt.Errorf("ultra rare limit not set for NFT bucket ID %s", nftID)
		}
		if nb.UltraRareLimit.Int < amount {
			return fmt.Errorf("insufficient ultra rare limit to deduct: required %d, available %d", amount, nb.UltraRareLimit.Int)
		}
		nb.UltraRareLimit.Int -= amount
		log.Printf("Deducted %d from ultra rare limit, remaining: %d", amount, nb.UltraRareLimit.Int)

	case "uncommon":
		if !nb.UncommonLimit.Valid {
			return fmt.Errorf("uncommon limit not set for NFT bucket ID %s", nftID)
		}
		if nb.UncommonLimit.Int < amount {
			return fmt.Errorf("insufficient uncommon limit to deduct: required %d, available %d", amount, nb.UncommonLimit.Int)
		}
		nb.UncommonLimit.Int -= amount
		log.Printf("Deducted %d from uncommon limit, remaining: %d", amount, nb.UncommonLimit.Int)

	default:
		return fmt.Errorf("invalid rarity: %s", rarity)
	}

	// Update the NFT bucket in the database
	columnToUpdate := fmt.Sprintf("%s_limit", strings.ToLower(rarity))
	if _, err := nb.Update(s.db, boil.Whitelist(columnToUpdate)); err != nil {
		return fmt.Errorf("failed to update %s limit for NFT bucket ID %s: %w", rarity, nftID, err)
	}
	log.Printf("Successfully updated the %s limit for NFT bucket ID %s", rarity, nftID)

	return nil
}

// CheckLimit checks if the NFT bucket's limit for a specific rarity is greater than zero.
func (s *DBStore) CheckLimit(nftID string, rarity string) (bool, error) {
	// Find the NFT bucket by its ID
	nb, err := schema.FindNFTBucket(s.db, nftID)
	if err != nil {
		return false, fmt.Errorf("failed to find NFT bucket by ID %s: %w", nftID, err)
	}

	// Check the limit based on rarity
	switch strings.ToLower(rarity) {
	case "common":
		if !nb.CommonLimit.Valid {
			return false, fmt.Errorf("common limit not set for NFT bucket ID %s", nftID)
		}
		return nb.CommonLimit.Int > 0, nil
	case "uncommon":
		if !nb.UncommonLimit.Valid {
			return false, fmt.Errorf("uncommon limit not set for NFT bucket ID %s", nftID)
		}
		return nb.UncommonLimit.Int > 0, nil
	case "rare":
		if !nb.RareLimit.Valid {
			return false, fmt.Errorf("rare limit not set for NFT bucket ID %s", nftID)
		}
		return nb.RareLimit.Int > 0, nil
	case "ultra_rare":
		if !nb.UltraRareLimit.Valid {
			return false, fmt.Errorf("ultra rare limit not set for NFT bucket ID %s", nftID)
		}
		return nb.UltraRareLimit.Int > 0, nil
	case "legendary":
		if !nb.LegendaryLimit.Valid {
			return false, fmt.Errorf("legendary limit not set for NFT bucket ID %s", nftID)
		}
		return nb.LegendaryLimit.Int > 0, nil
	default:
		return false, fmt.Errorf("invalid rarity: %s", rarity)
	}
}
