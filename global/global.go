package global

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	CONFIG Config
	LOGGER *zap.Logger
	DB     *gorm.DB
	// router *gin.Engine
	RedisDb *redis.Client
)
