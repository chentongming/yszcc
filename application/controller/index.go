package controller

import (
	"github.com/chentongming/yszcc/application/util"
	"github.com/chentongming/yszcc/application/util/logger"
	"html/template"
	"net/http"
	"path"
)

func IndexHandler(rw http.ResponseWriter, req *http.Request) {
	logger.Info(req.RequestURI, "info")
	indexTemplate, err := template.ParseFiles(path.Join(util.HtmlPath, "index.html"))
	if err != nil {
		errTemplate, _ := template.ParseFiles(path.Join(util.HtmlPath, "error", "error.html"))
		errTemplate.ExecuteTemplate(rw, "error.html", map[string]string{"a": "a"})
		return
	}
	indexTemplate.ExecuteTemplate(rw, "index.html", map[string]string{"a": "q"})
}

func HwHandler(rw http.ResponseWriter, req *http.Request) {
	logger.Info(req.RequestURI, "info")
	rw.Write([]byte("hello world"))
}
func HwHandler2(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("hello world"))
}
