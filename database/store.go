package database

import (
	"context"
	"errors"
	"time"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/schema"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/model"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/types"
	"github.com/volatiletech/null/v8"
)

const (
	MatchStatusUnknown   = "u"
	MatchStatusWaiting   = "w"
	MatchStatusLineups   = "l"
	MatchStatusGame      = "g"
	MatchStatusEnded     = "e"
	MatchStatusCancelled = "c"

	MatchPeriodPregame    = "p"
	MatchPeriodFirstHalf  = "f"
	MatchPeriodHalfTime   = "h"
	MatchPeriodSecondHalf = "s"
	MatchPeriodBreakX1    = "bx1"
	MatchPeriodFirstExt   = "x1"
	MatchPeriodBreakX2    = "bx2"
	MatchPeriodSecondExt  = "x2"
	MatchPeriodBreakP     = "bp"
	MatchPeriodPenalties  = "pe"
	MatchPeriodPostGame   = "pg"

	MatchTypeUnknown  = 0
	MatchTypeFree     = 1
	MatchTypeCash     = 2
	MatchTypeCashPlus = 3

	GameStatusWaiting  = "w"
	GameStatusGameplay = "g"
	GameStatusFinished = "f"

	MatchEventStatusActive    = 1
	MatchEventStatusCancelled = 2
	MatchEventStatusIgnored   = 3

	MatchHeadlineScreenTypeLobby    = 1
	MatchHeadlineScreenTypeGamePlay = 2
	MatchHeadlineScreenTypeFullTime = 3

	PositionDefender   = "d"
	PositionMidfielder = "m"
	PositionForward    = "f"
	PositionGoalkeeper = "g"
	PositionSubstitute = "s"
	PositionUnknown    = "u"

	SubscriptionTierNone    = 0
	SubscriptionTierPremium = 1
	SubscriptionTierLite    = 2

	PowerUpTypeEvent   = 1
	PowerUpTypeSpecial = 2

	TransactionTypeVirtual = "v"
	TransactionTypeCrypto  = "c"
	TransactionTypeNFT     = "n"

	StoreProductTypeSubscription  = "subscription"
	StoreProductTypeNFT           = "nft"
	StoreProductTypeConsumable    = "consumable"
	StoreProductTypeNonConsumable = "nonconsumable"
)

var UserAlreadyExists = errors.New("db: user already exists")
var UserWithNameAlreadyExists = errors.New("db: user with name already exists")
var ReferralCodeAlreadyExists = errors.New("db: referral code already exists")
var CannotGenerateUniqueReferralCode = errors.New("db: cannot generate unique referral code")
var FollowerAlreadyExists = errors.New("db: follower already exists")

type PointBucket struct {
	Low    int
	High   int
	Points null.Float64
}

type ActionSummary struct {
	Name   string
	Count  int
	Points float64
}

type RarityPercentages struct {
	Common    float64 `json:"common"`
	Uncommon  float64 `json:"uncommon"`
	Rare      float64 `json:"rare"`
	UltraRare float64 `json:"ultra_rare"`
	Legendary float64 `json:"legendary"`
}

type PlayerWithNFTDetails struct {
	Player *schema.AssignedPlayer
	NFT    *schema.NFTBucket
}

type Store interface {
	Transaction(func(s Store) error) error

	LockUser(userID string) (*schema.User, error)
	GetUserByID(id string) (*schema.User, error)
	GetUserByName(name string) (*schema.User, error)
	GetUserByReferralCode(code string) (*schema.User, error)
	CreateUser(user *schema.User) (*schema.User, error)
	UpdateUserName(user *schema.User, t time.Time) (*schema.User, error)
	UpdateUserPayPalEmail(user *schema.User) (*schema.User, error)
	UpdateUserReferrer(user *schema.User) (*schema.User, error)
	UpdateUserReferrerBonusUsed(user *schema.User) error
	UpdateBonusPowerUps(user *schema.User) (*schema.User, error)
	GetUserIDByFirebaseID(firebaseID string) (string, error)
	GetNextUserIndex() (int, error)
	SearchUsers(query string) (schema.UserSlice, error)

	CreateFollower(fromUser, toUser string) error
	DeleteFollower(fromUser, toUser string) error
	GetFollowers(userID string, search string) (schema.UserSlice, error)
	GetFollowings(userID string, search string) (schema.UserSlice, error)
	IsFollowing(fromUser, toUser string) (bool, error)
	ChangeFollowerCount(userID string, diff int) error
	ChangeFollowingCount(userID string, diff int) error

	GetMatchesInPeriod(from, to time.Time) (schema.MatchSlice, error)
	GetMatchesInGameWeek(gameWeek int) (schema.MatchSlice, error)
	GetNextActiveMatch(t time.Time) (*schema.Match, error)
	GetMatchSquad(matchID string) (schema.MatchPlayerSlice, error)
	GetMatchByID(matchID string) (*schema.Match, error)
	GetPlayerCount(matchID string) (int64, error)

	GetMatchPlayers(matchID string) (schema.MatchPlayerSlice, error)
	GetMatchPlayersPPG(matchID string) (map[string]float64, error)

	LockGame(gameID string) (*schema.Game, error)
	GetGamesLaterThan(userID string, from time.Time) (schema.GameSlice, error)
	GetGameByID(id string, userID string) (*schema.Game, error)
	GetGameByUserIDMatchID(userID, matchID string) (*schema.Game, error)
	GetActiveGameIDForMatch(matchID, userID string) (string, error)
	IsGameBelongToUser(gameID, userID string) (bool, error)
	CreateGame(game *schema.Game) (*schema.Game, error)
	GetNumberOfGames(userID string) (int, error)
	CreateGamePick(pick *schema.GamePick) (*schema.GamePick, error)
	GetGamePickByID(id, gameID string) (*schema.GamePick, error)
	GetActiveGamePicksWithPlayer(gameID, playerID string) (schema.GamePickSlice, error)
	UpdateGamePickEndedAt(pick *schema.GamePick) error
	GetGameEvents(gameID string, offset, limit int) (schema.GameEventSlice, error)
	GetLeaderBoardForGame(gameID string) (*schema.MatchLeaderboard, error)
	GetFollowingLeaderboard(matchID, userID string, limit int) (schema.MatchLeaderboardSlice, error)
	GetLeaderboard(matchID string, limit int) (schema.MatchLeaderboardSlice, error)
	GetGameActivePowerUps(gameID string, position int) (schema.GamePowerupSlice, error)
	GetGamePowerUps(gameID string) (schema.GamePowerupSlice, error)
	CreateGamePowerUp(pu *schema.GamePowerup) (*schema.GamePowerup, error)
	GetFinishedGames(userID string, limit int) (schema.GameSlice, error)

	GetMatchHeadlines(matchID string, screenType int) (schema.MatchHeadlineSlice, error)

	GetTeamByID(teamID string) (*schema.Team, error)

	GetPlayerByID(playerID string) (*schema.Player, error)

	GetActions() (schema.ActionSlice, error)
	GetActionByID(actionID string) (*schema.Action, error)
	UpdateAction(actionID string, options *model.UpdateActionRequest) (*schema.Action, error)
	CreateActions(actions *model.CreateActionsRequest) ([]schema.Action, error)
	DeleteAction(actionID string) error

	FindTransaction(matchID, userID string) (*schema.Transaction, error)
	FindRewardTransaction(matchID, userID string) (*schema.Transaction, error)
	CreateTransaction(transaction *schema.Transaction) (*schema.Transaction, error)
	GetTransactionByID(id string) (*schema.Transaction, error)

	GetPowerUpByID(id int) (*schema.Powerup, error)
	GetPowerUps() (schema.PowerupSlice, error)
	GetPowerUpActions() (schema.PowerupActionSlice, error)
	GetSportSubstitutionPowerUp(sportID string) (*schema.Powerup, error)

	CreateAMQPEvent(exchange, typ string, data interface{}) (*schema.AmqpEvent, error)
	CreateAssignedCardPack(acp *schema.AssignedCardPack) (*schema.AssignedCardPack, error)
	GetAssignedCardPackByID(id string) (*schema.AssignedCardPack, error)
	GetAllAssignedCardPacks() (schema.AssignedCardPackSlice, error)
	UpdateAssignedCardPack(acp *schema.AssignedCardPack) error
	DeleteAssignedCardPack(id string) error
	GetAssignedCardPackByStoreProductTransactionID(storeTransactionID string) (*schema.AssignedCardPack, error)
	RefundAssignedCardPackByID(id string) error
	GetAssignedCardPacksByUserID(userID string) (schema.AssignedCardPackSlice, error)
	GetAssignedCardPacksByCardPackTypeID(cardPackTypeID string) (schema.AssignedCardPackSlice, error)
	OpenAssignedCardPack(id string) error
	FilterAssignedCardPacksByDateRange(from, to time.Time) (schema.AssignedCardPackSlice, error)
	CountUnopenedAssignedCardPacks(userID string) (int64, error)
	GetRecentAssignedCardPacks(limit int) (schema.AssignedCardPackSlice, error)
	BulkUpdateOpenedStatus(ids []string, opened bool) error
	DeleteExpiredAssignedCardPacks(expiryDate time.Time) (int64, error)
	CreateAssignedPlayer(ap *schema.AssignedPlayer) (*schema.AssignedPlayer, error)
	GetAssignedPlayerByID(id string) (*schema.AssignedPlayer, error)
	GetAllAssignedPlayers() (schema.AssignedPlayerSlice, error)
	UpdateAssignedPlayer(ap *schema.AssignedPlayer) error
	DeleteAssignedPlayer(id string) error
	GetAssignedPlayersByUserID(userID string) (schema.AssignedPlayerSlice, error)
	GetAssignedPlayersByPlayerID(playerID string) (schema.AssignedPlayerSlice, error)
	PlayerPopularity() (map[string]int, error)
	UserEngagement() (map[string]float64, error)
	TeamBalanceAnalysis() (map[string]map[string]int, error)
	CreateCardPack(cp *schema.CardPackType) error
	GetCardPackByID(id string) (*schema.CardPackType, error)
	GetAllCardPacks() ([]*schema.CardPackType, error)
	UpdateCardPack(cp *schema.CardPackType) error
	DeleteCardPack(id string) error
	GetCardPackTypeByName(name string) (*schema.CardPackType, error)
	GetCardPackTypeByCode(code string) (*schema.CardPackType, error)
	GetAssignedCardPacksByType(cardPackTypeID string) ([]*schema.AssignedCardPack, error)
	CreatePlayerBucket(pb *schema.PlayerBucket) (*schema.PlayerBucket, error)
	GetPlayerBucketByID(id string) (*schema.PlayerBucket, error)
	GetAllPlayerBuckets() (schema.PlayerBucketSlice, error)
	UpdatePlayerBucket(pb *schema.PlayerBucket) error
	DeletePlayerBucket(id string) error
	GetPlayerBucketsByTeamID(teamID string) (schema.PlayerBucketSlice, error)
	AssignPlayerToTeam(playerID, teamID string) (*schema.PlayerBucket, error)
	CountPlayerBuckets() (int64, error)
	CountPlayersByTeam(ctx context.Context) (map[string]int, error)
	GetPlayersByStatus(status string) (schema.PlayerBucketSlice, error)
	GetPlayersByTeamAndStatus(teamID, status string) (schema.PlayerBucketSlice, error)
	CalculateRarityPercentages() (*RarityPercentages, error)
	GetPlayersByRarity(rarity string) (schema.PlayerBucketSlice, error)
	GetUniqueRandomTeams(count int) (schema.TeamSlice, error)
	GetRandomPlayerByTeamAndRarity(teamID string, rarity string) (*schema.PlayerBucket, error)
	GetUniqueRandomTeamExcluding(excluded map[string]struct{}) (*schema.Team, error)
	GetItemByID(id string) (*schema.Item, error)
	GetItems() (schema.ItemSlice, error)

	CreateOrder(order *schema.Order) (*schema.Order, error)
	UpdateOrder(order *schema.Order) error
	GetOrderByID(id string) (*schema.Order, error)
	LockOrder(orderID string) (*schema.Order, error)
	UpdateTeam(team *schema.Team) (int64, error)
	TeamExistsByID(teamID string) (bool, error)
	TeamExistsByName(teamName string) (bool, *schema.Team, error)
	ValidateTeam(team *schema.Team) error
	AddTeam(team *schema.Team) error
	GetOrCreateTeamByName(teamName string) (*schema.Team, error)
	GetTeamByName(teamName string) (*schema.Team, error)
	GetCountryByID(countryID string) (*schema.Country, error)
	GetAllCountries() (schema.CountrySlice, error)
	AddCountry(country *schema.Country) error
	UpdateCountry(country *schema.Country) (int64, error)
	DeleteCountry(countryID string) (int64, error)
	CountryExistsByID(countryID string) (bool, error)
	GetCountryByName(name string) (*schema.Country, error)
	GetRandomPlayerBucketByTeamID(teamID string) (*schema.PlayerBucket, error)
	CreateUserWithPassword(user *schema.User, password string) (*schema.User, error)
	SendVerificationEmail(user *schema.User) error
	VerifyEmail(token string) error
	AuthenticateUser(email, password string) (*schema.User, error)
	ResendVerificationEmail(userID string) error
	UpdatePassword(userID, oldPassword, newPassword string) error
	GetActionNameByActionID(actionID int) (string, error)
	CreateChatMessage(message *schema.ChatMessage) error
	GetChatRoomByName(name string) (*schema.ChatRoom, error)
	GetChatRoomByMatchID(matchID string) (*schema.ChatRoom, error)
	GetChatMessagesByRoomID(roomID string, offset, limit, minutesPrior int) (*[]types.ChatMessage, error)
	GetChatRoomMessageByID(roomID, messageID string) (*schema.ChatMessage, error)
	DeleteChatRoomMessageByID(roomID, messageID string) (bool, error)
	GetChatRoomMember(roomID, userID string) (*schema.ChatRoomMember, error)
	GetMatchEventsByMatchID(matchID string) (schema.MatchEventSlice, error)
	GetMatchEventsByMatchIDAndType(matchID string, eventType int) (schema.MatchEventSlice, error)
	UpdateUser(*schema.User) error
	SavePushToken(userID, token string) error
	RemoveFinishedGameByUserID(userID string, gameID string) error
	EmailExists(email string) (bool, error)
	UsernameExists(username string) (bool, error)
	DeleteUser(userID string) error
	DeleteAllGamesByUserID(userID string) error
	GetUnnotifiedGamesByUserID(userID string) (schema.GameSlice, error)
	SetGameNotified(gameID string) error
	SetAllGamesNotified() error
	GeneratePasswordResetToken(email string) (string, error)
	VerifyPasswordResetToken(token string) (*schema.User, error)
	UpdatePasswordWithToken(token, newPassword string) error
	AuthenticateUserWithSession(sequenceSessionID string) (*schema.User, error)
	GetSports() (schema.SportSlice, error)
	GetSportByID(id string) (*schema.Sport, error)
	GetStoreProducts() (schema.StoreProductSlice, error)
	GetStoreProductByProductID(productID string) (*schema.StoreProduct, error)
	GetStoreProductByStoreAndProductID(store, productID string) (*schema.StoreProduct, error)
	CreateStoreProductTransaction(
		externalTransactionID, originStore, storeProductID string,
		purchaseDate time.Time,
	) (string, error)
	CancelStoreProductTransaction(externalTransactionID string) error
	GetStoreProductTransactionByExternalID(externalTransactionID string) (*schema.StoreProductTransaction, error)
	ConfirmStoreProductTransactionByExternalID(userID, transactionID, externalTransactionID string) error

	CreateNFTBucket(nb *schema.NFTBucket) (*schema.NFTBucket, error)
	GetNFTBucketByID(id string) (*schema.NFTBucket, error)
	GetAllNFTBuckets() (schema.NFTBucketSlice, error)
	UpdateNFTBucket(nb *schema.NFTBucket) error
	DeleteNFTBucket(id string) error
	GetNFTBucketsByTeamID(teamID string) (schema.NFTBucketSlice, error)
	CountNFTBuckets() (int64, error)
	GetNFTsByRarity(rarity string) (schema.NFTBucketSlice, error)
	GetNFTsByTeamAndRarity(teamID, rarity string) (schema.NFTBucketSlice, error)
	GetRandomNFTByTeamAndRarity(teamID string, rarity string) (*schema.NFTBucket, error)
	GetRandomNFTBucketByTeamID(teamID string) (*schema.NFTBucket, error)
	GetNFTBucketsByPosition(position string) (schema.NFTBucketSlice, error)
	GetNFTBucketsByGamePosition(gamePosition string) (schema.NFTBucketSlice, error)
	GetNFTBucketsByPositionAndTeam(teamID, position string) (schema.NFTBucketSlice, error)
	GetNFTBucketsByTeamAndNationality(teamID, nationality string) (schema.NFTBucketSlice, error)
	GetNFTBucketsByNationality(nationality string) (schema.NFTBucketSlice, error)
	GetNFTBucketsByAgeRange(minAge, maxAge int) (schema.NFTBucketSlice, error)
	GetNFTBucketsByRarityAndPosition(rarity, position string) (schema.NFTBucketSlice, error)
	CheckLimit(nftID string, rarity string) (bool, error)
	CountNFTBucketsByRarity() (map[string]int64, error)
	AssignNFTToTeam(nftID, teamID string) (*schema.NFTBucket, error)
	GetRandomNFTBucket() (*schema.NFTBucket, error)
	GetNFTBucketByName(name string) (*schema.NFTBucket, error)
	GetRandomNFTBucketByTeamAndRarityAndStarRating(teamID string, rarity string, starRating int, cardpackCollection string) (*schema.NFTBucket, error)
	CountAllEntries() (int64, error)
	SetAssignedCardPackRevealed(id string) error
	SetAssignedCardPackOpened(id string) error
	UpdateAssignedCardPackCardIds(id string, cardIds []string) error
	GetAssignedPlayersWithNFTDetails(uuids []string) (
		[]struct {
			Player *schema.AssignedPlayer
			NFT    *schema.NFTBucket
		}, error,
	)
	GetAssignedPlayersByNFTIDs(nftIDs []string) ([]PlayerWithNFTDetails, error)

	GetGameWeeks(limit int) (schema.GameWeekSlice, error)
	GetGameWeek(id string) (*schema.GameWeek, error)
	GetCurrentGameWeek() (*schema.GameWeek, error)

	GetCurrentWeekLeaderboard(
		seasonStartAt, seasonEndAt, weekStartAt, weekEndAt time.Time,
		userID string,
	) ([]*types.LeaderboardEntry, error)
	GetConcludedWeekLeaderboard(weekID, userID string) ([]*types.LeaderboardEntry, error)

	GetLatestUserGameWeekHistory(userID string) (*schema.UserGameWeekHistory, error)
	ListUserGameWeekHistories(userID string) (schema.UserGameWeekHistorySlice, error)
	GetUserGameWeekHistory(userID string, ID string) (*schema.UserGameWeekHistory, error)

	GetGameWeekDivision(weekID, divisionID string) (*schema.GameWeekDivision, error)
	GetDivisionRewards(weekID string) (schema.DivisionRewardSlice, error)
	GetDivisionByID(divisionID string) (*schema.Division, error)

	DeductFromPackLimit(cardPackTypeID string, amount int) error
	RestockPackLimit(cardPackTypeID string, amount int) error
	DeductFromRarityLimit(nftID string, rarity string, amount int) error

	GetUserByWalletAddress(walletAddress string) (bool, error)
	GetUsernameByWalletAddress(walletAddress string) (string, error)

	GetAppInboxByUserID(userID string) (schema.AppInboxSlice, error)
	GetAppInboxByID(id string, userID string) (*schema.AppInbox, error)
	CreateAppInbox(inbox *schema.AppInbox) (*schema.AppInbox, error)
	MarkAppInboxAsRead(id string, userID string) error
	DeleteAppInboxByID(id string, userID string) error
	GetUnreadAppInboxCountByUserID(userID string) (int64, error)
	DeleteAllAppInboxesByUserID(userID string) error
	GetRecentAppInboxByUserID(userID string, limit int) (schema.AppInboxSlice, error)
	SetAppInboxesAsReadByUserID(userID string) error
	MarkAppInboxAsClaimed(id string, userID string) error
	MarkAppInboxAsUnclaimed(id string, userID string) error
	GetClaimedAppInboxByUserID(userID string) (schema.AppInboxSlice, error)
	GetUnclaimedAppInboxByUserID(userID string) (schema.AppInboxSlice, error)
	SetAppInboxesAsClaimedByUserID(userID string) error
	DeleteAllClaimedAppInboxesByUserID(userID string) error
	GetRewardByID(rewardID string) (*schema.Reward, error)
	CreateReward(reward *schema.Reward) error

	GetPackLimitsByCardPackName(name string) (int, error)
	GetPackLimitsByCardPackCode(code string) (int, error)

	GetNFTBucketByIDandRarity(id string, rarity string) (*schema.NFTBucket, error)
	GetNFTBucketByIDandRarityFiltered(id string, rarity string) (*types.GenericRarityNFT, error)

	GetBadges() (schema.BadgeSlice, error)
	GetBadgeNameByID(badgeID int) (string, error)
	GetBadgeByID(badgeID string) (*schema.Badge, error)
	DeleteBadge(badgeID string) error
	UpdateBadge(badgeID string, options *model.UpdateBadgeRequest) (*schema.Badge, error)
	CreateBadges(badges *model.CreateBadgesRequest) ([]schema.Badge, error)
	GetTypeDefaultBadges() (schema.BadgeSlice, error)

	GetBanners() (schema.BannerSlice, error)
	GetBannerNameByID(bannerID int) (string, error)
	GetBannerByID(bannerID string) (*schema.Banner, error)
	DeleteBanner(bannerID string) error
	UpdateBanner(bannerID string, options *model.UpdateBannerRequest) (*schema.Banner, error)
	CreateBanners(banners *model.CreateBannersRequest) ([]schema.Banner, error)
	GetTypeDefaultBanners() (schema.BannerSlice, error)

	GetFrames() (schema.FrameSlice, error)
	GetFrameNameByID(frameID int) (string, error)
	GetFrameByID(frameID string) (*schema.Frame, error)
	DeleteFrame(frameID string) error
	UpdateFrame(frameID string, options *model.UpdateFrameRequest) (*schema.Frame, error)
	CreateFrames(frames *model.CreateFramesRequest) ([]schema.Frame, error)
	GetDefaultFrames() (schema.FrameSlice, error)

	GetUserBadges(userID string) (schema.UserBadgeSlice, error)
	GetUserBadgeByID(userBadgeID int) (*schema.UserBadge, error)
	DeleteUserBadge(userBadgeID int) error
	CreateUserBadge(userID string, badgeID int) (*schema.UserBadge, error)
	ChangeAllUsersBadgesToUnselected() error
	ChangeUserBadgeToSelected(userBadgeID int) error

	GetUserBanners(userID string) (schema.UserBannerSlice, error)
	GetUserBannerByID(userBannerID int) (*schema.UserBanner, error)
	DeleteUserBanner(userBannerID int) error
	CreateUserBanner(userID string, bannerID int) (*schema.UserBanner, error)
	ChangeAllUsersBannersToUnselected() error
	SelectUserBanner(userBannerID int) error

	GetUserFrames(userID string) (schema.UserFrameSlice, error)
	GetUserFrameByID(userFrameID int) (*schema.UserFrame, error)
	DeleteUserFrame(userFrameID int) error
	CreateUserFrame(userID string, frameID int) (*schema.UserFrame, error)
	ChangeAllUsersFramesToUnselected() error
	ChangeUserFrameToSelected(userFrameID int) error
}

type ActionStore interface {
	GetAction(id int) (*schema.Action, error)
}
