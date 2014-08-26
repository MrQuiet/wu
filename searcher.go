package wu

import (
	"fmt"
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

func (sea *Searcher) Query(query string) {
	fmt.Printf("Query: %s\n", query)
	queryResult := oleutil.MustCallMethod(sea.updateSearcher, "Search", query).ToIDispatch()
	updates := oleutil.MustGetProperty(queryResult, "Updates").ToIDispatch()
	count := int(oleutil.MustGetProperty(updates, "Count").Val)
	fmt.Printf("Query result count: %d\n", count)
	for i := 0; i < count; i++ {
		item := oleutil.MustGetProperty(updates, "Item", i).ToIDispatch()
		title := oleutil.MustGetProperty(item, "Title").ToString()
		fmt.Printf("[%d] %s\n", i, title)
	}
}
