package test

import (
	"fmt"
	"github.com/chentongming/yszcc/application/util/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	err := config.Load("E:/work/mia/code/go/src/github.com/chentongming/yszcc/application/config/test.ini")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(config.Get("port"))
}
