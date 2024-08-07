package main

import (
	"github.com/cyansobble/global"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	//jwt "github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Uuid        uuid.UUID
	UserName    string
	NickName    string
	AuthorityId int
}

type Claim struct {
	UserClaim
	jwt.StandardClaims
	//jwt.RegisteredClaims
}

func GetClainm(u UserClaim) *Claim {
	return &Claim{UserClaim: u, StandardClaims: jwt.StandardClaims{
		Audience:  "cyan",
		NotBefore: 1,
		ExpiresAt: 1,
		//Subject:  "sub",
		Issuer: global.CONFIG.Jwt.Issuer,
	}}
}
