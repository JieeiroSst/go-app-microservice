package jwt

import (
	"github.com/JIeeiroSst/go-app/config"
	"github.com/JIeeiroSst/go-app/internal/models/user"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenUser struct {
	u user.Users
	conf *config.Config
}

func NewTokenUser(u user.Users,conf *config.Config) *TokenUser{
	return &TokenUser{u:u,conf:conf}
}

func (t *TokenUser) GenerateToken() (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = t.u.Id
	atClaims["username"]= t.u.Username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(t.conf.Server.JwtSecretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t *TokenUser) ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.conf.Server.JwtSecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "Missing Authentication Token", err
	}
}
