package mp

import (
	"github.com/chentongming/yszcc/application/service/mp"
	"testing"
)

func TestStock(t *testing.T) {
	info := mp.Sh()
	t.Error(info)
}
