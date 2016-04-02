package mp

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"strings"
)

var template = `
%s (%s)
昨日收盘价:%s
今日开盘价:%s
今日最高价:%s
今日最低价:%s
成交量:%s手
成交额:%s元`

func Sh(stockCode string) string {
	info := SinaStock(stockCode)
	return fmt.Sprintf(template,
		info.StockName,
		info.StockCode,
		info.YesterdayEnd,
		info.DayStart,
		info.DayHigh,
		info.DayDown,
		info.DayChange,
		info.DayTotal)
}

func SinaStock(stockCode string) *StockInfo {
	info := new(StockInfo)
	sh := "http://hq.sinajs.cn/rn=1459529465729&list=" + stockCode
	response, err := http.Get(sh)
	if err != nil {
		return info
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return info
	}
	gbk := mahonia.NewDecoder("GBK")
	_, cdata, _ := gbk.Translate(result, false)
	s := string(cdata)
	start := strings.Index(s, "\"")
	end := strings.LastIndex(s, "\"")
	infoStr := s[start+1 : end]
	infoArray := strings.Split(infoStr, ",")
	info.StockName = infoArray[0]
	info.StockCode = stockCode
	info.DayStart = infoArray[1]
	info.YesterdayEnd = infoArray[2]
	info.DayHigh = infoArray[4]
	info.DayDown = infoArray[5]
	info.DayChange = infoArray[8]
	info.DayTotal = infoArray[9]
	info.Date = infoArray[30]
	info.Time = infoArray[31]
	return info
}

type StockInfo struct {
	StockName    string
	StockCode    string
	Date         string
	Time         string
	DayStart     string
	YesterdayEnd string
	DayEnd       string
	DayHigh      string
	DayDown      string
	DayChange    string
	DayTotal     string
}
