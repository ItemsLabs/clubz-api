package routes

import (
	"log"
	"os"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/config"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/types"

	"github.com/gameon-app-inc/laliga-matchfantasy-api/database"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/action"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/database/dbstore"

	_ "github.com/gameon-app-inc/laliga-matchfantasy-api/docs"
	"github.com/gameon-app-inc/laliga-matchfantasy-api/routes/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRouter(r *echo.Echo, rawStore *dbstore.DBStore, wsConfigs map[string]types.WebsocketConfig) { //, m *melody.Melody, bus event_bus.Bus, chatWS *melody.Melody) {
	store := database.Store(rawStore)

	// Middleware setup
	r.Pre(middleware.RemoveTrailingSlash()) // Trailing slash
	r.HTTPErrorHandler = handlers.ErrorHandler

	r.GET("/swagger/*", echoSwagger.WrapHandler)
	ps := handlers.NewPlayerService(rawStore)
	redisClient := config.NewRedisClient()
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	env := &handlers.Env{
		Store:            store,
		ActionStore:      action.NewCachedActionStore(store),
		RedisClient:      redisClient,
		WebsocketConfigs: wsConfigs,
		// Melody:      m,
		// ChatWS:      chatWS,
	}

	env.InitEventPipe()

	// Base API route
	api := r.Group("/api", middleware.Logger()) // API

	api.GET("/chat/ws", env.ChatWebsocketHandler, env.DebounceMiddleware) // Chat websocket

	api.GET("/leaderboards/ws", env.LeaderboardWebsocketHandler, env.DebounceMiddleware) // Leaderboard websocket

	// JWT Middleware setup
	jwtSecretKey := os.Getenv("JWT_SECRET")
	if jwtSecretKey == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}
	protected := api.Group("", env.JWTMiddleware(jwtSecretKey))
	{
		protected.GET("/users/:id", env.GetUserByID)                   // User profile
		protected.POST("/users/:id/follow", env.FollowUser)            // Follow user
		protected.POST("/users/:id/unfollow", env.UnFollowUser)        // Unfollow user
		protected.GET("/users/search", env.SearchUsers)                // Search users
		protected.GET("/users/profile", env.CurrentProfile)            // Current user profile
		protected.POST("/users/profile/name", env.UpdateUserName)      // Update user name
		protected.POST("/users/profile/referral", env.UseReferralCode) // Use referral code
		// protected.POST("/users/profile/paypal_email", env.UpdatePayPalEmail) // Update paypal email
		protected.GET("/users/profile/followers", env.GetFollowers)                            // Get followers
		protected.GET("/users/profile/followings", env.GetFollowings)                          // Get followings
		protected.GET("/users/profile/wallet/balance/:contractType", env.GetUserWalletBalance) // Get user wallet balance
		// protected.PUT("/update-password", env.UpdatePassword, env.DebounceMiddleware) // Update password
		protected.POST("/remove-finished-game/:id", env.SetGameNotified, env.DebounceMiddleware)      // Remove finished game
		protected.POST("/remove-all-finished-games", env.SetAllGamesNotified, env.DebounceMiddleware) // Remove all finished games
		protected.GET("/users/card-packs", env.GetAssignedCardPacksByUserIDHandler)                   // Get assigned card packs
		protected.GET("/users/players", env.GetMintedPlayersFromCardPacksHandler)
		protected.GET("/users/players-filtered", env.GetMintedPlayersFromCardPacksFilteredHandler)
		protected.POST("/users/change-email", env.ChangeEmail)     // Get minted players
		protected.GET("/users/badges", env.GetUserBadges)          // Get user badges
		protected.GET("/users/banners", env.GetUserBanners)        // Get user banners
		protected.GET("/users/frames", env.GetUserFrames)          // Get user frames
		protected.POST("/users/badges/:id", env.SelectUserBadge)   // Change user badge to selected
		protected.POST("/users/banners/:id", env.SelectUserBanner) // Change user banner to selected
		protected.POST("/users/frames/:id", env.SelectUserFrame)   // Change user frame to selected

	}

	// IAP section
	iap := api.Group("/iap") // General IAP section
	{
		iap.GET("/card-pack-types", env.GetCardPackTypes)
	}
	rc := iap.Group("/rc") // RevenueCat section
	{
		rc.POST("/create-transaction", env.CreateRevenueCatPurchase)
		rc.POST("/cancel-transaction", env.CancelRevenueCatPurchase)
	}

	// Auth section
	auth := api.Group("/auth")
	{
		// auth.POST("/register", env.Register, env.DebounceMiddleware)
		// auth.POST("/login", env.Login, env.DebounceMiddleware)
		// auth.POST("/resend-verification", env.ResendVerificationEmail, env.DebounceMiddleware)
		// auth.GET("/verify-email", env.VerifyEmail, env.DebounceMiddleware)
		auth.POST("/register-guest", env.RegisterGuest, env.DebounceMiddleware)                                // Register guest
		auth.POST("/refresh-token", env.RefreshToken, env.DebounceMiddleware, env.JWTMiddleware(jwtSecretKey)) // Refresh token
		// auth.POST("/upgrade-to-premium", env.UpgradeToPremium, env.JWTMiddleware(jwtSecretKey)) // Upgrade to premium
		auth.POST("/push-token", env.RegisterPushToken, env.JWTMiddleware(jwtSecretKey))                         // Register Expo Push token
		auth.POST("/delete-user", env.DeleteUser, env.JWTMiddleware(jwtSecretKey))                               // Delete user
		auth.POST("/login-sequence", env.LoginSequence, env.DebounceMiddleware)                                  // Login sequence
		auth.POST("/upgrade-to-premium-sequence", env.UpgradeToPremiumSequence, env.JWTMiddleware(jwtSecretKey)) // Upgrade to premium sequence
		auth.POST("/upgrade-or-login", env.UpgradeOrLogin, env.JWTMiddleware(jwtSecretKey))                      // Logout
	}
	chat := protected.Group("/chat")
	{
		chat.GET("/:room", env.GetChatMessages, env.DebounceMiddleware)                     // Get chat messages
		chat.POST("/:room", env.PostChatMessage, env.DebounceMiddleware)                    // Post chat message
		chat.DELETE("/:room/:messageId", env.DeleteChatRoomMessage, env.DebounceMiddleware) // Delete chat message
	}
	ws := api.Group("/ws")
	{
		ws.GET("", env.EventPipe) // Websocket
	}
	matches := api.Group("/matches")
	{
		matches.GET("/upcoming", env.UpcomingMatches)                                // Upcoming matches
		matches.GET("/:id", env.GetMatchInfo)                                        // Match info
		protectedChat := matches.Group("/:id/chat", env.JWTMiddleware(jwtSecretKey)) // Match chat
		{
			protectedChat.GET("", env.GetMatchChatMessages)                                             // Get chat messages
			protectedChat.POST("", env.PostMatchChatMessages, env.DebounceMiddleware)                   // Post chat message
			protectedChat.DELETE("/:messageId", env.DeleteMatchChatRoomMessage, env.DebounceMiddleware) // Delete chat message
		}
		matches.GET("/:id/squad", env.GetMatchSquad)                                                                          // Match squad
		matches.GET("/:id/headlines/lobby", env.GetLobbyHeadlines)                                                            // Lobby headlines
		matches.GET("/:id/headlines/gameplay", env.GetGamePlayHeadlines)                                                      // Match headlines
		matches.GET("/:id/headlines/fulltime", env.GetFullTimeHeadlines)                                                      // Fulltime headlines
		matches.GET("/:matchID/leaderboard/:count", env.GetLeaderboard)                                                       // Match leaderboard
		matches.GET("/:id/following_leaderboard", env.FollowingLeaderboard, env.JWTMiddleware(jwtSecretKey))                  // Following leaderboard
		matches.GET("/:id/user_games/:user_id", env.GetUserGame, env.JWTMiddleware(jwtSecretKey))                             // User game
		matches.GET("/:match_id/players/:player_id/pregame_stats", ps.GetPregamePlayerStats, env.JWTMiddleware(jwtSecretKey)) // Pregame stats
		matches.GET("/:match_id/players/:player_id/live_stats", ps.GetLivePlayerStats, env.JWTMiddleware(jwtSecretKey))       // Live stats
		events := matches.Group("/:id/events")                                                                                // Match events
		{
			events.GET("", env.GetMatchEvents)            // Get match events
			events.GET("/type", env.GetMatchEventsByType) // Get match events by type
		}
	}
	games := protected.Group("/games")
	{
		games.GET("", env.GamesList)      // Game List
		games.POST("/join", env.JoinGame) // Join game
		games.POST("/new-join", env.NewJoinGame)
		games.GET("/:id", env.GetGameByID)                                        // Get game by ID
		games.POST("/:id/swap", env.SwapPlayer, env.DebounceMiddleware)           // Swap player
		games.POST("/:id/substitution", env.Substitution, env.DebounceMiddleware) // Substitution
		games.GET("/:id/actions", env.GetGameEvents)                              // Get game actions
		games.GET("/:id/leaderboard_position", env.GetLeaderBoardPosition)        // Get leaderboard position
		games.GET("/history", env.GetGameHistory)                                 // Get game history
		games.GET("/:id/powerups", env.GetGamePowerUps)                           // Get game powerups
		games.POST("/:id/powerups", env.ApplyPowerUp)                             // Apply powerup
		games.GET("/finished", env.GetUnnotifiedGames)                            // Get unnotified games
	}
	web3 := protected.Group("/web3", env.DebounceMiddleware)
	{
		// web3.POST("/create-card-pack", env.CreateCardPack)
		// web3.POST("/assign-card-pack", env.AssignCardPackToPlayerHandler)
		web3.GET("/balance/:address/:contractType", env.GetBalanceOf)
		// web3.POST("/approve", env.ApproveToken)
		// web3.POST("/setApprovalForAll", env.SetApprovalForAll)
		// web3.POST("/transfer", env.TransferToken)
		// web3.POST("/burn", env.BurnToken)
		web3.GET("/tokenURI/:tokenId/:contractType", env.GetTokenURI)
		web3.GET("/ownerOf/:tokenId/:contractType", env.OwnerOfToken)
		// web3.GET("/isApprovedForAll/:ownerAddress/:operatorAddress/:contractType", env.IsApprovedForAll)
		web3.GET("/getOwner/:contractType", env.GetOwner)
		// web3.GET("/getApproved/:tokenId/:contractType", env.GetApproved)
		// web3.POST("/safeTransferFrom", env.SafeTransferFrom)
		// web3.POST("/sendSignedTransaction", env.SendSignedTransaction)
		web3.POST("/open-card-pack/:cardPackID", env.OpenCardPack)
		web3.POST("/new-open-card-pack/:cardPackID", env.NewOpenCardPack)
		web3.POST("/open-card-pack-collection/:cardPackID", env.OpenCollection)
		web3.GET("/get-card-packs-players", env.GetCardPackWithNFTDetailsHandler)
		// web3.POST("/import-data", env.ImportNFTsHandler)
	}
	api.GET("/teams/:id", env.GetTeamByID)     // Get team by ID
	api.GET("/players/:id", env.GetPlayerByID) // Get player by ID
	actions := api.Group("/actions")           // Actions
	{
		actions.GET("", env.ActionsList) // Actions list
		// actions.POST("", env.CreateActions)	// Create action
		// actions.POST("/:id", env.UpdateAction)	// Update action
		// actions.DELETE("/:id", env.DeleteAction)	// Delete action
	}
	sports := api.Group("/sports") // Sports
	{
		sports.GET("", env.SportList)                // Sports list
		sports.GET("/:id", env.GetSportByID)         // Get sport by ID
		sports.GET("/:id/powerups", env.PowerUpList) // Powerups list
	}
	inbox := protected.Group("/inbox")
	{
		inbox.GET("/internal", env.GetAppInboxByUserID)                              // Get inbox messages
		inbox.POST("/internal/:id", env.SetMarkAppInboxAsRead)                       // Post inbox message
		inbox.POST("/internal/read", env.SetAllAppInboxesAsReadByUserID)             // Mark all inbox messages as read
		inbox.POST("/internal/claim/:id", env.ClaimInboxItem)                        // Claim inbox item
		inbox.POST("/internal/claim-prizes/:id", env.UpdateUserDetailsAndClaimInbox) // Claim inbox item
	}
	// api.GET("/settings", env.Settings)
	// api.GET("/items/:id", env.GetItemByID)
	// api.GET("/items", env.ItemList)

	purchase := protected.Group("/purchase") // Purchase
	{
		purchase.POST("", env.Purchase)        // Make
		purchase.GET("/:id", env.GetOrderByID) // Get
	}
	api.POST("/webhook", env.WebhooksPurchase) // Webhook

	r.GET("/readiness", env.Readiness) // Readiness
	r.GET("/healthz", env.Health)      // Health
	api.GET("/user-with-wallet/:walletAddress", env.GetUserByWalletAddress)
	api.GET("/username-with-wallet/:walletAddress", env.GetUserByUsernameByWalletAddress)
	api.GET("/cardpacks-by-code/:code", env.GetPackLimitsByCardPackCodeHandler)
	weeks := api.Group("/weeks") // Weeks
	{
		weeks.GET("", env.GetWeekList)                        // Week list
		weeks.GET("/:id/leaderboard", env.GetWeekLeaderboard) // Week leaderboard
	}
	protected.GET("/leaderboard/current_week", env.GetCurrentWeekLeaderboard) // Current week leaderboard

	badges := api.Group("/badges") // Badges
	{
		badges.GET("", env.BadgesList) // List
		badges.GET("/:id", env.GetBadge)
	}
	banners := api.Group("/banners") // Banners
	{
		banners.GET("", env.BannersList) // List
		banners.GET("/:id", env.GetBanner)
	}
	frames := api.Group("/frames") // Frames
	{
		frames.GET("", env.FramesList) // List
		frames.GET("/:id", env.GetFrame)
	}
	// UserGameWeekHistories routes
	userGameWeekHistories := protected.Group("/user_game_week_histories") // UserGameWeekHistories
	{
		userGameWeekHistories.GET("", env.ListUserGameWeekHistories)           // List
		userGameWeekHistories.GET("/:id", env.GetUserGameWeekHistory)          // Get
		userGameWeekHistories.GET("/latest", env.GetLatestUserGameWeekHistory) // Get latest
	}

}
