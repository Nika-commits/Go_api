package middleware

import (
	"errors"
	"net/http"
	"github.com/Nika-commits/Go_api/api"
	// "github.com/Nika-commits/Go_api/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalid credentials")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request){
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username =="" || token == "" {
			
		}
	})
}