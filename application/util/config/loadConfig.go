package config

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var (
	defaultSection = "_default"
	config         = make(map[string]string)
	// loadConfig map[string]int
)

func Load(iniPath string) error {
	file, err := os.Open(iniPath)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var line string
	for {
		tmpLine, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if isPrefix == true {
			line += string(tmpLine)
			continue
		} else {
			line = string(tmpLine)
		}
		lineConfig := strings.Split(line, "=")
		config[lineConfig[0]] = lineConfig[1]
	}
	return nil
}

func Get(key string) string {
	return config[key]
}
