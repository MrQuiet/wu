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
	item                *ole.IDispatch
	Title               string
	Description         string
	CanRequireSource    bool
	IsBeta              bool
	IsDownloaded        bool
	IsInstalled         bool
	IsUninstallable     bool
	IsHidden            bool
	IsMandatory         bool
	RebootRequired      bool
	EulaAccepted        bool
	RevisionNumber      int
	UpdateID            string
	SupportUrl          string
	SecurityBulletinIDs []string
	SupersededUpdateIDs []string
}

func newUpdate(item *ole.IDispatch) *Update {
	up := new(Update)
	up.item = item
	up.Title = up.GetString("Title")
	up.Description = up.GetString("Description")
	up.CanRequireSource = up.GetBool("CanRequireSource")
	up.IsBeta = up.GetBool("IsBeta")
	up.IsDownloaded = up.GetBool("IsDownloaded")
	up.IsInstalled = up.GetBool("IsInstalled")
	up.IsUninstallable = up.GetBool("IsUninstallable")
	up.IsHidden = up.GetBool("IsHidden")
	up.IsMandatory = up.GetBool("IsMandatory")
	up.RebootRequired = up.GetBool("RebootRequired")
	up.EulaAccepted = up.GetBool("EulaAccepted")
	{
		identity := oleutil.MustGetProperty(item, "Identity").ToIDispatch()
		revisionNumber := int(oleutil.MustGetProperty(identity, "RevisionNumber").Val)
		up.RevisionNumber = revisionNumber
		updateID := oleutil.MustGetProperty(identity, "UpdateID").ToString()
		up.UpdateID = updateID
	}
	up.SupportUrl = up.GetString("SupportUrl")
	{
		securityBulletinIDs := oleutil.MustGetProperty(item, "SecurityBulletinIDs").ToIDispatch()
		count := int(oleutil.MustGetProperty(securityBulletinIDs, "Count").Val)
		up.SecurityBulletinIDs = make([]string, count)
		for i := 0; i < count; i++ {
			securityBulletinID := oleutil.MustGetProperty(securityBulletinIDs, "Item", i).ToString()
			up.SecurityBulletinIDs[i] = securityBulletinID
		}
	}
	{
		supersededUpdateIDs := oleutil.MustGetProperty(item, "SupersededUpdateIDs").ToIDispatch()
		count := int(oleutil.MustGetProperty(supersededUpdateIDs, "Count").Val)
		up.SupersededUpdateIDs = make([]string, count)
		for i := 0; i < count; i++ {
			supersededUpdateID := oleutil.MustGetProperty(supersededUpdateIDs, "Item", i).ToString()
			up.SupersededUpdateIDs[i] = supersededUpdateID
		}
	}
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
