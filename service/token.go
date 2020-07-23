package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/i-coder-robot/go-micro-action-user/proto/auth"
	"github.com/lecex/core/env"
	"strconv"
	"time"
)

var (
	privateKey = []byte(env.Getenv("APP_KEY", "8ca96774aadf77668e62931b9c0a14e5"))
)

type Authable interface {
	Encode(user *auth.User) (string, error)
	Decode(tokenStr string) (*CustomClaims, error)
}

type CustomClaims struct {
	User *auth.User
	jwt.StandardClaims
}

type TokenService struct {
}

func (srv *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, err
}

func (srv *TokenService) Encode(user *auth.User) (string, error) {
	validityPeriod, _ := strconv.ParseInt(env.Getenv("TOKEN_VALIDITY_PERIOD", 3), 10, 64)
	expireTime := time.Now().Add(time.Hour * 24 * time.Duration(validityPeriod)).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			Issuer:    `user`,
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(privateKey)
}
