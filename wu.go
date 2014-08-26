package wu

import (
	"github.com/mattn/go-ole"
)

func init() {
	ole.CoInitialize(0)
}

func btou(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
