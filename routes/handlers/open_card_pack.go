package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types"
	"golang.org/x/exp/rand"
)

type NFTS []struct {
	TierData schema.NFTBucket
	Rarity   string
}

// OpenCardPack godoc
// @Summary Open a card pack
// @Description Open a card pack and assign players to the user
// @Tags cardPacks
// @Security BearerAuth
// @ID open-card-pack
// @Param cardPackID path string true "Card pack ID"
// @Success 200 {object} map[string]string "Card pack opened successfully"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Card pack not found"
// @Failure 500 {object} map[string]string "Failed to open card pack"
// @Router /card-packs/{cardPackID}/open [post]
func (e *Env) OpenCardPack(c echo.Context) error {
	userID := userID(c)
	cardPackID := c.Param("cardPackID")

	user, err := e.Store.GetUserByID(userID)
	if user == nil {
		return handleError(c, 10000, "User not found. Please log in and try again.", http.StatusNotFound)
	}

	if err != nil {
		return handleError(c, 10000, "Failed to retrieve user information. Please try again later.", http.StatusInternalServerError)
	}

	if user.WalletAddress.IsZero() {
		return handleError(c, 10000, "Sign up to open the card pack.", http.StatusNotFound)
	}

	acp, err := e.Store.GetAssignedCardPackByID(cardPackID)
	if err != nil {
		return handleError(c, 10001, "The requested card pack was not found. Please check the card pack ID and try again.", http.StatusNotFound)
	}
	if acp.UserID != userID || acp.Opened {
		return handleError(c, 10002, "Unauthorized access or the card pack has already been opened. Please verify your access rights or card pack status.", http.StatusUnauthorized)
	}

	cardPackType, err := e.Store.GetCardPackByID(acp.CardPackTypeID)
	if err != nil {
		return handleError(c, 10003, "Failed to retrieve card pack type information. Please try again later or contact support if the issue persists.", http.StatusInternalServerError)
	}

	var nfts NFTS
	usedTeams := make(map[string]struct{})
	totalPlayers := 7

	tierRarities := []struct {
		TierData types.JSON
		Rarity   string
	}{
		{cardPackType.Tier1, "common"},
		{cardPackType.Tier2, "uncommon"},
		{cardPackType.Tier3, "rare"},
		{cardPackType.Tier4, "ultra_rare"},
		{cardPackType.Tier5, "legendary"},
	}

	// Process each tier and assign players
	for _, tier := range tierRarities {
		var tiers [][]int
		if err := json.Unmarshal(tier.TierData, &tiers); err != nil {
			return handleError(c, 10004, "Failed to parse tier data: "+err.Error(), http.StatusInternalServerError)
		}

		for _, tierEntry := range tiers {
			amount, starRating := tierEntry[0], tierEntry[1]
			for i := 0; i < amount; i++ {
				player, err := getPlayerByRarityAndStarRating(e, tier.Rarity, starRating, usedTeams, cardPackType.Collection.String)
				if err != nil {
					continue
				}
				nfts = append(nfts, struct {
					TierData schema.NFTBucket
					Rarity   string
				}{
					TierData: *player,
					Rarity:   tier.Rarity,
				})
			}
		}
	}

	// Fill with random players if needed
	for len(nfts) < totalPlayers {
		player, rarity, err := getPlayerByRarityAndStarRatingRandom(e, usedTeams, cardPackType.Collection.String)
		if err != nil {
			return handleError(c, 10005, "Failed to assign additional random players: "+err.Error(), http.StatusInternalServerError)
		}
		nfts = append(nfts, struct {
			TierData schema.NFTBucket
			Rarity   string
		}{
			TierData: *player,
			Rarity:   rarity,
		})
	}

	assignedPlayers := make([]*schema.AssignedPlayer, len(nfts))
	nftCount, err := e.Store.CountAllEntries()
	if err != nil {
		return handleError(c, 10014, "Failed to retrieve NFT count. Please try again later.", http.StatusInternalServerError)
	}

	// Convert nftCount to int and create the increments slice
	nftCountInt := int(nftCount)
	increments := make([]*big.Int, 7)
	for i := 0; i < 7; i++ {
		increments[i] = big.NewInt(int64(nftCountInt + i + 1))
	}

	for i, nft := range nfts {
		assignedPlayer := &schema.AssignedPlayer{
			ID:          uuid.New().String(),
			PlayerNFTID: null.StringFrom(nft.TierData.ID),
			NFTID:       null.StringFrom(increments[i].String()),
			Rarity:      null.StringFrom(nft.Rarity),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err := e.Store.DeductFromRarityLimit(nft.TierData.ID, nft.Rarity, 1)
		if err != nil {
			return handleError(c, 10010, "Failed to deduct from rarity limit. Please try again later.", http.StatusInternalServerError)
		}

		if _, err := e.Store.CreateAssignedPlayer(assignedPlayer); err != nil {
			return handleError(c, 10009, "Failed to assign player to user: "+err.Error(), http.StatusInternalServerError)
		}
		assignedPlayers[i] = assignedPlayer
	}

	assignedPlayersJSONStrings := make([]string, len(assignedPlayers))
	for i, player := range assignedPlayers {
		playerJSON, err := json.Marshal(player)
		if err != nil {
			return handleError(c, 10013, "Failed to serialize player information. Please try again later.", http.StatusInternalServerError)
		}
		assignedPlayersJSONStrings[i] = string(playerJSON)
	}

	if err := e.Store.SetAssignedCardPackOpened(cardPackID); err != nil {
		return handleError(c, 10015, "Failed to mark the card pack as opened. Please contact support for further assistance.", http.StatusInternalServerError)
	}

	// Collect all UUIDs from assigned players
	playerUUIDs := make([]string, len(assignedPlayers))
	for i, player := range assignedPlayers {
		playerUUIDs[i] = player.ID
	}
	playersWithNFTs, err := e.Store.GetAssignedPlayersWithNFTDetails(playerUUIDs)
	if err != nil {
		return handleError(c, 10016, "Failed to retrieve players with NFT details. Please contact support for further assistance.", http.StatusInternalServerError)
	}
	if err := e.Store.UpdateAssignedCardPackCardIds(cardPackID, playerUUIDs); err != nil {
		return handleError(c, 10016, "Failed to update card IDs of the assigned card pack. Please contact support for further assistance.", http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Card pack opened successfully.",
		"players": playersWithNFTs,
	})
}

// handleError is a utility function to simplify error responses
func handleError(c echo.Context, code int, message string, httpStatus int) error {
	return c.JSON(httpStatus, echo.Map{
		"code":    code,
		"message": message,
	})
}

// getPlayerByRarityAndStarRating retrieves a player by rarity and star rating, ensuring no team is reused
func getPlayerByRarityAndStarRating(e *Env, rarity string, starRating int, usedTeams map[string]struct{}, cardpackCollection string) (*schema.NFTBucket, error) {
	var player *schema.NFTBucket
	retryLimit := 30

	for attempts := 0; attempts < retryLimit; attempts++ {
		// Get a unique random team that hasn't been used yet
		team, err := e.Store.GetUniqueRandomTeamExcluding(usedTeams)
		if err != nil {
			fmt.Println("Error retrieving team:", err)
			continue // Skip to the next iteration if no unique team found
		}

		// Try to get a random NFT bucket with the specified team, rarity, and star rating
		player, err = e.Store.GetRandomNFTBucketByTeamAndRarityAndStarRating(team.ID, rarity, starRating, cardpackCollection)
		if err != nil {
			fmt.Println("Error retrieving player:", err)
			continue // Retry with a new team if there was an error
		}

		// Check if the player has a valid limit
		isLimited, err := e.Store.CheckLimit(player.ID, rarity)
		if err != nil {
			fmt.Println("Error checking limit:", err)
			continue // Retry if there was an error checking the limit
		}

		if isLimited {
			usedTeams[team.ID] = struct{}{}
			return player, nil // Return the player if the limit is still valid
		}

		// If the limit is not valid, continue the loop to try again
	}

	// Return an error if the function fails to find a valid player after the retry limit
	return nil, fmt.Errorf("failed to assign a player of the specified rarity [%s] and star rating [%d] after %d attempts", rarity, starRating, retryLimit)
}

// getPlayerByRarityAndStarRatingRandom attempts to retrieve a random player by iterating over different rarities and star ratings.
func getPlayerByRarityAndStarRatingRandom(e *Env, usedTeams map[string]struct{}, cardpackCollection string) (*schema.NFTBucket, string, error) {
	var player *schema.NFTBucket
	retryLimit := 20

	// Rarity probabilities
	rarityProbabilities := map[string]float64{
		"common":     0.375,
		"uncommon":   0.325,
		"rare":       0.2,
		"ultra_rare": 0.08,
		"legendary":  0.02,
	}

	// Star rating probabilities
	starRatingProbabilities := map[int]float64{
		1: 0.2,
		2: 0.25,
		3: 0.35,
		4: 0.15,
		5: 0.05,
	}

	// Seed the random number generator
	rand.Seed(uint64(time.Now().UnixNano()))

	// Function to select a rarity based on weighted probabilities
	selectRarity := func() string {
		r := rand.Float64()
		accumulated := 0.0
		for rarity, probability := range rarityProbabilities {
			accumulated += probability
			if r < accumulated {
				return rarity
			}
		}
		return "common" // Default to common if no other rarity is selected
	}

	// Function to select a star rating based on weighted probabilities
	selectStarRating := func() int {
		r := rand.Float64()
		accumulated := 0.0
		for starRating, probability := range starRatingProbabilities {
			accumulated += probability
			if r < accumulated {
				return starRating
			}
		}
		return 1 // Default to 1-star if no other star rating is selected
	}

	for attempts := 0; attempts < retryLimit; attempts++ {
		rarity := selectRarity()
		starRating := selectStarRating()

		team, err := e.Store.GetUniqueRandomTeamExcluding(usedTeams)
		if err != nil {
			fmt.Println("Error retrieving team:", err)
			continue
		}

		player, err = e.Store.GetRandomNFTBucketByTeamAndRarityAndStarRating(team.ID, rarity, starRating, cardpackCollection)
		if err != nil {
			fmt.Println("Error retrieving player:", err)
			continue // Retry with a new team if there was an error
		}

		// Check if the player has a valid limit
		isLimited, err := e.Store.CheckLimit(player.ID, rarity)
		if err != nil {
			fmt.Println("Error checking limit:", err)
			continue // Retry if there was an error checking the limit
		}

		if isLimited {
			usedTeams[team.ID] = struct{}{} // Mark the team as used
			return player, rarity, nil      // Return the player if the limit is still valid
		}

		// If the limit is not valid, continue the loop to try again
	}

	// Return an error if the function fails to find a valid player after the retry limit
	return nil, "", fmt.Errorf("failed to assign a player after %d attempts", retryLimit)
}

// OpenCardPack godoc
// @Summary Open a card pack
// @Description Open a card pack and assign players to the user
// @Tags cardPacks
// @Security BearerAuth
// @ID open-card-pack
// @Param cardPackID path string true "Card pack ID"
// @Success 200 {object} map[string]string "Card pack opened successfully"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Card pack not found"
// @Failure 500 {object} map[string]string "Failed to open card pack"
// @Router /card-packs/{cardPackID}/open [post]
func (e *Env) NewOpenCardPack(c echo.Context) error {
	userID := userID(c)
	cardPackID := c.Param("cardPackID")

	user, err := e.Store.GetUserByID(userID)
	if user == nil {
		return handleError(c, 10000, "User not found. Please log in and try again.", http.StatusNotFound)
	}

	if err != nil {
		return handleError(c, 10000, "Failed to retrieve user information. Please try again later.", http.StatusInternalServerError)
	}

	if user.WalletAddress.IsZero() {
		return handleError(c, 10000, "Sign up to open the card pack", http.StatusNotFound)
	}

	acp, err := e.Store.GetAssignedCardPackByID(cardPackID)
	if err != nil {
		return handleError(c, 10001, "The requested card pack was not found. Please check the card pack ID and try again.", http.StatusNotFound)
	}
	if acp.UserID != userID || acp.Opened {
		return handleError(c, 10002, "Unauthorized access or the card pack has already been opened. Please verify your access rights or card pack status.", http.StatusUnauthorized)
	}

	cardPackType, err := e.Store.GetCardPackByID(acp.CardPackTypeID)
	if err != nil {
		return handleError(c, 10003, "Failed to retrieve card pack type information. Please try again later or contact support if the issue persists.", http.StatusInternalServerError)
	}

	var nfts NFTS
	usedTeams := make(map[string]struct{})
	totalPlayers := 7

	// Add iteration for Rarities and Star Ratings
	var starRatings [][]int
	var rarities [][]interface{}

	if cardPackType.Rarities.Valid {
		log.Printf("Raw Rarities data: %s", cardPackType.Rarities.JSON)
		if err := json.Unmarshal(cardPackType.Rarities.JSON, &rarities); err != nil {
			return handleError(c, 10004, "Failed to parse rarities data: "+err.Error(), http.StatusInternalServerError)
		}
	}

	// Convert the parsed array into the desired struct format
	var parsedRarities []struct {
		Rarity string
		Amount int
	}

	for _, entry := range rarities {
		if len(entry) != 2 {
			return handleError(c, 10004, "Invalid rarity format.", http.StatusInternalServerError)
		}

		rarity, ok1 := entry[0].(string)
		amount, ok2 := entry[1].(float64) // JSON numbers are parsed as float64

		if !ok1 || !ok2 {
			return handleError(c, 10004, "Failed to parse rarity data.", http.StatusInternalServerError)
		}

		parsedRarities = append(parsedRarities, struct {
			Rarity string
			Amount int
		}{
			Rarity: rarity,
			Amount: int(amount),
		})
	}

	// Process parsed rarities
	for _, rarityEntry := range parsedRarities {
		for i := 0; i < rarityEntry.Amount; i++ {
			player, err := getPlayerByRarityAndRandomStarRating(e, rarityEntry.Rarity, usedTeams, cardPackType.Collection.String)
			if err != nil {
				continue
			}
			nfts = append(nfts, struct {
				TierData schema.NFTBucket
				Rarity   string
			}{
				TierData: *player,
				Rarity:   rarityEntry.Rarity,
			})
		}
	}
	if cardPackType.StarRatings.Valid {
		log.Printf("Raw StarRatings data: %s", cardPackType.StarRatings.JSON)
		if err := json.Unmarshal(cardPackType.StarRatings.JSON, &starRatings); err != nil {
			return handleError(c, 10004, "Failed to parse star ratings data: "+err.Error(), http.StatusInternalServerError)
		}
	}

	// Assign based on specific star ratings if no rarities are specified
	for _, starEntry := range starRatings {
		starRating, amount := starEntry[0], starEntry[1]
		for i := 0; i < amount; i++ {
			// Assign player with specific star rating, random rarity
			player, rarity, err := getPlayerByStarRatingAndRandomRarity(e, starRating, usedTeams, cardPackType.Collection.String)
			if err != nil {
				continue
			}
			nfts = append(nfts, struct {
				TierData schema.NFTBucket
				Rarity   string
			}{
				TierData: *player,
				Rarity:   rarity,
			})
		}
	}

	// Fill with random players if needed
	for len(nfts) < totalPlayers {
		player, rarity, err := getPlayerByRarityAndStarRatingRandom(e, usedTeams, cardPackType.Collection.String)
		if err != nil {
			return handleError(c, 10005, "Failed to assign additional random players: "+err.Error(), http.StatusInternalServerError)
		}
		nfts = append(nfts, struct {
			TierData schema.NFTBucket
			Rarity   string
		}{
			TierData: *player,
			Rarity:   rarity,
		})
	}

	assignedPlayers := make([]*schema.AssignedPlayer, len(nfts))
	nftCount, err := e.Store.CountAllEntries()
	if err != nil {
		return handleError(c, 10014, "Failed to retrieve NFT count. Please try again later.", http.StatusInternalServerError)
	}

	// Convert nftCount to int and create the increments slice
	nftCountInt := int(nftCount)
	increments := make([]*big.Int, 7)
	for i := 0; i < 7; i++ {
		increments[i] = big.NewInt(int64(nftCountInt + i + 1))
	}

	for i, nft := range nfts {
		assignedPlayer := &schema.AssignedPlayer{
			ID:          uuid.New().String(),
			PlayerNFTID: null.StringFrom(nft.TierData.ID),
			NFTID:       null.StringFrom(increments[i].String()),
			Rarity:      null.StringFrom(nft.Rarity),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		err := e.Store.DeductFromRarityLimit(nft.TierData.ID, nft.Rarity, 1)
		if err != nil {
			return handleError(c, 10010, "Failed to deduct from rarity limit. Please try again later.", http.StatusInternalServerError)
		}

		if _, err := e.Store.CreateAssignedPlayer(assignedPlayer); err != nil {
			return handleError(c, 10009, "Failed to assign player to user: "+err.Error(), http.StatusInternalServerError)
		}
		assignedPlayers[i] = assignedPlayer
	}

	assignedPlayersJSONStrings := make([]string, len(assignedPlayers))
	for i, player := range assignedPlayers {
		playerJSON, err := json.Marshal(player)
		if err != nil {
			return handleError(c, 10013, "Failed to serialize player information. Please try again later.", http.StatusInternalServerError)
		}
		assignedPlayersJSONStrings[i] = string(playerJSON)
	}

	if err := e.Store.SetAssignedCardPackOpened(cardPackID); err != nil {
		return handleError(c, 10015, "Failed to mark the card pack as opened. Please contact support for further assistance.", http.StatusInternalServerError)
	}

	// Collect all UUIDs from assigned players
	playerUUIDs := make([]string, len(assignedPlayers))
	for i, player := range assignedPlayers {
		playerUUIDs[i] = player.ID
	}
	playersWithNFTs, err := e.Store.GetAssignedPlayersWithNFTDetails(playerUUIDs)
	if err != nil {
		return handleError(c, 10016, "Failed to retrieve players with NFT details. Please contact support for further assistance.", http.StatusInternalServerError)
	}
	if err := e.Store.UpdateAssignedCardPackCardIds(cardPackID, playerUUIDs); err != nil {
		return handleError(c, 10016, "Failed to update card IDs of the assigned card pack. Please contact support for further assistance.", http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Card pack opened successfully.",
		"players": playersWithNFTs,
	})
}

// getPlayerByRarityAndRandomStarRating retrieves a player by rarity and assigns a random star rating
func getPlayerByRarityAndRandomStarRating(e *Env, rarity string, usedTeams map[string]struct{}, cardpackCollection string) (*schema.NFTBucket, error) {
	starRating := selectRandomStarRating()
	return getPlayerByRarityAndStarRating(e, rarity, starRating, usedTeams, cardpackCollection)
}

// getPlayerByStarRatingAndRandomRarity retrieves a player by star rating and assigns a random rarity
func getPlayerByStarRatingAndRandomRarity(e *Env, starRating int, usedTeams map[string]struct{}, cardpackCollection string) (*schema.NFTBucket, string, error) {
	rarity := selectRandomRarity()
	player, err := getPlayerByRarityAndStarRating(e, rarity, starRating, usedTeams, cardpackCollection)
	return player, rarity, err
}

// selectRandomRarity selects a rarity based on weighted probabilities
func selectRandomRarity() string {
	rarityProbabilities := map[string]float64{
		"common":     0.375,
		"uncommon":   0.325,
		"rare":       0.2,
		"ultra_rare": 0.08,
		"legendary":  0.02,
	}
	r := rand.Float64()
	accumulated := 0.0
	for rarity, probability := range rarityProbabilities {
		accumulated += probability
		if r < accumulated {
			return rarity
		}
	}
	return "common"
}

// selectRandomStarRating selects a star rating based on weighted probabilities
func selectRandomStarRating() int {
	starRatingProbabilities := map[int]float64{
		1: 0.2,
		2: 0.25,
		3: 0.35,
		4: 0.15,
		5: 0.05,
	}
	r := rand.Float64()
	accumulated := 0.0
	for starRating, probability := range starRatingProbabilities {
		accumulated += probability
		if r < accumulated {
			return starRating
		}
	}
	return 1
}

func (e *Env) OpenCollection(c echo.Context) error {
	userID := userID(c)
	cardPackID := c.Param("cardPackID")

	user, err := e.Store.GetUserByID(userID)
	if user == nil {
		return handleError(c, 10000, "User not found. Please log in and try again.", http.StatusNotFound)
	}

	if err != nil {
		return handleError(c, 10000, "Failed to retrieve user information. Please try again later.", http.StatusInternalServerError)
	}

	if user.WalletAddress.IsZero() {
		return handleError(c, 10000, "Sign up to open the card pack.", http.StatusNotFound)
	}

	acp, err := e.Store.GetAssignedCardPackByID(cardPackID)
	if err != nil {
		return handleError(c, 10001, "The requested card pack was not found. Please check the card pack ID and try again.", http.StatusNotFound)
	}
	if acp.UserID != userID || acp.Opened {
		return handleError(c, 10002, "Unauthorized access or the card pack has already been opened. Please verify your access rights or card pack status.", http.StatusUnauthorized)
	}

	cardPackType, err := e.Store.GetCardPackByID(acp.CardPackTypeID)
	if err != nil {
		return handleError(c, 10003, "Failed to retrieve card pack type information. Please try again later or contact support if the issue persists.", http.StatusInternalServerError)
	}
	if cardPackType.Collection.String == "" {
		return handleError(c, 10006, "Collection not found. Please contact support for further assistance.", http.StatusNotFound)
	}
	if cardPackType.Collection.String == "kickoff" {
		return e.OpenCardPack(c)
	}
	if cardPackType.Collection.String == "season" {
		return e.NewOpenCardPack(c)
	}

	return handleError(c, 10007, "Unknown collection type. Please contact support.", http.StatusBadRequest)
}
