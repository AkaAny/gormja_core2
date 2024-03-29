package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	backend "gormja_core2"
	"net/http"
)

func RegisterDataRouter(group gin.IRouter, runtimeRegistry *backend.RuntimeRegistry) {
	group.POST("/lookup", func(c *gin.Context) {
		var runtimeID = c.Query("runtimeID")
		var topic = c.Query("topic")
		var ctx = context.Background()
		condMap, err := unmarshalCondMap(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		serviceRuntime, err := runtimeRegistry.Get(runtimeID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": err.Error(),
			})
			return
		}
		dests, err := serviceRuntime.Lookup(ctx, topic, condMap)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": dests,
		})
	})
	group.PUT("/manualRefresh", func(c *gin.Context) {
		var runtimeID = c.Query("runtimeID")
		var topic = c.Query("topic")
		var ctx = context.Background()
		condMap, err := unmarshalCondMap(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": fmt.Errorf("unamarhsal cond map with err:%w", err).Error(),
			})
			return
		}
		serviceRuntime, err := runtimeRegistry.Get(runtimeID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": err.Error(),
			})
			return
		}
		if err := serviceRuntime.ManualRefresh(ctx, topic, condMap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	})
	group.GET("/description", func(c *gin.Context) {
		runtimeID, serviceID := c.Query("runtimeID"), c.Query("topic")
		serviceRuntime, err := runtimeRegistry.Get(runtimeID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
			})
			return
		}
		serviceObj, err := serviceRuntime.GetServiceByID(serviceID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
			})
			return
		}
		var serviceDesc = serviceObj.GetDescription()
		c.JSON(http.StatusOK, gin.H{
			"data": serviceDesc,
		})
	})
}

func unmarshalCondMap(c *gin.Context) (map[string]interface{}, error) {
	var condMap = make(map[string]interface{})
	err := c.BindJSON(&condMap)
	if err != nil {
		return nil, fmt.Errorf("unamarhsal cond map with err:%w", err)
	}
	return condMap, nil
}
