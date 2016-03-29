package wxUtil

import (
	"crypto/sha1"
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
