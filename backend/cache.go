package backend

import "context"

type CacheClient interface {
	Set(ctx context.Context, cacheKey string, dests []interface{}, expireMillisecond int64) error
	Get(ctx context.Context, cacheKey string) (dests []interface{}, cacheExist bool, err error)
}

type NoCacheCacheClient struct {
}

func (ncc *NoCacheCacheClient) Set(ctx context.Context, cacheKey string, dests []interface{}, expireMillisecond int64) error {
	return nil
}

func (ncc *NoCacheCacheClient) Get(ctx context.Context, cacheKey string) (dests []interface{}, cacheExist bool, err error) {
	return nil, false, nil
}
