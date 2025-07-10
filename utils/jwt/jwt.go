package jwt

import (
	"net/http"
	"time"

	"github.com/emenesism/Decentralized-voting-backend/config"
	"github.com/golang-jwt/jwt"
)

func GenToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id": id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.AppConfig.Jwt_secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil	
}

func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, http.ErrAbortHandler
        }
        return []byte(config.AppConfig.Jwt_secret), nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, err
    }

    if err := claims.Valid(); err != nil {
        return nil, err
    }

    return claims, nil
}