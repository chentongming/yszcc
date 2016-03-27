package controller

import (
	"github.com/chentongming/yszcc/application/util"
	"html/template"
	"net/http"
	"path"
)

func IndexHandler(rw http.ResponseWriter, req *http.Request) {
	indexTemplate, err := template.ParseFiles(path.Join(util.HtmlPath, "index.html"))
	if err != nil {
		errTemplate, _ := template.ParseFiles(path.Join(util.HtmlPath, "error", "error.html"))
		errTemplate.ExecuteTemplate(rw, "error.html", map[string]string{"a": "a"})
		return
	}
	indexTemplate.ExecuteTemplate(rw, "index.html", map[string]string{"a": "q"})
}
