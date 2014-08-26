package wu

import (
	"github.com/mattn/go-ole"
	"github.com/mattn/go-ole/oleutil"
)

type SystemInfo struct {
	updateSystemInfo *ole.IDispatch
}

func NewSystemInfo() *SystemInfo {
	sys := new(SystemInfo)
	if unknown, err := oleutil.CreateObject("Microsoft.Update.SystemInfo"); err != nil {
		panic(err)
	} else {
		updateSystemInfo, _ := unknown.QueryInterface(ole.IID_IDispatch)
		sys.updateSystemInfo = updateSystemInfo
	}
	return sys
}

func (sys *SystemInfo) RebootRequired() bool {
	return oleutil.MustGetProperty(sys.updateSystemInfo, "RebootRequired").Val == 1
}
