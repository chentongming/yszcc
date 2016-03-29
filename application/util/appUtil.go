package util

import (
	"path"
	"path/filepath"
)

var (
	RootPath string
	basePath string
	HtmlPath string
)

func init() {
	basePath, _ = filepath.Abs(".")
	RootPath = path.Join(basePath, "application")
	HtmlPath = path.Join(RootPath, "html")
}

func FilePath(args ...string) string {
	filepath := RootPath
	for _, v := range args {
		filepath = path.Join(filepath, v)
	}
	return filepath
}
