package middleware

import (
	"github.com/cyansobble/global"
	"github.com/cyansobble/response"
	"github.com/cyansobble/utils"
	"github.com/gin-gonic/gin"
)

func AuthorityAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			global.LOGGER.Info("no claims")
			response.JSONResponse(c, 0, "请登录哦！", nil)
			c.Abort()
			return
		}

		if claims.(*utils.CustomClaim).AuthorityId != 7 {
			global.LOGGER.Error("权限不足")
			response.JSONResponse(c, 0, "jwt auth failed", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
