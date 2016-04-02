package mp

import (
	"github.com/chentongming/yszcc/application/service/mp"
	"testing"
)

func TestStock(t *testing.T) {
	info := mp.SinaStock("sz399001")
	if info == nil {
		t.Error("msg err")
	}
	if info.StockCode != "sz399001" {
		t.Error("msg code error")
	}
}
