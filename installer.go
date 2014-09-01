package wu

import (
	"github.com/mattn/go-ole"
	"github.com/mattn/go-ole/oleutil"
)

type Installer struct {
	updateInstaller *ole.IDispatch
}

func (ses *Session) NewInstaller(up *Updates) *Installer {
	ins := new(Installer)
	ins.updateInstaller = oleutil.MustCallMethod(ses.updateSession, "CreateUpdateInstaller").ToIDispatch()
	if up != nil {
		oleutil.MustPutProperty(ins.updateInstaller, "Updates", up.updates)
	}
	oleutil.MustPutProperty(ins.updateInstaller, "ForceQuiet", true) // FIXME: probably not always true...
	return ins
}

func (ins *Installer) IsBusy() bool {
	return oleutil.MustGetProperty(ins.updateInstaller, "IsBusy").Val == 1
}

func (ins *Installer) Install() (WUError, error) {
	ret, err := oleutil.CallMethod(ins.updateInstaller, "Install")
	val := WUError(ret.Val)
	return val, err
}
