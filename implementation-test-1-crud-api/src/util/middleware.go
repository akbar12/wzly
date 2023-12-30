package util

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		authHead := strings.Split(r.Header.Get("Authorization"), " ")
		var ctx context.Context
		if len(authHead) > 1 {
			jwtData, err := VerifyJWT(authHead[1])
			if err != nil {
				ResponseError401(w, ErrUnauthorized)
				return
			}
			ctx = context.WithValue(r.Context(), "username", jwtData.Username)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type JWTData struct {
	Username string
}

func VerifyJWT(tokenString string) (*JWTData, error) {
	var MySigningKey = []byte(GetConfigString("jwt.key"))
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return MySigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uname, ok := claims["UNAME"].(string)
		if !ok {
			return nil, errors.New("failed claim auth UNAME data")
		}
		exp, ok := claims["EXP"].(float64)
		if !ok {
			return nil, errors.New("failed claim auth EXPIRED data")
		}
		// check expired
		expiredTime := time.Unix(int64(exp), 0)
		if time.Now().After(expiredTime) {
			return nil, fmt.Errorf("ACCESS_TOKEN_EXPIRED")
		}

		return &JWTData{
			Username: uname,
		}, nil
	}
	return nil, fmt.Errorf("ACCESS_TOKEN_EXPIRED")
}

// func ()
