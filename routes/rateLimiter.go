//path routes/rateLimiter.go

package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
)

func rateLimitMiddleware(next http.Handler) http.Handler {
	limiter := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Hour})
	limiter.SetMessage("Too many requests")
	limiter.SetMessageContentType("application/json; charset=utf-8")
	limiter.SetOnLimitReached(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Limit reached")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("You have reached maximum request limit"))
	})
	return tollbooth.LimitHandler(limiter, next)
}
