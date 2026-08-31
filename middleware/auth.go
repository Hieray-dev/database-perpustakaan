package middleware

import (
	"context"
	"net/http"
	"perpustakaan-api/utils"
	"strings"
)

type contextKey string

const UserCtxKey contextKey = "user_claims"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"message":"Header Authorization diperlukan"}`, http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, `{"message":"Token tidak valid atau expired"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserCtxKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(UserCtxKey).(*utils.Claims)
		if !ok || claims.IDRole != 1 {
			http.Error(w, `{"message":"Akses ditolak! Khusus Admin"}`, http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func PetugasOrAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(UserCtxKey).(*utils.Claims)
		if !ok || (claims.IDRole != 1 && claims.IDRole != 2) {
			http.Error(w, `{"message":"Akses ditolak! Khusus Petugas atau Admin"}`, http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
