package wu

import (
	"github.com/mattn/go-ole"
)

func init() {
	ole.CoInitialize(0)
}
