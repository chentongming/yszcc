package main

import (
	"github.com/chentongming/yszcc/application/controller"
	"github.com/chentongming/yszcc/application/controller/weixin"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/weixin_mp/", weixin.MpHandler)
	http.HandleFunc("/", controller.IndexHandler)
	log.Fatal(http.ListenAndServe(":1200", nil))
}
