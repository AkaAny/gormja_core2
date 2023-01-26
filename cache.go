package gormja_core2

import "context"

type CacheClient interface {
	Set(ctx context.Context, cacheKey string, dests []interface{}, expireMillisecond int64) error
	Get(ctx context.Context, cacheKey string) (dests []interface{}, cacheExist bool, err error)
}
