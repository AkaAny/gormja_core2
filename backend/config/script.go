package config

import (
	"fmt"
	"github.com/guonaihong/gout"
	"os"
)

type ScriptConfig struct {
	Type   string
	Source string
}

func (x *ScriptConfig) Open() string {
	switch x.Type {
	case "file":
		return openFile(x.Source)
	case "http":
		return openHttpConfig(x.Source)
	default:
		panic(fmt.Errorf("unknown script source type:%s", x.Type))
	}
}

func openHttpConfig(url string) string {
	var rawData = make([]byte, 0)
	if err := gout.New().GET(url).BindBody(rawData).Do(); err != nil {
		panic(err)
	}
	return string(rawData)
}

func openFile(path string) string {
	rawData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(rawData)
}
