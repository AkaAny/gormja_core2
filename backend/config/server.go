package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	Port int
}

func (x *ServerConfig) RunGinEngine(engine *gin.Engine) {
	var bindAddress = fmt.Sprintf(":%d", x.Port)
	if err := engine.Run(bindAddress); err != nil {
		panic(err)
	}
}
