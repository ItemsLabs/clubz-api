package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

type GraphQLResponse struct {
	Data struct {
		User struct {
			NftBalance int `json:"nftBalance"`
			Nfts       []struct {
				TokenID string `json:"tokenID"`
			} `json:"nfts"`
		} `json:"user"`
	} `json:"data"`
}

type NFT struct {
	TokenID string `json:"tokenId"`
	// Other fields...
}

// Struct to parse the Alchemy API response
type NFTResponse struct {
	OwnedNfts []NFT  `json:"ownedNfts"`
	PageKey   string `json:"pageKey,omitempty"` // Add this line
}

// CreateCardPack godoc
// @Summary Create a card pack
// @Description Create a card pack in the database with provided details.
// @Accept  json
// @Produce  json
// @Param   name        body    string     true        "Name of the card pack"
// @Param   description body    string     true        "Description of the card pack"
// @Param   price       body    int        true        "Price of the card pack"
// @Param   image       body    string     true        "URL to the image of the card pack"
// @Success 200 {object} CardPackType  "Returns the created card pack"
// @Failure 400 {object} string        "Returns error message if creation fails"
// @Router /card_packs [post]
func (e *Env) CreateCardPack(c echo.Context) error {
	var cardPack schema.CardPackType
	if err := c.Bind(&cardPack); err != nil {
		log.Printf("Error creating card pack: %v, CardPack: %+v", err, cardPack)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	cardPack.CreatedAt = time.Now()
	cardPack.UpdatedAt = time.Now()

	if err := e.Store.CreateCardPack(&cardPack); err != nil {
		log.Printf("Error creating card pack: %v, CardPack: %+v", err, cardPack)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Failed to create card pack: %s", err.Error())})
	}

	return c.JSON(http.StatusCreated, cardPack)
}

// GetAssignedCardPacksByUserIDHandler godoc
// @Summary Get assigned card packs by user ID
// @Description Retrieve all assigned card packs for a given user ID.
// @Accept  json
// @Produce  json
// @Param   user_id  path   string  true  "User ID"
// @Success 200 {array} schema.AssignedCardPack  "Returns a list of assigned card packs"
// @Failure 400 {object} string                  "Returns error message if retrieval fails"
// @Router /users/{user_id}/assigned_card_packs [get]
func (e *Env) GetAssignedCardPacksByUserIDHandler(c echo.Context) error {
	userID := userID(c)
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	assignedCardPacks, err := e.Store.GetAssignedCardPacksByUserID(userID)
	if err != nil {
		log.Printf("Error retrieving assigned card packs by user ID: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Failed to retrieve assigned card packs: %s", err.Error())})
	}

	type CardPackTypeResponse struct {
		ID          string      `boil:"id" json:"id" toml:"id" yaml:"id"`
		CreatedAt   time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
		UpdatedAt   time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
		Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
		Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
		Image       null.String `boil:"image" json:"image,omitempty" toml:"image" yaml:"image,omitempty"`
		PackLimits  null.Int    `boil:"pack_limits" json:"pack_limits,omitempty" toml:"pack_limits" yaml:"pack_limits,omitempty"`
	}

	type AssignedCardPackResponse struct {
		ID             string                `json:"id"`
		CreatedAt      time.Time             `json:"created_at"`
		UpdatedAt      time.Time             `json:"updated_at"`
		Opened         bool                  `json:"opened"`
		CardPackTypeID string                `json:"card_pack_type_id"`
		UserID         string                `json:"user_id"`
		OpenedAt       *time.Time            `json:"opened_at,omitempty"`
		Revealed       bool                  `json:"revealed"`
		RevealedAt     *time.Time            `json:"revealed_at,omitempty"`
		CardIds        *json.RawMessage      `json:"card_ids,omitempty"`
		TransactionID  *string               `json:"transaction_id_id,omitempty"`
		CardPackType   *CardPackTypeResponse `json:"card_pack_type,omitempty"`
	}

	var response []AssignedCardPackResponse
	for _, acp := range assignedCardPacks {
		var cardIds *json.RawMessage
		if acp.CardIds.Valid {
			cardIds = (*json.RawMessage)(&acp.CardIds.JSON)
		}

		resp := AssignedCardPackResponse{
			ID:             acp.ID,
			CreatedAt:      acp.CreatedAt,
			UpdatedAt:      acp.UpdatedAt,
			Opened:         acp.Opened,
			CardPackTypeID: acp.CardPackTypeID,
			UserID:         acp.UserID,
			OpenedAt:       acp.OpenedAt.Ptr(),
			Revealed:       acp.Revealed,
			RevealedAt:     acp.RevealedAt.Ptr(),
			CardIds:        cardIds,
			TransactionID:  acp.StoreTransactionID.Ptr(),
		}
		if acp.R != nil && acp.R.CardPackType != nil {
			resp.CardPackType = &CardPackTypeResponse{
				ID:          acp.R.CardPackType.ID,
				CreatedAt:   acp.R.CardPackType.CreatedAt,
				UpdatedAt:   acp.R.CardPackType.UpdatedAt,
				Name:        acp.R.CardPackType.Name,
				Description: acp.R.CardPackType.Description,
				Image:       acp.R.CardPackType.Image,
				PackLimits:  acp.R.CardPackType.PackLimits,
			}
		}
		response = append(response, resp)
	}

	return c.JSON(http.StatusOK, response)
}

// @Summary Assign a card pack to a player
// @Description Assign an existing card pack to a player.
// @Accept  json
// @Produce  json
// @Param   user_id        path    string     true  "User ID"
// @Param   card_pack_type_id   body    string     true  "Card Pack Type ID"
// @Success 200 {object} schema.AssignedCardPack  "Returns the assigned card pack"
// @Failure 400 {object} string                  "Returns error message if assignment fails"
// @Router /users/{user_id}/assign_card_pack [post]
func (e *Env) AssignCardPackToPlayerHandler(c echo.Context) error {
	userID := userID(c)
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	type RequestBody struct {
		CardPackTypeID string `json:"card_pack_type_id" validate:"required"`
	}

	var requestBody RequestBody
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	assignedCardPack := &schema.AssignedCardPack{
		ID:             uuid.New().String(),
		UserID:         userID,
		CardPackTypeID: requestBody.CardPackTypeID,
		Opened:         false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := e.Store.DeductFromPackLimit(requestBody.CardPackTypeID, 1)
	if err != nil {
		log.Printf("Error deducting from pack limit: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("Failed to assign card pack: %s", err.Error())})
	}
	if _, err := e.Store.CreateAssignedCardPack(assignedCardPack); err != nil {
		log.Printf("Error assigning card pack to player: %v", err)
		errMap := map[string]string{"error": fmt.Sprintf("Failed to assign card pack: %s", err.Error())}
		if err := e.Store.RestockPackLimit(requestBody.CardPackTypeID, 1); err != nil {
			errMap["error"] = errMap["error"] + " - RestockPackLimit: " + err.Error()
		}
		return c.JSON(http.StatusInternalServerError, errMap)
	}

	return c.JSON(http.StatusCreated, assignedCardPack)
}

// GetCardPackWithNFTDetailsHandler godoc
// @Summary Get card pack with NFT details
// @Description Retrieve card pack and related NFT details using card pack ID
// @Accept  json
// @Produce  json
// @Param   cardPackID  path   string  true  "Card Pack ID"
// @Success 200 {object} map[string]interface{}  "Returns card pack with NFT details"
// @Failure 400 {object} string                  "Returns error message if retrieval fails"
// @Router /card_packs/{cardPackID}/details [get]
func (e *Env) GetCardPackWithNFTDetailsHandler(c echo.Context) error {
	cardPackID := c.Param("cardPackID")
	if cardPackID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Card Pack ID is required"})
	}

	// Fetch the assigned card pack
	assignedCardPack, err := e.Store.GetAssignedCardPackByID(cardPackID)
	if err != nil {
		log.Printf("Error retrieving assigned card pack by ID: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve assigned card pack"})
	}

	// Extract card IDs from the assigned card pack
	cardIDs := assignedCardPack.CardIds
	if cardIDs.IsZero() {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Card IDs are empty"})
	}

	// Convert JSON array to slice of strings
	var cardUUIDs []string
	if err := cardIDs.Unmarshal(&cardUUIDs); err != nil {
		log.Printf("Error unmarshaling card IDs: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse card IDs"})
	}

	// Fetch assigned players and their corresponding NFT details
	playersWithNFTs, err := e.Store.GetAssignedPlayersWithNFTDetails(cardUUIDs)
	if err != nil {
		log.Printf("Error retrieving assigned players with NFT details: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve assigned players with NFT details"})
	}

	// Prepare the response
	var nftDetails []map[string]interface{}
	for _, item := range playersWithNFTs {
		nftDetails = append(nftDetails, map[string]interface{}{
			"player": item.Player,
			"nft":    item.NFT,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"assignedCardPack": assignedCardPack,
		"nftDetails":       nftDetails,
	})
}

// GetMintedPlayersFromCardPacksHandler godoc
// @Summary Get minted players from card packs
// @Description Retrieve all players minted from card packs for a given user ID, filtered by opened_at timestamp.
// @Accept  json
// @Produce  json
// @Param   user_id  path   string  true  "User ID"
// @Success 200 {array} map[string]interface{}  "Returns a list of minted players with NFT details"
// @Failure 400 {object} string                  "Returns error message if retrieval fails"
// @Router /users/{user_id}/minted_players [get]
func (e *Env) GetMintedPlayersFromCardPacksHandler(c echo.Context) error {
	userID := userID(c)
	log.Printf("Starting GetMintedPlayersFromCardPacksHandler for User ID: %s", userID)
	if userID == "" {
		log.Printf("Error: User ID is empty")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		log.Printf("Error retrieving user by ID %s: %v", userID, err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User not found"})
	}

	assignedCardPacks, err := e.Store.GetAssignedCardPacksByUserID(userID)
	if err != nil {
		log.Printf("Error retrieving assigned card packs by user ID %s: %v", userID, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve assigned card packs"})
	}

	if len(assignedCardPacks) == 0 {
		log.Printf("No assigned card packs found for user ID %s", userID)
		return c.JSON(http.StatusOK, []map[string]interface{}{})
	}

	if assignedCardPacks[0].UserID != user.ID {
		log.Printf("Unauthorized access attempt for user ID %s", userID)
		return handleError(c, 10002, "Unauthorized access or the card pack has already been opened. Please verify your access rights or card pack status.", http.StatusUnauthorized)
	}

	var mintedPlayers []map[string]interface{}
	// Use the subgraph to get the balance and NFT IDs
	nftIDs, err := getNFTDataFromCurl(user.WalletAddress.String)
	if err != nil {
		log.Printf("Error fetching NFT data from subgraph for address %s: %v", user.WalletAddress.String, err)

	}

	playersWithNFTs, err := e.Store.GetAssignedPlayersByNFTIDs(nftIDs)
	if err != nil {
		log.Printf("Error retrieving assigned players by NFT IDs: %v", err)

	}

	for _, playerWithNFT := range playersWithNFTs {
		playerNFTID := playerWithNFT.Player.PlayerNFTID.String
		playerRarity := playerWithNFT.Player.Rarity.String
		nftDetails, err := e.Store.GetNFTBucketByIDandRarity(playerNFTID, playerRarity)
		if err != nil {
			log.Printf("Error retrieving NFT bucket by ID: %v", err)
			continue
		}

		mintedPlayers = append(mintedPlayers, map[string]interface{}{
			"player": playerWithNFT.Player,
			"nft":    playerWithNFT.NFT,
			"name":   nftDetails.Name,
		})
	}
	for _, cardPack := range assignedCardPacks {
		log.Printf("Processing card pack ID %s for user ID %s", cardPack.ID, userID)
		if cardPack.OpenedAt.Valid {
			if !cardPack.Revealed {
				log.Printf("Processing card pack opened within 48 hours: %v", cardPack.ID)
				cardPackWithNFTDetails, err := e.Store.GetAssignedCardPackByID(cardPack.ID)
				if err != nil {
					log.Printf("Error retrieving assigned card pack by ID: %v", err)
					continue
				}

				cardIDs := cardPackWithNFTDetails.CardIds
				if cardIDs.IsZero() {
					log.Println("Card IDs are zero")
					continue
				}

				var cardUUIDs []string
				if err := cardIDs.Unmarshal(&cardUUIDs); err != nil {
					log.Printf("Error unmarshaling card IDs: %v", err)
					continue
				}

				playersWithNFTs, err := e.Store.GetAssignedPlayersWithNFTDetails(cardUUIDs)
				if err != nil {
					log.Printf("Error retrieving assigned players with NFT details: %v", err)
					continue
				}

				for _, playerWithNFT := range playersWithNFTs {
					playerNFTID := playerWithNFT.Player.PlayerNFTID.String
					playerRarity := playerWithNFT.Player.Rarity.String
					nftDetails, err := e.Store.GetNFTBucketByIDandRarity(playerNFTID, playerRarity)
					if err != nil {
						log.Printf("Error retrieving NFT bucket by ID: %v", err)
						continue
					}

					mintedPlayers = append(mintedPlayers, map[string]interface{}{
						"player": playerWithNFT.Player,
						"nft":    playerWithNFT.NFT,
						"name":   nftDetails.Name,
					})
				}
			} else {
				log.Printf("CardPack already minted: %v", cardPack.ID)

			}
		}
	}

	if len(mintedPlayers) == 0 {
		log.Printf("No minted players found after processing all card packs for user ID %s", userID)
	} else {
		log.Printf("Found %d minted players for user ID %s", len(mintedPlayers), userID)
	}

	return c.JSON(http.StatusOK, mintedPlayers)
}

// GetMintedPlayersFromCardPacksHandler godoc
// @Summary Get minted players from card packs
// @Description Retrieve all players minted from card packs for a given user ID, filtered by opened_at timestamp.
// @Accept  json
// @Produce  json
// @Param   user_id  path   string  true  "User ID"
// @Success 200 {array} map[string]interface{}  "Returns a list of minted players with NFT details"
// @Failure 400 {object} string                  "Returns error message if retrieval fails"
// @Router /users/{user_id}/minted_players [get]
func (e *Env) GetMintedPlayersFromCardPacksFilteredHandler(c echo.Context) error {
	userID := userID(c)
	log.Printf("Starting GetMintedPlayersFromCardPacksHandler for User ID: %s", userID)
	if userID == "" {
		log.Printf("Error: User ID is empty")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User ID is required"})
	}

	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		log.Printf("Error retrieving user by ID %s: %v", userID, err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User not found"})
	}

	assignedCardPacks, err := e.Store.GetAssignedCardPacksByUserID(userID)
	if err != nil {
		log.Printf("Error retrieving assigned card packs by user ID %s: %v", userID, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve assigned card packs"})
	}

	if len(assignedCardPacks) == 0 {
		log.Printf("No assigned card packs found for user ID %s", userID)
		return c.JSON(http.StatusOK, []map[string]interface{}{})
	}

	if assignedCardPacks[0].UserID != user.ID {
		log.Printf("Unauthorized access attempt for user ID %s", userID)
		return handleError(c, 10002, "Unauthorized access or the card pack has already been opened. Please verify your access rights or card pack status.", http.StatusUnauthorized)
	}

	var mintedPlayers []map[string]interface{}
	// Use the subgraph to get the balance and NFT IDs
	nftIDs, err := getNFTDataFromCurl(user.WalletAddress.String)
	if err != nil {
		log.Printf("Error fetching NFT data from subgraph for address %s: %v", user.WalletAddress.String, err)

	}

	playersWithNFTs, err := e.Store.GetAssignedPlayersByNFTIDs(nftIDs)
	if err != nil {
		log.Printf("Error retrieving assigned players by NFT IDs: %v", err)

	}

	for _, playerWithNFT := range playersWithNFTs {
		playerNFTID := playerWithNFT.Player.PlayerNFTID.String
		playerRarity := playerWithNFT.Player.Rarity.String
		nftDetails, err := e.Store.GetNFTBucketByIDandRarityFiltered(playerNFTID, playerRarity)
		if err != nil {
			log.Printf("Error retrieving NFT bucket by ID: %v", err)
			continue
		}

		mintedPlayers = append(mintedPlayers, map[string]interface{}{
			"player": playerWithNFT.Player,
			"nft":    nftDetails,
			"name":   nftDetails.Name,
		})
	}
	for _, cardPack := range assignedCardPacks {
		log.Printf("Processing card pack ID %s for user ID %s", cardPack.ID, userID)
		if cardPack.OpenedAt.Valid {
			if !cardPack.Revealed {
				log.Printf("Processing card pack opened within 48 hours: %v", cardPack.ID)
				cardPackWithNFTDetails, err := e.Store.GetAssignedCardPackByID(cardPack.ID)
				if err != nil {
					log.Printf("Error retrieving assigned card pack by ID: %v", err)
					continue
				}

				cardIDs := cardPackWithNFTDetails.CardIds
				if cardIDs.IsZero() {
					log.Println("Card IDs are zero")
					continue
				}

				var cardUUIDs []string
				if err := cardIDs.Unmarshal(&cardUUIDs); err != nil {
					log.Printf("Error unmarshaling card IDs: %v", err)
					continue
				}

				playersWithNFTs, err := e.Store.GetAssignedPlayersWithNFTDetails(cardUUIDs)
				if err != nil {
					log.Printf("Error retrieving assigned players with NFT details: %v", err)
					continue
				}

				for _, playerWithNFT := range playersWithNFTs {
					playerNFTID := playerWithNFT.Player.PlayerNFTID.String
					playerRarity := playerWithNFT.Player.Rarity.String
					nftDetails, err := e.Store.GetNFTBucketByIDandRarityFiltered(playerNFTID, playerRarity)
					if err != nil {
						log.Printf("Error retrieving NFT bucket by ID: %v", err)
						continue
					}

					mintedPlayers = append(mintedPlayers, map[string]interface{}{
						"player": playerWithNFT.Player,
						"nft":    nftDetails,
						"name":   nftDetails.Name,
					})
				}
			} else {
				log.Printf("CardPack already minted: %v", cardPack.ID)

			}
		}
	}

	if len(mintedPlayers) == 0 {
		log.Printf("No minted players found after processing all card packs for user ID %s", userID)
	} else {
		log.Printf("Found %d minted players for user ID %s", len(mintedPlayers), userID)
	}

	return c.JSON(http.StatusOK, mintedPlayers)
}

func getNFTDataFromCurl(walletAddress string) ([]string, error) {
	alchemyURL := os.Getenv("ALCHEMY_URL")
	contractAddress := os.Getenv("CONTRACT_ADDRESS")

	if alchemyURL == "" {
		return nil, fmt.Errorf("ALCHEMY_URL environment variable is not set")
	}

	if contractAddress == "" {
		return nil, fmt.Errorf("CONTRACT_ADDRESS environment variable is not set")
	}

	var nftIDs []string
	pageKey := ""

	for {
		url := fmt.Sprintf("%s/nft/v3/FfCMsxRJDQaIcai9atLm1gDgVRMGTXd5/getNFTsForOwner?owner=%s&withMetadata=true&pageSize=100&contract=%s", alchemyURL, walletAddress, contractAddress)
		if pageKey != "" {
			url = fmt.Sprintf("%s&pageKey=%s", url, pageKey)
		}

		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error making GET request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("received non-ok status code: %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %v", err)
		}

		var nftResponse NFTResponse
		err = json.Unmarshal(body, &nftResponse)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling response: %v", err)
		}

		for _, nft := range nftResponse.OwnedNfts {
			nftIDs = append(nftIDs, nft.TokenID)
		}

		// Check if there is a next page
		if nftResponse.PageKey == "" {
			break
		}

		pageKey = nftResponse.PageKey
	}

	log.Printf("NFT IDs: %v", nftIDs)
	return nftIDs, nil
}

func (e *Env) GetPackLimitsByCardPackCodeHandler(c echo.Context) error {
	code := c.Param("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Card Pack Code is required"})
	}

	limit, err := e.Store.GetPackLimitsByCardPackCode(code)
	if err != nil {
		log.Printf("Error retrieving pack limits by card pack name: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve pack limits"})
	}

	return c.JSON(http.StatusOK, map[string]int{"pack_limits": limit})
}
