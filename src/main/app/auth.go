package app

import (
	"net/http"
	"strings"
	"context"
	"os"
	jwt "github.com/dgrijalva/jwt-go"
	"main/models"
	u "main/utils"
)

var JwtAuthentication = func (next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r * http.Request) {
		res := make(map[string] interface{})
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			res = u.Message(false, "Missing Authorization header")
			u.Respond(w, res, http.StatusForbidden)
			return
		}
	
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			res = u.Message(false, "Invalid Authorization")
			u.Respond(w, res, http.StatusForbidden)
			return
		}
	
		requestToken := splitted[1]
		tk := &models.Token{}
		token, e := jwt.ParseWithClaims(requestToken, tk, func (token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwt_secret")), nil
		})
	
		if e != nil || !token.Valid {
			res = u.Message(false, "Invalid token")
			u.Respond(w, res, http.StatusForbidden)
			return
		}
	
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}