package backend

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gormja_core2/config"
	"testing"
)

func TestServiceRuntime_LoadScript(t *testing.T) {
	var scriptContent = GetScriptContent()
	var cfg = config.Unmarshal(config.NewDefaultViper())
	var dbRegistry = NewDBRegistry()
	const dataSourceID = "hdu-oracle"
	var db = cfg.DB[dataSourceID].ToDB()
	dbRegistry.Register(dataSourceID, db)
	var redisClient = cfg.Cache["main"].ToClient().(*redis.Client)
	var cacheClient = NewRedisCacheClient(redisClient)
	var serviceRuntime = NewServiceRuntime(dbRegistry, cacheClient, logrus.New())
	serviceRuntime.Init()
	serviceRuntime.LoadScript(scriptContent)
	dests, err := serviceRuntime.Lookup(context.Background(), "ProfileService", map[string]interface{}{
		"StaffID": "20113128",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(dests)
}
