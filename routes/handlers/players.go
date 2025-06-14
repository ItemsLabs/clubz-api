package handlers

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/apiconv"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/palantir/stacktrace"
	"github.com/volatiletech/null/v8"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Test struct {
	// This is a test
	Test model.PregamePlayerStat
}

type PlayerService struct {
	store PlayerStore
}

func NewPlayerService(store PlayerStore) *PlayerService {
	return &PlayerService{store: store}
}

type PlayerStore interface {
	GetUserByID(id string) (*schema.User, error)
	GetMatchByID(matchID string) (*schema.Match, error)
	GetLastMatchPlayers(playerID string, limit int, mods ...qm.QueryMod) (schema.MatchPlayerSlice, error)
	GetPercentOfPicks(matchID, playerID string) (float64, error)
	GetAveragePointsDistribution(playerID string, matches []string) ([]*database.PointBucket, error)
	GetPointsInInterval(matchID, playerID string, startTime, endTime time.Time) (int, error)
	GetActionSummary(matchID, playerID string) ([]*database.ActionSummary, error)
}

// GetPlayerByID godoc
// @Summary Get player by ID
// @Description Get player by ID
// @ID get-player-by-id
// @Produce json
// @Param id path string true "Player ID"
// @Success 200 {object} schema.Player
// @Router /players/{id} [get]
func (e *Env) GetPlayerByID(c echo.Context) error {
	player, err := e.Store.GetPlayerByID(c.Param("id"))
	if err != nil {
		return err
	}

	return e.RespondSuccess(c, apiconv.ToPlayer(player))
}

// GetPregamePlayerStats godoc
// @Summary Get pregame player stats
// @Description Get pregame player stats for a specific player in a match
// @ID get-pregame-player-stats
// @Produce json
// @Param match_id path string true "Match ID"
// @Param player_id path string true "Player ID"
// @Success 200 {object} model.PregamePlayerStat
// @Router /matches/{match_id}/players/{player_id}/pregame [get]
func (s *PlayerService) GetPregamePlayerStats(c echo.Context) error {
	matchID := c.Param("match_id")
	playerID := c.Param("player_id")

	currentUser, err := s.store.GetUserByID(userID(c))
	if err != nil {
		return err
	}

	if currentUser.SubscriptionTier != database.SubscriptionTierPremium {
		return ErrNotAuthorized
	}

	pickPercent, err := s.store.GetPercentOfPicks(matchID, playerID)
	if err != nil {
		return stacktrace.Propagate(err, "cannot calculate percent of picks %v -%v", matchID, playerID)
	}

	matchPlayers, err := s.store.GetLastMatchPlayers(playerID, 5, qm.Load("Match.HomeTeam"), qm.Load("Match.AwayTeam"))
	if err != nil {
		return stacktrace.Propagate(err, "cannot get last match players %v", playerID)
	}

	for _, mp := range matchPlayers {
		if mp.R.Match.AwayTeamID == mp.TeamID {
			mp.R.Team = mp.R.Match.R.HomeTeam
		} else {
			mp.R.Team = mp.R.Match.R.AwayTeam
		}
	}

	matchIds := make([]string, 0, len(matchPlayers))
	for _, mp := range matchPlayers {
		matchIds = append(matchIds, mp.MatchID)
	}

	pointDistribution, err := s.store.GetAveragePointsDistribution(playerID, matchIds)
	if err != nil {
		return stacktrace.Propagate(err, "cannot get average points distribution, %v, %v", playerID, matchIds)
	}

	return RespondSuccess(c, apiconv.ToPregamePlayerStat(matchPlayers, pickPercent, pointDistribution))
}

// GetLivePlayerStats godoc
// @Summary Get live player stats
// @Description Get live player stats
// @ID get-live-player-stats
// @Produce json
// @Param match_id path string true "Match ID"
// @Param player_id path string true "Player ID"
// @Success 200 {object} model.LiveGamePlayerStat
// @Router /matches/{match_id}/players/{player_id}/live [get]
func (s *PlayerService) GetLivePlayerStats(c echo.Context) error {
	matchID := c.Param("match_id")
	playerID := c.Param("player_id")
	now := time.Now()

	currentUser, err := s.store.GetUserByID(userID(c))
	if err != nil {
		return err
	}

	if currentUser.SubscriptionTier != database.SubscriptionTierPremium {
		return ErrNotAuthorized
	}

	// calculate pick percent
	pickPercent, err := s.store.GetPercentOfPicks(matchID, playerID)
	if err != nil {
		return stacktrace.Propagate(err, "cannot calculate percent of picks %v - %v", matchID, playerID)
	}

	pointDistribution, err := s.store.GetAveragePointsDistribution(playerID, []string{matchID})
	if err != nil {
		return stacktrace.Propagate(err, "cannot get average points distribution, %v, %v", playerID, matchID)
	}

	last10MinPoints, err := s.store.GetPointsInInterval(matchID, playerID, now.Add(-time.Minute*10), now)
	if err != nil {
		return err
	}

	actionSummary, err := s.store.GetActionSummary(matchID, playerID)
	if err != nil {
		return err
	}

	return RespondSuccess(c, apiconv.ToLiveGamePlayerStat(pickPercent, pointDistribution, last10MinPoints, actionSummary))
}

// parseCSV parses the CSV file and returns slices of teams and players.

func parseCSVFromFile(filePath string) ([]*schema.Country, []*schema.Team, []*schema.NFTBucket, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to open CSV file '%s': %w", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to read records from file '%s': %w", filePath, err)
	}

	var nfts []*schema.NFTBucket
	teamMap := make(map[string]*schema.Team)
	countryMap := make(map[string]*schema.Country)

	for i, record := range records {
		if i < 1 { // Skip the first row
			continue
		}

		teamName := record[4]
		if _, exists := teamMap[teamName]; !exists {
			teamMap[teamName] = &schema.Team{
				Name: teamName,
				// Populate other team fields as necessary
			}
		}

		countryName := record[2]
		if _, exists := countryMap[countryName]; !exists {
			countryMap[countryName] = &schema.Country{
				Name: countryName,
				// Populate other country fields as necessary
			}
		}

		nft := &schema.NFTBucket{
			ID:                    uuid.New().String(),
			Name:                  record[0],
			Age:                   null.IntFrom(parseToInt(record[1])),
			Nationality:           null.StringFrom(record[2]),
			Position:              record[3],
			TeamID:                null.StringFrom(teamName),
			GamePosition:          record[5],
			StarRating:            parseToFloat(record[6]),
			CommonStopping:        null.Float64From(parseToFloat(record[14])),
			CommonClaiming:        null.Float64From(parseToFloat(record[15])),
			CommonDistribution:    null.Float64From(parseToFloat(record[16])),
			CommonShooting:        null.Float64From(parseToFloat(record[17])),
			CommonDribbling:       null.Float64From(parseToFloat(record[18])),
			CommonPassing:         null.Float64From(parseToFloat(record[19])),
			CommonDefence:         null.Float64From(parseToFloat(record[20])),
			UncommonStopping:      null.Float64From(parseToFloat(record[22])),
			UncommonClaiming:      null.Float64From(parseToFloat(record[23])),
			UncommonDistribution:  null.Float64From(parseToFloat(record[24])),
			UncommonShooting:      null.Float64From(parseToFloat(record[25])),
			UncommonDribbling:     null.Float64From(parseToFloat(record[26])),
			UncommonPassing:       null.Float64From(parseToFloat(record[27])),
			UncommonDefence:       null.Float64From(parseToFloat(record[28])),
			RareStopping:          null.Float64From(parseToFloat(record[30])),
			RareClaiming:          null.Float64From(parseToFloat(record[31])),
			RareDistribution:      null.Float64From(parseToFloat(record[32])),
			RareShooting:          null.Float64From(parseToFloat(record[33])),
			RareDribbling:         null.Float64From(parseToFloat(record[34])),
			RarePassing:           null.Float64From(parseToFloat(record[35])),
			RareDefence:           null.Float64From(parseToFloat(record[36])),
			UltraRareStopping:     null.Float64From(parseToFloat(record[38])),
			UltraRareClaiming:     null.Float64From(parseToFloat(record[39])),
			UltraRareDistribution: null.Float64From(parseToFloat(record[40])),
			UltraRareShooting:     null.Float64From(parseToFloat(record[41])),
			UltraRareDribbling:    null.Float64From(parseToFloat(record[42])),
			UltraRarePassing:      null.Float64From(parseToFloat(record[43])),
			UltraRareDefence:      null.Float64From(parseToFloat(record[44])),
			LegendaryStopping:     null.Float64From(parseToFloat(record[46])),
			LegendaryClaiming:     null.Float64From(parseToFloat(record[47])),
			LegendaryDistribution: null.Float64From(parseToFloat(record[48])),
			LegendaryShooting:     null.Float64From(parseToFloat(record[49])),
			LegendaryDribbling:    null.Float64From(parseToFloat(record[50])),
			LegendaryPassing:      null.Float64From(parseToFloat(record[51])),
			LegendaryDefence:      null.Float64From(parseToFloat(record[52])),
			CommonLimit:           null.NewInt(128, true),
			UncommonLimit:         null.NewInt(64, true),
			RareLimit:             null.NewInt(32, true),
			UltraRareLimit:        null.NewInt(16, true),
			LegendaryLimit:        null.NewInt(5, true),
			CommonImage:           null.StringFrom(record[64]),
			CommonMetadata:        null.StringFrom(record[59]),
			UncommonImage:         null.StringFrom(record[65]),
			UncommonMetadata:      null.StringFrom(record[60]),
			RareImage:             null.StringFrom(record[66]),
			RareMetadata:          null.StringFrom(record[61]),
			UltraRareImage:        null.StringFrom(record[67]),
			UltraRareMetadata:     null.StringFrom(record[62]),
			LegendaryImage:        null.StringFrom(record[68]),
			LegendaryMetadata:     null.StringFrom(record[63]),
			OptaID:                null.StringFrom(record[55]),
			PlayersGroup:          null.StringFrom("season"),
			CreatedAt:             time.Now(),
			UpdatedAt:             time.Now(),
		}

		nfts = append(nfts, nft)
	}

	var teams []*schema.Team
	for _, team := range teamMap {
		teams = append(teams, team)
	}

	var countries []*schema.Country
	for _, country := range countryMap {
		countries = append(countries, country)
	}

	return countries, teams, nfts, nil
}

func parseToInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return val
}

func parseToFloat(str string) float64 {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return val
}

func (e *Env) ImportNFTsHandler(c echo.Context) error {
	filePath := "./players_season.csv"

	countries, teams, nfts, err := parseCSVFromFile(filePath)
	if err != nil {
		return fmt.Errorf("error parsing CSV from file: %w", err)
	}

	for _, country := range countries {
		newCountry := &schema.Country{
			ID:        uuid.New().String(),
			Name:      country.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := e.Store.AddCountry(newCountry); err != nil {
			return fmt.Errorf("error creating country '%s': %w", country.Name, err)
		}
	}

	for _, team := range teams {
		newTeam := &schema.Team{
			ID:        uuid.New().String(),
			Name:      team.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if err := e.Store.AddTeam(newTeam); err != nil {
			return fmt.Errorf("error creating team '%s': %w", team.Name, err)
		}
	}

	for _, nft := range nfts {
		team, err := e.Store.GetTeamByName(nft.TeamID.String)
		if err != nil {
			return fmt.Errorf("error processing team '%s': %w", nft.TeamID.String, err)
		}
		country, err := e.Store.GetCountryByName(nft.Nationality.String)
		if err != nil {
			return fmt.Errorf("error processing country '%s': %w", nft.Nationality.String, err)
		}
		nft.Nationality = null.StringFrom(country.ID)
		nft.TeamID = null.StringFrom(team.ID)

		if _, err := e.Store.CreateNFTBucket(nft); err != nil {
			return fmt.Errorf("error creating NFT bucket for '%s': %w", nft.Name, err)
		}
	}

	return e.RespondSuccess(c, "NFTs imported successfully.")
}
