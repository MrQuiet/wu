package wu

import (
	"github.com/mattn/go-ole"
	"github.com/mattn/go-ole/oleutil"
)

type Updates struct {
	updates *ole.IDispatch // ??? Not sure why we're hanging on to this.
	Updates []*Update
}

type Update struct {
	item            *ole.IDispatch
	Title           string
	Description     string
	IsInstalled     bool
	IsUninstallable bool
	IsMandatory     bool
	EulaAccepted    bool
	RevisionNumber  int
	UpdateID        string
	SupportUrl      string
}

func newUpdate(item *ole.IDispatch) *Update {
	up := new(Update)
	up.item = item
	up.Title = up.GetString("Title")
	up.Description = up.GetString("Description")
	up.IsInstalled = up.GetBool("IsInstalled")
	up.IsUninstallable = up.GetBool("IsUninstallable")
	up.IsMandatory = up.GetBool("IsMandatory")
	up.EulaAccepted = up.GetBool("EulaAccepted")
	identity := oleutil.MustGetProperty(item, "Identity").ToIDispatch()
	revisionNumber := int(oleutil.MustGetProperty(identity, "RevisionNumber").Val)
	up.RevisionNumber = revisionNumber
	updateID := oleutil.MustGetProperty(identity, "UpdateID").ToString()
	up.UpdateID = updateID
	up.SupportUrl = up.GetString("SupportUrl")
	return up
}

func (up *Update) AcceptEula(accept bool) {
	oleutil.MustCallMethod(up.item, "AcceptEula", accept)
}

func (up *Update) GetString(attr string) string {
	return oleutil.MustGetProperty(up.item, attr).ToString()
}

func (up *Update) GetBool(attr string) bool {
	return oleutil.MustGetProperty(up.item, attr).Value().(bool)
}

func (up *Update) GetInt(attr string) int {
	return int(oleutil.MustGetProperty(up.item, attr).Val)
}
