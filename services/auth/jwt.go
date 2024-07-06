package auth

import (
	"strconv"
	"time"

	"github.com/Ayobami6/go_ecom/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	sec, err := strconv.ParseInt(config.Envs.JWTExpiration, 10, 64)
	if err!= nil {
        return "", err
    }
	expiration := time.Second * time.Duration(sec)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err!= nil {
        return "", err
    }
	return tokenString, nil

} 