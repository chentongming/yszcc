package weixin

import (
	"fmt"
	"github.com/chentongming/yszcc/application/util/config"
	"github.com/chentongming/yszcc/application/util/wxUtil"
	"io/ioutil"
	"net/http"
	"strconv"
)

func MpHandler(rw http.ResponseWriter, req *http.Request) {
	bytes, _ := ioutil.ReadAll(req.Body)
	fmt.Println("post", string(bytes))
	fmt.Println("income request:", req.RequestURI, " data:", string(bytes))
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
		if serverValid.Valid(config.Get("token")) {
			rw.Write([]byte(req.FormValue("echostr")))
		}
	} else {
		var wxMsg wxUtil.WxMsgReceived
		wxMsg.Set(string(bytes))
		dealTextMsg(&wxMsg, &rw)
	}

}

func dealTextMsg(wxTextMsg *wxUtil.WxMsgReceived, rw *http.ResponseWriter) {
	http.ResponseWriter(*rw).Write([]byte(wxTextMsg.Content))
}
