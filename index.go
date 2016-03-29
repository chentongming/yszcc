package main

import (
	"fmt"
	"github.com/chentongming/yszcc/application/controller"
	"github.com/chentongming/yszcc/application/controller/weixin"
	"github.com/chentongming/yszcc/application/util"
	"github.com/chentongming/yszcc/application/util/config"
	"log"
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
	http.HandleFunc("/weixin_mp/msg", weixin.MpHandler)
	http.HandleFunc("/", controller.IndexHandler)
	fmt.Println(":" + config.Get("port"))
	log.Fatal(http.ListenAndServe(":"+config.Get("port"), nil))
}
