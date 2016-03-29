package wxUtil

import (
	"encoding/xml"
)

type WxResponseBase struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
}

type WxResponseText struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
	Content      string
}

func (w *WxResponseText) ToXml() string {
	bytes, err := xml.Marshal(w)
	if err != nil {
		return ""
	}
	bytes = append([]byte(xml.Header), bytes...)
	return string(bytes)
}
