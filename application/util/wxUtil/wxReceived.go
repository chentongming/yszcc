package wxUtil

import (
	"encoding/xml"
)

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
