package handlers

import (
	"log"
	"net/http"

	"github.com/itemslabs/clubz-api/database/schema"
	"github.com/itemslabs/clubz-api/routes/apiconv"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
)

// GetAppInboxByUserID godoc
// @Summary Get user inbox messages
// @Description Retrieve all inbox messages for a specific user
// @Tags inbox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} []model.AppInbox
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Router /inbox [get]
func (e *Env) GetAppInboxByUserID(c echo.Context) error {
	userID := userID(c)

	inboxes, err := e.Store.GetAppInboxByUserID(userID)
	if err != nil {
		log.Printf("Error retrieving inbox messages for user ID %s: %v", userID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve inbox messages")
	}

	return e.RespondSuccess(c, apiconv.ToAppInboxSlice(inboxes))
}

// SetAppInboxesAsReadByUserID godoc
// @Summary Mark all inbox messages as read
// @Description Mark all inbox messages as read for a specific user
// @Tags inbox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Router /inbox/read [post]
func (e *Env) SetMarkAppInboxAsRead(c echo.Context) error {
	userID := userID(c)
	id := c.Param("id")
	err := e.Store.MarkAppInboxAsRead(id, userID)
	if err != nil {
		log.Printf("Error setting inbox messages as read for user ID %s: %v", userID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to set inbox messages as read")
	}

	return e.RespondSuccess(c, "ok")
}

// SetAllAppInboxesAsReadByUserID godoc
// @Summary Mark all inbox messages as read
// @Description Mark all inbox messages as read for a specific user
// @Tags inbox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Router /inbox/read/all [post]
func (e *Env) SetAllAppInboxesAsReadByUserID(c echo.Context) error {
	userID := userID(c)
	err := e.Store.SetAppInboxesAsReadByUserID(userID)
	if err != nil {
		log.Printf("Error setting all inbox messages as read for user ID %s: %v", userID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to set all inbox messages as read")
	}

	return e.RespondSuccess(c, "ok")
}

// ClaimInboxItem godoc
// @Summary Claim inbox item
// @Description Claim an inbox item for a specific user
// @Tags inbox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Inbox item ID"
// @Success 200 {object} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Router /inbox/claim/{id} [post]
func (e *Env) ClaimInboxItem(c echo.Context) error {
	userID := userID(c)
	inboxID := c.Param("id")

	if inboxID == "" {
		log.Printf("Inbox item ID is missing")
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	// Mark the inbox item as claimed
	err := e.Store.MarkAppInboxAsClaimed(inboxID, userID)
	if err != nil {
		log.Printf("Error claiming inbox item ID %s for user ID %s: %v", inboxID, userID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to claim inbox item")
	}

	// Retrieve the claimed inbox item
	inboxItem, err := e.Store.GetAppInboxByID(inboxID, userID)
	if err != nil {
		log.Printf("Error retrieving claimed inbox item ID %s for user ID %s: %v", inboxID, userID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve claimed inbox item")
	}

	// Handle the reward based on the reward type
	if inboxItem.RewardID.Valid {
		reward, err := e.Store.GetRewardByID(inboxItem.RewardID.String)
		if err != nil {
			log.Printf("Error retrieving reward ID %s for inbox item ID %s: %v", inboxItem.RewardID.String, inboxID, err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve reward details")
		}

		// Determine the reward type and create the corresponding transaction
		text := "Reward for claiming inbox item"
		if reward.Credits > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Amount:     reward.Credits,
				Text:       null.StringFrom(text),
				ObjectType: "v", // Virtual currency
				Delivered:  true,
			})
			if err != nil {
				log.Printf("Error creating credit transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create credit transaction")
			}
		}

		if reward.GameToken > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Amount:     reward.GameToken,
				Text:       null.StringFrom(text),
				ObjectType: "g", // Game token
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating game token transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create game token transaction")
			}
		}

		if reward.LaptToken > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Amount:     reward.LaptToken,
				Text:       null.StringFrom(text),
				ObjectType: "l", // LAPT token
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating LAPT token transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create LAPT token transaction")
			}
		}

		if reward.EventTickets > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Quantity:   int64(reward.EventTickets),
				Text:       null.StringFrom(text),
				ObjectType: "e", // Event ticket
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating event ticket transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create event ticket transaction")
			}
		}

		if reward.Ball > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Quantity:   int64(reward.Ball),
				Text:       null.StringFrom(text),
				ObjectType: "b", // Balls
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating ball transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create ball transaction")
			}
		}

		if reward.SignedBall > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Quantity:   int64(reward.SignedBall),
				Text:       null.StringFrom(text),
				ObjectType: "a", // Signed Balls
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating signed ball transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create signed ball transaction")
			}
		}

		if reward.Shirt > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Quantity:   int64(reward.Shirt),
				Text:       null.StringFrom(text),
				ObjectType: "s", // Shirts
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating shirt transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create shirt transaction")
			}
		}

		if reward.SignedShirt > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Quantity:   int64(reward.SignedShirt),
				Text:       null.StringFrom(text),
				ObjectType: "h", // Signed Shirts
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating signed shirt transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create signed shirt transaction")
			}
		}
	}

	return e.RespondSuccess(c, "ok")
}

// UpdateUserDetailsAndClaimInbox godoc
// @Summary Update user real name and email, and mark inbox item as claimed
// @Description Update the real name and email of the user if provided, and mark the inbox item as claimed
// @Tags inbox
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Inbox item ID"
// @Param real_name body string true "New real name"
// @Param email body string false "New email (optional)"
// @Success 200 {object} string "ok"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 403 {string} string "forbidden"
// @Router /inbox/update-user-and-claim/{id} [post]
func (e *Env) UpdateUserDetailsAndClaimInbox(c echo.Context) error {
	userID := userID(c)
	inboxID := c.Param("id")

	if inboxID == "" {
		log.Printf("Inbox item ID is missing")
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	// Bind request body to get the real_name and email
	request := struct {
		RealName string `json:"real_name"`
		Email    string `json:"email,omitempty"`
	}{}

	if err := c.Bind(&request); err != nil {
		log.Printf("Failed to bind request data: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	// Fetch the user by ID
	user, err := e.Store.GetUserByID(userID)
	if err != nil {
		log.Printf("Failed to retrieve user ID %s: %v", userID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user")
	}

	// Update the real_name of the user
	user.RealName = null.StringFrom(request.RealName)

	// Optionally update the email if provided
	if request.Email != "" {
		if exists, err := e.Store.EmailExists(request.Email); err != nil || exists {
			log.Printf("Email already exists or error checking email existence: %v", err)
			return echo.NewHTTPError(http.StatusBadRequest, "email already in use")
		}
		user.Email = null.StringFrom(request.Email)
	}

	// Update user details in the database
	if err := e.Store.UpdateUser(user); err != nil {
		log.Printf("Failed to update user details: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user details")
	}

	// Mark the inbox item as claimed
	err = e.Store.MarkAppInboxAsClaimed(inboxID, userID)
	if err != nil {
		log.Printf("Error claiming inbox item ID %s for user ID %s: %v", inboxID, userID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to claim inbox item")
	}

	// Retrieve the claimed inbox item
	inboxItem, err := e.Store.GetAppInboxByID(inboxID, userID)
	if err != nil {
		log.Printf("Error retrieving claimed inbox item ID %s for user ID %s: %v", inboxID, userID, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve claimed inbox item")
	}

	// Handle the reward based on the reward type
	if inboxItem.RewardID.Valid {
		reward, err := e.Store.GetRewardByID(inboxItem.RewardID.String)
		if err != nil {
			log.Printf("Error retrieving reward ID %s for inbox item ID %s: %v", inboxItem.RewardID.String, inboxID, err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve reward details")
		}

		// Determine the reward type and create the corresponding transaction
		text := "Reward for claiming inbox item"
		if reward.Credits > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Amount:     reward.Credits,
				Text:       null.StringFrom(text),
				ObjectType: "v", // Virtual currency
				Delivered:  true,
			})
			if err != nil {
				log.Printf("Error creating credit transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create credit transaction")
			}
		}

		if reward.GameToken > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Amount:     reward.GameToken,
				Text:       null.StringFrom(text),
				ObjectType: "g", // Game token
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating game token transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create game token transaction")
			}
		}

		if reward.LaptToken > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Amount:     reward.LaptToken,
				Text:       null.StringFrom(text),
				ObjectType: "l", // LAPT token
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating LAPT token transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create LAPT token transaction")
			}
		}

		if reward.EventTickets > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Quantity:   int64(reward.EventTickets),
				Text:       null.StringFrom(text),
				ObjectType: "e", // Event ticket
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating event ticket transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create event ticket transaction")
			}
		}

		if reward.Ball > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Quantity:   int64(reward.Ball),
				Text:       null.StringFrom(text),
				ObjectType: "b", // Balls
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating ball transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create ball transaction")
			}
		}

		if reward.SignedBall > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Quantity:   int64(reward.SignedBall),
				Text:       null.StringFrom(text),
				ObjectType: "a", // Signed Balls
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating signed ball transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create signed ball transaction")
			}
		}

		if reward.Shirt > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Quantity:   int64(reward.Shirt),
				Text:       null.StringFrom(text),
				ObjectType: "s", // Shirts
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating shirt transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create shirt transaction")
			}
		}

		if reward.SignedShirt > 0 {
			_, err := e.Store.CreateTransaction(&schema.Transaction{
				ID:         uuid.New().String(),
				UserID:     userID,
				Quantity:   int64(reward.SignedShirt),
				Text:       null.StringFrom(text),
				ObjectType: "h", // Signed Shirts
				Delivered:  false,
			})
			if err != nil {
				log.Printf("Error creating signed shirt transaction for user ID %s: %v", userID, err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create signed shirt transaction")
			}
		}
	}
	return e.RespondSuccess(c, "User details updated and inbox item claimed successfully")
}
