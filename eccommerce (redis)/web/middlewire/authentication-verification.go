package middlewire

import (
	"context"
	"ecommerce/config"
	"ecommerce/web/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var userKey = "user"

type AuthClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func unauthorizedResponse(w http.ResponseWriter) {
	utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// collect token from header
		header := r.Header.Get("authorization")
		tokenStr := ""
		fmt.Println(header)
		// collect token from query
		if len(header) == 0 {
			tokenStr = r.URL.Query().Get("auth")
		} else {
			tokens := strings.Split(header, " ")
			if len(tokens) != 2 {
				unauthorizedResponse(w)
				return
			}
			tokenStr = tokens[1]
		}

		// parse jwt
		var claims AuthClaims
		token, err := jwt.ParseWithClaims(
			tokenStr,
			&claims,
			func(t *jwt.Token) (interface{}, error) {
				return []byte(config.GetConfig().JwtSecretKey), nil
			},
		)
		fmt.Println(claims, token, "hi")
		if err != nil {
			unauthorizedResponse(w)
			return
		}

		// get user id from token
		if !token.Valid {
			unauthorizedResponse(w)
			return
		}

		// set user id in the context
		wrappedRequest := r.WithContext(context.WithValue(r.Context(), userKey, claims.Id))
		next.ServeHTTP(w, wrappedRequest)
	})
}

func GetUserId(r *http.Request) (int, error) {
	userIdVal := r.Context().Value(userKey)
	userId, ok := userIdVal.(int)
	if !ok {
		return 0, fmt.Errorf("Unauthorized")
	}
	return userId, nil
}
