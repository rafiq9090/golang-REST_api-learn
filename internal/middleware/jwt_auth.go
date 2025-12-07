package middleware

import (
	"net/http"
	"strings"
	"task-api/internal/config"
	"task-api/internal/utils"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.App.JWTSecret)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.JSONError(w, "Unauthorized", http.StatusUnauthorized, nil)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.JSONError(w, "Unauthorized", http.StatusUnauthorized, nil)
			return
		}

		tokenStr := parts[1]
		token, err := jwt.ParseWithClaims(tokenStr, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			utils.JSONError(w, "Unauthorized", http.StatusUnauthorized, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
