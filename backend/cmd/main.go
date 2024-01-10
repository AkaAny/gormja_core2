package main

import (
	"github.com/AkaAny/config-tv"
	"github.com/AkaAny/config-tv/plugin/k8s_configmap"
	"github.com/AkaAny/config-tv/plugin/k8s_secret"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	backend "gormja_core2"
	"gormja_core2/config"
	"gormja_core2/router"
)

func main() {
	var configPluginConfig = config_tv.GetConfigPluginConfigFromEnv()
	var pluginMap = make(config_tv.TypePluginMap)
	{
		var pluginConfigMap = configPluginConfig.Plugin[k8s_secret.PluginName]
		var k8sSecretTypeKVPlugin = k8s_secret.NewK8sSecretPluginFromConfig(pluginConfigMap)
		pluginMap[k8s_secret.PluginName] = k8sSecretTypeKVPlugin
	}
	{
		var pluginConfigMap = configPluginConfig.Plugin[k8s_configmap.PluginName]
		var k8sConfigMapTypeKVPlugin = k8s_configmap.NewK8sConfigPluginFromConfig(pluginConfigMap)
		pluginMap[k8s_configmap.PluginName] = k8sConfigMapTypeKVPlugin
	}
	var mainConfig = new(config.Config)
	{
		config_tv.GetAndUnmarshalMainConfigFromEnv(mainConfig, pluginMap)
	}
	var dbRegistry = backend.NewDBRegistry()
	for id, cfgItem := range mainConfig.DB {
		var db = cfgItem.ToDB()
		dbRegistry.Register(id, db)
	}
	var cacheClient backend.CacheClient = nil
	if mainConfig.EnableCache {
		var cacheConfigItem = mainConfig.Cache["main"]

		var redisClient = cacheConfigItem.ToClient().(*redis.Client)
		cacheClient = backend.NewRedisCacheClient(redisClient)
	} else {
		cacheClient = &backend.NoCacheCacheClient{}
	}
	var runtimeRegistry = backend.NewRuntimeRegistry()
	var logger = logrus.New()
	//init startup script
	for id, cfgItem := range mainConfig.Script {
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
	mainConfig.Server.RunGinEngine(engine)
}
