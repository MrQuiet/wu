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

func (sea *Searcher) Query(query string) *Updates {
	queryResult := oleutil.MustCallMethod(sea.updateSearcher, "Search", query).ToIDispatch()
	updatesSet := new(Updates)
	updates := oleutil.MustGetProperty(queryResult, "Updates").ToIDispatch()
	count := int(oleutil.MustGetProperty(updates, "Count").Val)
	updatesSet.Updates = make([]*Update, count)
	for i := 0; i < count; i++ {
		item := oleutil.MustGetProperty(updates, "Item", i).ToIDispatch()
		updatesSet.Updates[i] = newUpdate(item)
	}
	return updatesSet
}
