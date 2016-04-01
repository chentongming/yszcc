package requestUtil

import (
	"github.com/chentongming/yszcc/application/util/requestUtil"
	"testing"
)

func TestGet(t *testing.T) {
	result, _ := requestUtil.Get("http://www.mia.com/instant/account/loginService", nil)
	if result["code"] == 500 {
		t.Error(result["code"])
	}
}
