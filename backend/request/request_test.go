package request

import (
	"fmt"
	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gormja_core2/utils"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	var engine = gin.Default()
	engine.POST("/test", func(c *gin.Context) {
		fmt.Println(c.Request)
		var expectedStatusCode = http.StatusOK
		var expectedStatusCodeStr = c.GetHeader("X-Expected-StatusCode")
		if expectedStatusCodeStr != "" {
			expectedStatusCodeInt64, err := strconv.ParseInt(expectedStatusCodeStr, 10, 64)
			if err != nil {
				panic(err)
			}
			expectedStatusCode = int(expectedStatusCodeInt64)
		}
		c.JSON(expectedStatusCode, gin.H{
			"status": "ok",
		})
	})
	go func() {
		err := engine.Run(":18085")
		if err != nil {
			panic(err)
		}
		fmt.Println("ready")
	}()
	time.Sleep(100 * time.Millisecond)
	var runtime = goja.New()
	var baseLogger = logrus.New()
	var jsConsole = utils.NewJSConsole(baseLogger)
	jsConsole.Attach(runtime)
	MakeFetch(runtime)
	runtime.RunString(`
		fetch('http://localhost:18085/test',{
			method:"POST",
			headers:{
				"hk1":"hv1",
				"Content-Type":"application/json",
				"X-Expected-StatusCode":"400",
			},
			body:"{\"a\":\"b\"}"
		}).then((resp)=>{
			return resp.json()
		}).then((jsonObj)=>{
			console.log(jsonObj)
		})
`)
}
