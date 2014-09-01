package wu

import (
	"github.com/mattn/go-ole"
	"github.com/mattn/go-ole/oleutil"
)

type Downloader struct {
	updateDownloader *ole.IDispatch
}

func (ses *Session) NewDownloader(up *Updates) *Downloader {
	dow := new(Downloader)
	dow.updateDownloader = oleutil.MustCallMethod(ses.updateSession, "CreateUpdateDownloader").ToIDispatch()
	if up != nil {
		oleutil.MustPutProperty(dow.updateDownloader, "Updates", &up.updates)
	}
	return dow
}

func (dow *Downloader) IsBusy() bool {
	return oleutil.MustGetProperty(dow.updateDownloader, "IsBusy").Val == 1
}

// FIXME: return value is probably going to be set to something useful
func (dow *Downloader) Download() (code int64, err error) {
	ret, err := oleutil.CallMethod(dow.updateDownloader, "Download")
	val := ret.Val
	return val, err
}
