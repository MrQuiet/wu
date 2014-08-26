package wu

import (
	"github.com/mattn/go-ole"
	"github.com/mattn/go-ole/oleutil"
)

type Session struct {
	updateSession *ole.IDispatch
}

func NewSession() *Session {
	ses := new(Session)
	if unknown, err := oleutil.CreateObject("Microsoft.Update.Session"); err != nil {
		panic(err)
	} else {
		updateSession, _ := unknown.QueryInterface(ole.IID_IDispatch)
		ses.updateSession = updateSession
	}
	return ses
}
