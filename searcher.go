package wu

import (
	"github.com/mattn/go-ole"
	"github.com/mattn/go-ole/oleutil"
)

type Searcher struct {
	updateSearcher *ole.IDispatch
}

func (ses *Session) NewSearcher(online bool) *Searcher {
	sea := new(Searcher)
	sea.updateSearcher = oleutil.MustCallMethod(ses.updateSession, "CreateUpdateSearcher").ToIDispatch()
	oleutil.MustPutProperty(sea.updateSearcher, "Online", online)
	return sea
}
