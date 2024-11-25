package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type contextKey string

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

const UserContextKey = contextKey("user")

func JWTAuthMiddleware(secret string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                http.Error(w, "Authorization header required", http.StatusUnauthorized)
                return
            }

            parts := strings.Split(authHeader, " ")
            if len(parts) != 2 || parts[0] != "Bearer" {
                http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
                return
            }

            tokenStr := parts[1]
            claims := &jwt.StandardClaims{}
            token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
                return []byte(secret), nil
            })
			

            if err != nil {
                log.Printf("Token parsing error: %v", err)
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            if !token.Valid {
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            ctx := context.WithValue(r.Context(), UserContextKey, claims.Subject)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}