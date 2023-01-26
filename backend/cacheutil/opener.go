package cacheutil

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/url"
	"strconv"
)

type CacheOpener struct {
	Type string
	DSN  string
}

func (x *CacheOpener) ToCache() interface{} {
	switch x.Type {
	case "redis":
		return redisOpener(x.DSN)
	default:
		panic(fmt.Errorf("unknown cache type:%s", x.Type))
	}
}

func redisOpener(dsn string) *redis.Client {
	dsnURL, err := url.Parse(dsn)
	if err != nil {
		panic(fmt.Errorf("parse dsn with err:%w", err))
	}
	var userName = dsnURL.Query().Get("userName")
	var password = dsnURL.Query().Get("password")
	var dbStr = dsnURL.Query().Get("db")
	db, err := strconv.ParseInt(dbStr, 10, 64)
	if err != nil {
		panic(fmt.Errorf("invalid db str:%s", dbStr))
	}
	var redisClient = redis.NewClient(&redis.Options{
		Network:  dsnURL.Scheme,
		Addr:     dsnURL.Host,
		Username: userName,
		Password: password,
		DB:       int(db),
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return redisClient
}
