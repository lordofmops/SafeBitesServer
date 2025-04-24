package middleware

import (
	"context"
	"net/http"
	"strings"

	"SafeBitesServer/pkg/response"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret = []byte("your-secret-key") // TODO: загружать из конфига

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			response.Error(w, http.StatusUnauthorized, "отсутствует токен")
			return
		}
		parts := strings.Split(auth, "Bearer ")
		if len(parts) != 2 {
			response.Error(w, http.StatusUnauthorized, "неверный формат токена")
			return
		}
		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			response.Error(w, http.StatusUnauthorized, "невалидный токен")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Error(w, http.StatusUnauthorized, "невалидные claims")
			return
		}

		idStr, ok := claims["user_id"].(string)
		if !ok {
			response.Error(w, http.StatusUnauthorized, "не найден user_id")
			return
		}
		userID, err := uuid.Parse(idStr)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, "невалидный UUID")
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
