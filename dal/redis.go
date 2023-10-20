package dal

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/csguojin/reserve/util/logger"
)

func Lock(ctx context.Context, rdb *redis.Client, key string, value string, t time.Duration) (bool, error) {
	lockAcquired, err := rdb.SetNX(ctx, key, value, t).Result()
	if err != nil {
		logger.L.Errorln("Error acquiring lock:", err)
		return false, err
	}
	if !lockAcquired {
		logger.L.Errorln("Failed to acquire lock. Cur lock is performing an operation.")
		return false, err
	}
	return true, nil
}

func UnLock(ctx context.Context, rdb *redis.Client, key string, value string) error {
	_, err := rdb.Eval(ctx, UnlockScript, []string{key}, value).Result()
	if err != nil {
		logger.L.Errorln("Error releasing lock:", err)
		return err
	}
	return nil
}

const UnlockScript = `
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end
`
