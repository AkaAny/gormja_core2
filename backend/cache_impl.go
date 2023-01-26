package backend

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCacheClient struct {
	redisClient *redis.Client
}

func NewRedisCacheClient(redisClient *redis.Client) *RedisCacheClient {
	return &RedisCacheClient{redisClient: redisClient}
}

func (rcc *RedisCacheClient) Set(ctx context.Context, cacheKey string, dests []interface{}, expireMillisecond int64) error {
	rawData, err := json.Marshal(dests)
	if err != nil {
		return fmt.Errorf("marshal dests:%v with err:%w", dests, err)
	}
	var expireDuration = time.Millisecond * time.Duration(expireMillisecond)
	var cmdHandle = rcc.redisClient.Set(ctx, cacheKey, string(rawData), expireDuration)
	if err := cmdHandle.Err(); err != nil {
		return fmt.Errorf("set cmd handle with err:%w", err)
	}
	return nil
}

func (rcc *RedisCacheClient) Get(ctx context.Context, cacheKey string) (dests []interface{}, cacheExist bool, err error) {
	var cmdHandle = rcc.redisClient.Get(ctx, cacheKey)
	if err := cmdHandle.Err(); err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, false, nil
		} else {
			return nil, false, fmt.Errorf("get cmd handle with err:%w", err)
		}
	}
	valueStr, _ := cmdHandle.Result() //checked if err above
	var rawData = []byte(valueStr)
	if err := json.Unmarshal(rawData, &dests); err != nil {
		return nil, true, fmt.Errorf("unmarshal dests with err:%w", err)
	}
	return dests, true, nil
}
