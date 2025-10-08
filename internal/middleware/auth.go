package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/Fedesin/goapi-practica/internal/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := r.Header.Get("Authorization")
		if h == "" {
			http.Error(w, "Token faltante", http.StatusUnauthorized)
			return
		}
		if len(h) < len("Bearer "){
			http.Error(w, "Header Authorization inválido", http.StatusUnauthorized)
			return
		}
		tokenStr := h[len("Bearer "):]

		token, err := utils.ParseToken(tokenStr)
		if err != nil || token == nil || !token.Valid {
			http.Error(w, "Token inválido o expirado", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Token malformado", http.StatusUnauthorized)
			return
		}

		emailAny := claims["email"]
		email, ok := emailAny.(string)
		if !ok {
			http.Error(w, "Claim email ausente", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
