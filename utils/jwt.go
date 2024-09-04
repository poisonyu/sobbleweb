package utils

import (
	"errors"
	"time"

	"github.com/cyansobble/global"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type UserClaim struct {
	Uuid        uuid.UUID
	UserID      uint
	UserName    string
	NickName    string
	AuthorityId int
}

type CustomClaim struct {
	UserClaim
	//jwt.StandardClaims
	jwt.RegisteredClaims
	//jwt.RegisteredClaims
}

func CreateCustomClaim(u UserClaim) CustomClaim {
	return CustomClaim{
		UserClaim: u,
		// 	StandardClaims: jwt.StandardClaims{
		// 	Audience:  "cyan",
		// 	NotBefore: 1,
		// 	ExpiresAt: 1,
		// 	//Subject:  "sub",
		// 	Issuer: global.CONFIG.Jwt.Issuer,
		// }
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: global.CONFIG.Jwt.Issuer,
			// Subject: ,
			Audience:  jwt.ClaimStrings{global.CONFIG.Jwt.Audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			// NotBefore: ,
			// IssuedAt: ,
		},
	}
}

func CreateToken(claims CustomClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.CONFIG.Jwt.SignKey))
}

func SetToken(c *gin.Context, token string, maxAge int) {
	c.SetCookie("jwt-token", token, maxAge, "/", "", false, false)
	// c.Header("authorization", token)
}

func GetToken(c *gin.Context) (string, error) {
	token, err := c.Cookie("jwt-token")
	if token == "" {
		token = c.Request.Header.Get("jwt-token")
	}
	return token, err
}

func ParseToken(tokenString string) (*CustomClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString,
		&CustomClaim{},
		func(token *jwt.Token) (i interface{}, e error) {
			return []byte(global.CONFIG.Jwt.SignKey), nil
		},
	)
	if err != nil {
		// global.LOGGER.Error("parse token", zap.Error(err))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")

}
