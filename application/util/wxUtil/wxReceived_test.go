package wxUtil

import (
	"fmt"
	"github.com/chentongming/yszcc/application/util/wxUtil"
	"testing"
)

func TestToXml(t *testing.T) {
	response := wxUtil.WxResponseText{
		ToUserName:   "abg",
		FromUserName: "abg",
		CreateTime:   12358,
		Content:      "abg",
	}
	str, _ := response.ToXml()
	t.Error(str)
	fmt.Println(str)
}
