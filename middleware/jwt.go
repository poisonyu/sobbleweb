package middleware

import (
	"time"

	"github.com/cyansobble/global"
	"github.com/cyansobble/response"
	"github.com/cyansobble/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := utils.GetToken(c)
		if token == "" {
			global.LOGGER.Info("empty token")
			response.JSONResponse(c, 0, "jwt auth failed", nil)
			c.Abort()
			return
		}
		//global.LOGGER.Info(token)
		claims, err := utils.ParseToken(token)
		if err != nil {
			global.LOGGER.Error("parse token", zap.Error(err))
			response.JSONResponse(c, 0, "jwt auth failed", nil)
			c.Abort()
			return
		}
		if claims.ExpiresAt.Unix() <= time.Now().Unix() {

			global.LOGGER.Error("token expired", zap.Error(err))
			response.JSONResponse(c, 0, "jwt auth failed", nil)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
