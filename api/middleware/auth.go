package middleware

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

var (
	adminToken = "admin"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check auth header exists and starts with Bearer
		authorization := r.Header.Get("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}
		// check token
		encodedToken := strings.TrimPrefix(authorization, "Bearer ")
		token, err := base64.StdEncoding.DecodeString(encodedToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}

		// Compare token
		if string(token) != adminToken {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}

		log.Println("request authorized")
		next.ServeHTTP(w, r)
	})
}
