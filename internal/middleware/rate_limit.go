package middleware

import (
	"net/http"
	"sync"
	"task-api/internal/utils"
	"time"

	"golang.org/x/time/rate"
)

var (
	visitors = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	limiter, exists := visitors[ip]

	if !exists {
		limiter = rate.NewLimiter(rate.Every(time.Second/1), 5)
		visitors[ip] = limiter
	}
	return limiter
}

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		limiter := getLimiter(ip)
		if !limiter.Allow() {
			utils.JSONError(w, "Too many requests", http.StatusTooManyRequests, nil)
			return
		}
		next.ServeHTTP(w, r)
	})
}
