package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	backend "gormja_core2"
	"gormja_core2/config"
	"gormja_core2/router"
)

func main() {
	var vh = config.NewDefaultViper()
	var cfg = config.Unmarshal(vh)
	var dbRegistry = backend.NewDBRegistry()
	for id, cfgItem := range cfg.DB {
		var db = cfgItem.ToDB()
		dbRegistry.Register(id, db)
	}
	var cacheConfigItem = cfg.Cache["main"]
	var redisClient = cacheConfigItem.ToClient().(*redis.Client)
	var cacheClient = backend.NewRedisCacheClient(redisClient)
	var runtimeRegistry = backend.NewRuntimeRegistry()
	var logger = logrus.New()
	//init startup script
	for id, cfgItem := range cfg.Script {
		var scriptStr = cfgItem.Open()
		var serviceRuntime = backend.NewServiceRuntime(dbRegistry, cacheClient, logger)
		serviceRuntime.Init()
		serviceRuntime.LoadScript(scriptStr)
		runtimeRegistry.Put(id, serviceRuntime)
	}
	//router
	var engine = gin.Default()
	router.RegisterDataRouter(engine.Group("/data"), runtimeRegistry)
	router.RegisterManageRouter(engine.Group("/runtime"), runtimeRegistry,
		func(script string) (*backend.ServiceRuntime, error) {
			var serviceRuntime = backend.NewServiceRuntime(dbRegistry, cacheClient, logger)
			serviceRuntime.Init()
			serviceRuntime.LoadScript(script)
			return serviceRuntime, nil
		})
	cfg.Server.RunGinEngine(engine)
}
