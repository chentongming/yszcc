package test

import (
	"github.com/chentongming/yszcc/application/util/wxUtil"
	"testing"
)

func TestCalc(t *testing.T) {
	wxValidData := wxUtil.WxServerValidData{Timestamp: 1234, Nonce: "abc"}
	a := wxValidData.Calc("aa")
	if a == "" {
		t.Error(a)
	}
}

func TestValid(t *testing.T) {
	wxValidData := wxUtil.WxServerValidData{Timestamp: 1234, Nonce: "abc"}
	if wxValidData.Valid("aa") {
		t.Error("false")
	}
}

func TestGetAccessToken(t *testing.T) {
	result := wxUtil.GetAccessToken("123456789", "1234567890")
	if result["access_token"] == "" {
		t.Error(result["access_token"])
	}
}
func TestWxMsg(t *testing.T) {
	xmlStr := `<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[FromUser]]></FromUserName>
<CreateTime>123456789</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Event><![CDATA[subscribe]]></Event>
</xml>`
	var wxMsg wxUtil.WxMsgReceived
	err := wxMsg.Set(xmlStr)
	if err != nil {
		t.Error(err)
	}
	if wxMsg.MsgType != "text" {
		t.Error(wxMsg.MsgType, "!=text")
	}
}
