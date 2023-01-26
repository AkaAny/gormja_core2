package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	backend "gormja_core2"
	"net/http"
)

func NewRouter(serviceRuntime *backend.ServiceRuntime) *gin.Engine {
	var engine = gin.Default()
	engine.POST("/lookup", func(c *gin.Context) {
		var topic = c.Query("topic")
		var ctx = context.Background()
		condMap, err := unmarshalCondMap(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": fmt.Errorf("unamarhsal cond map with err:%w", err),
			})
			return
		}
		dests, err := serviceRuntime.Lookup(ctx, topic, condMap)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": dests,
		})
	})
	engine.PUT("/manualRefresh", func(c *gin.Context) {
		var topic = c.Query("topic")
		var ctx = context.Background()
		condMap, err := unmarshalCondMap(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": fmt.Errorf("unamarhsal cond map with err:%w", err),
			})
			return
		}
		if err := serviceRuntime.ManualRefresh(ctx, topic, condMap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	})
	return engine
}

func unmarshalCondMap(c *gin.Context) (map[string]interface{}, error) {
	var condMap = make(map[string]interface{})
	err := c.BindJSON(condMap)
	if err != nil {
		return nil, fmt.Errorf("unamarhsal cond map with err:%w", err)
	}
	return condMap, nil
}
