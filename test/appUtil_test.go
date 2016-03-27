package test

import (
	"github.com/chentongming/yszcc/application/util"
	"testing"
)

func Test(t *testing.T) {
	if util.BasePath == nil {
		t.Error("fail")
	}
}
