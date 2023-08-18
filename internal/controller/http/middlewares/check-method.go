package middlewares

import (
	"errors"
	"log"
	"net/http"
)

var ErrUnexpectedMethodRequest = errors.New("send unexpected method request")

// Middleware for check request method
func MethodCheck(allowedMethod string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedMethod != r.Method {
			http.Error(w, ErrUnexpectedMethodRequest.Error(), http.StatusMethodNotAllowed)
			log.Printf("[ERR] получен неожиданный метод запроса '%s', ожидается '%s'", r.Method, allowedMethod)

			return
		}

		next.ServeHTTP(w, r)
	})
}
