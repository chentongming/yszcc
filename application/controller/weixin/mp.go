package weixin

import (
	"github.com/chentongming/yszcc/application/util/wxUtil"
	"io/ioutil"
	"net/http"
)

func MpHandler(rw http.ResponseWriter, req *http.Request) {
	bytes, _ := ioutil.ReadAll(req.Body)
	var wxMsg wxUtil.WxMsgReceived
	wxMsg.Set(string(bytes))
	dealTextMsg(&wxMsg, &rw)
}

func dealTextMsg(wxTextMsg *wxUtil.WxMsgReceived, rw *http.ResponseWriter) {
	http.ResponseWriter(*rw).Write([]byte(wxTextMsg.Content))
}
