package wu

import (
	"github.com/mattn/go-ole"
	"github.com/mattn/go-ole/oleutil"
)

type Installer struct {
	updateInstaller *ole.IDispatch
}

func (ses *Session) NewInstaller() *Installer {
	ins := new(Installer)
	ins.updateInstaller = oleutil.MustCallMethod(ses.updateSession, "CreateUpdateInstaller").ToIDispatch()
	return ins
}

func (ins *Installer) IsBusy() bool {
	return oleutil.MustGetProperty(ins.updateInstaller, "IsBusy").Val == 1
}
