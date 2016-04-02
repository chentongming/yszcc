package weixin

import (
	"github.com/chentongming/yszcc/application/service/mp"
	"github.com/chentongming/yszcc/application/util/config"
	"github.com/chentongming/yszcc/application/util/logger"
	"github.com/chentongming/yszcc/application/util/wxUtil"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func MpHandler(rw http.ResponseWriter, req *http.Request) {
	bytes, _ := ioutil.ReadAll(req.Body)
	logger.Info(req.RequestURI, "info")
	if req.Method == "GET" {
		timeStamp, err := strconv.Atoi(req.FormValue("timestamp"))
		if err != nil {
			return
		}
		serverValid := &wxUtil.WxServerValidData{
			Signature: req.FormValue("signature"),
			Timestamp: int(timeStamp),
			Nonce:     req.FormValue("nonce"),
			Echostr:   req.FormValue("echostr"),
		}
		if serverValid.Valid(config.Get("wxToken")) {
			rw.Write([]byte(req.FormValue("echostr")))
		}
	} else {
		var wxMsg wxUtil.WxMsgReceived
		wxMsg.Set(string(bytes))
		dealTextMsg(&wxMsg, &rw)
	}

}

func dealTextMsg(wxTextMsg *wxUtil.WxMsgReceived, rw *http.ResponseWriter) {
	var content string
	if wxTextMsg.Content == "上证" {
		stockCode := "sh000001"
		content = mp.Sh(stockCode)
	} else if wxTextMsg.Content == "深成" {
		stockCode := "sz399001"
		content = mp.Sh(stockCode)
	} else {
		content = wxTextMsg.Content
	}
	wxResponseText := wxUtil.WxResponseText{
		ToUserName:   wxTextMsg.FromUserName,
		FromUserName: config.Get("wxUser"),
		CreateTime:   int(time.Now().Unix()),
		Content:      content,
		MsgType:      "text",
	}
	http.ResponseWriter(*rw).Write([]byte(wxResponseText.ToXml()))
}
