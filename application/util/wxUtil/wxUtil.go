package wxUtil

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/chentongming/yszcc/application/util/requestUtil"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func GetAccessToken(appId string, appSecret string) map[string]interface{} {
	params := url.Values{}
	params.Add("grant_type", "client_credential")
	params.Add("appid", appId)
	params.Add("secret", appSecret)
	result, err := requestUtil.Get("https://api.weixin.qq.com/cgi-bin/token", params)
	if err != nil {
		return nil
	}
	return result
}

type WxServerValidData struct {
	Signature string
	Timestamp int
	Nonce     string
	Echostr   string
}

func (validData *WxServerValidData) Calc(token string) string {
	data := []string{token, strconv.Itoa(validData.Timestamp), validData.Nonce}
	sort.Strings(data)
	sha1Result := sha1.Sum([]byte(strings.Join(data, "")))
	return fmt.Sprintf("%x", sha1Result)

}

func (validData *WxServerValidData) Valid(token string) bool {
	if validData.Calc(token) == validData.Signature {
		return true
	} else {
		return false
	}
}

type WxMsgReceived struct {
	ToUserName   string
	FromUserName string
	CreateTime   int32
	MsgType      string
	MsgId        int64
	Content      string //文本
	Event        string //事件
	MediaId      string //视频，图片，语音 下载id
	ThumbMediaId string //视频消息缩略图
	Recognition  string //语音文本
	Format       string //语音格式
	Location_X   string //地理位置
	Location_Y   string
	Scale        string
	Label        string
	Title        string
	Description  string
	Url          string
}

func (wxMsgReceived *WxMsgReceived) Set(xmlStr string) error {
	err := xml.Unmarshal([]byte(xmlStr), wxMsgReceived)
	return err
}
