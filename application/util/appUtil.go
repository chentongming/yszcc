package util

import (
	"path"
	"path/filepath"
)

var (
	RootPath string
	BasePath string
	HtmlPath string
)

func init() {
	BasePath, _ = filepath.Abs(".")
	RootPath = path.Join(BasePath, "application")
	HtmlPath = path.Join(RootPath, "html")
}
