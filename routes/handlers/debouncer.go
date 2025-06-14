package handlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

var requestTimestamps = make(map[string]time.Time)
var mu sync.Mutex

// DebounceMiddleware ensures that the handler can only be called once every 5 seconds per user.
func (e *Env) DebounceMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := userID(c)
		if userID == "" {
			userID = c.RealIP()
		}
		mu.Lock()
		lastRequestTime, found := requestTimestamps[userID]
		if found && time.Since(lastRequestTime) < 3*time.Second {
			mu.Unlock()
			return echo.NewHTTPError(http.StatusTooManyRequests, "Please wait before making another request.")
		}
		requestTimestamps[userID] = time.Now()
		mu.Unlock()
		return next(c)
	}
}
