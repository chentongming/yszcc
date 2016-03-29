package main

import (
	"github.com/chentongming/yszcc/application/controller"
	"github.com/chentongming/yszcc/application/controller/weixin"
	"github.com/chentongming/yszcc/application/util"
	"github.com/chentongming/yszcc/application/util/config"
	"github.com/chentongming/yszcc/application/util/logger"
	"net/http"
	"os"
)

var (
	env  string
	port string
)

func init() {
	if len(os.Args) < 2 {
		env = "test"
	} else {
		env = os.Args[1]
	}
	configPath := util.FilePath("config", env+".ini")
	config.Load(configPath)
}

func main() {
	logger.Info("Listen port:"+config.Get("port")+" start", "info")
	all := map[string]func(rw http.ResponseWriter, req *http.Request){
		"/weixin_mp/msg": weixin.MpHandler,
		"/":              controller.IndexHandler,
		"/hw":            controller.HwHandler,
		"/hw2":           controller.HwHandler2,
	}
	for pattern, handler := range all {
		http.HandleFunc(pattern, handler)
		logger.Info("pattern:"+pattern+",regist handler ", "info")
	}

	http.ListenAndServe(":"+config.Get("port"), nil)
}
