package global

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG Config
	LOGGER *zap.Logger
	DB     *gorm.DB
	router *gin.Engine
)
