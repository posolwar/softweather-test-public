package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/posolwar/softweather-test/internal/helpers"
)

var ErrAccessDenied = errors.New("Access denied")

// Middleware for check access to handlers
func AccessCheck(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAccess := r.Header.Get(helpers.UserAccess)

		if userAccess != helpers.AllowedUser {
			http.Error(w, ErrAccessDenied.Error(), http.StatusUnauthorized)
			log.Printf("[ERR] запрет доступа для %s на хандлер - %s, user-agent - %s", r.RemoteAddr, r.URL.Path, userAccess)

			return
		}

		next.ServeHTTP(w, r)
	})
}
