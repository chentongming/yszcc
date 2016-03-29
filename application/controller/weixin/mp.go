package weixin

import (
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
	wxResponseText := wxUtil.WxResponseText{
		ToUserName:   wxTextMsg.FromUserName,
		FromUserName: config.Get("wxUser"),
		CreateTime:   int(time.Now().Unix()),
		Content:      wxTextMsg.Content,
		MsgType:      "text",
	}
	http.ResponseWriter(*rw).Write([]byte(wxResponseText.ToXml()))
}
