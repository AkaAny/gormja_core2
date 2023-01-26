package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	backend "gormja_core2"
	"io"
	"net/http"
)

func ManagerAPIs(group gin.IRouter,
	runtimeRegistry *backend.RuntimeRegistry, factoryFunc ServiceRuntimeFactoryFunc) {
	group.PUT("/put", func(c *gin.Context) {
		var runtimeID = c.Query("id")
		fh, err := c.FormFile("scriptFile")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": fmt.Errorf("read multipart file header with err:%w", err),
			})
			return
		}
		f, err := fh.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": fmt.Errorf("open multipart file with err:%w", err),
			})
			return
		}
		defer f.Close()
		scriptFileData, err := io.ReadAll(f)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": fmt.Errorf("read script file with err:%w", err),
			})
			return
		}
		serviceRuntime, err := factoryFunc(string(scriptFileData))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": fmt.Errorf("factory func with err:%w", err),
			})
			return
		}
		runtimeRegistry.Put(runtimeID, serviceRuntime)
		c.JSON(http.StatusOK, gin.H{})
	})
}
