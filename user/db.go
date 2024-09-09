package user

import (
	"context"
	"time"

	"github.com/cyansobble/global"
	"go.uber.org/zap"
)

func GetUserByID(id uint) (User, error) {
	var user User
	err := global.DB.First(&user, id).Error
	return user, err
}

func SaveUser(u User) error {
	result := global.DB.Save(&u)
	return result.Error
}

func SetStringInRedis(key string, value interface{}, expiration time.Duration) (val string, err error) {
	ctx := context.Background()
	val, err = global.RedisDb.Set(ctx, key, value, expiration).Result()
	if val != "OK" {
		global.LOGGER.Error("set failed", zap.Error(err))
		return
	}
	return
}
