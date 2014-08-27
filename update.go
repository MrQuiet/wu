package wu

import (
	"github.com/mattn/go-ole"
	"github.com/mattn/go-ole/oleutil"
	"time"
)

type Updates struct {
	updates *ole.IDispatch // ??? Not sure why we're hanging on to this.
	Updates []*Update
}

type Update struct {
	item                     *ole.IDispatch
	Title                    string
	Description              string
	CanRequireSource         bool
	IsBeta                   bool
	IsDownloaded             bool
	IsInstalled              bool
	IsUninstallable          bool
	IsHidden                 bool
	IsMandatory              bool
	RebootRequired           bool
	EulaAccepted             bool
	RevisionNumber           int
	UpdateID                 string
	SupportUrl               string
	MsrcSeverity             string
	Deadline                 time.Time
	LastDeploymentChangeTime time.Time
	SecurityBulletinIDs      []string
	SupersededUpdateIDs      []string
	KBArticleIDs             []string
	MaxDownloadSize          int
	MinDownloadSize          int
	Categories               Categories
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
	up.MsrcSeverity = up.GetString("MsrcSeverity")
	up.Deadline = up.GetDate("Deadline")
	up.LastDeploymentChangeTime = up.GetDateTime("LastDeploymentChangeTime")
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
	{
		kbArticleIDs := oleutil.MustGetProperty(item, "KBArticleIDs").ToIDispatch()
		count := int(oleutil.MustGetProperty(kbArticleIDs, "Count").Val)
		up.KBArticleIDs = make([]string, count)
		for i := 0; i < count; i++ {
			kbArticleID := oleutil.MustGetProperty(kbArticleIDs, "Item", i).ToString()
			up.KBArticleIDs[i] = kbArticleID
		}
	}
	up.MaxDownloadSize = up.GetInt("MaxDownloadSize")
	up.MinDownloadSize = up.GetInt("MinDownloadSize")
	{
		up.Categories = Categories{}
		categories := oleutil.MustGetProperty(item, "Categories").ToIDispatch()
		count := int(oleutil.MustGetProperty(categories, "Count").Val)
		up.Categories.Categories = make([]*Category, count)
		for i := 0; i < count; i++ {
			category := oleutil.MustGetProperty(categories, "Item", i).ToIDispatch()
			up.Categories.Categories[i] = newCategory(category)
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

func (up *Update) GetInt64(attr string) int64 {
	return oleutil.MustGetProperty(up.item, attr).Val
}

func (up *Update) GetDateTime(attr string) time.Time {
	val := oleutil.MustGetProperty(up.item, attr).Val
	if val == 0 {
		return time.Time{}
	}
	t := time.Unix(val/10000000-11644473600, 0) // FIXME: This is supposed to be a Windows DateTime, but I'm not having a good time on this one.
	return t
}

func (up *Update) GetDate(attr string) time.Time {
	val := oleutil.MustGetProperty(up.item, attr).Val
	if val == 0 {
		return time.Time{}
	}
	t := time.Unix(val/10000000-11644473600, 0) // FIXME: This is supposed to be a Windows Date, not an attempt at a DateTime.
	return t
}
