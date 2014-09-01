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
	IsPresent                bool
	RebootRequired           bool
	EulaAccepted             bool
	PerUser                  bool
	BrowseOnly               bool
	RevisionNumber           int
	UpdateID                 string
	SupportUrl               string
	MsrcSeverity             string
	Deadline                 time.Time
	LastDeploymentChangeTime time.Time
	SecurityBulletinIDs      []string
	SupersededUpdateIDs      []string
	KBArticleIDs             []string
	CveIDs                   []string
	MaxDownloadSize          int
	MinDownloadSize          int
	Categories               Categories
	RecommendedCPUSpeed      int
	RecommendedHardDiskSpace int
	RecommendedMemory        int
	Type                     UpdateType
	AutoDownload             AutoDownload
	AutoSelection            AutoSelection
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
	up.IsPresent = up.GetBool("IsPresent")
	up.RebootRequired = up.GetBool("RebootRequired")
	up.EulaAccepted = up.GetBool("EulaAccepted")
	up.PerUser = up.GetBool("PerUser")
	up.BrowseOnly = up.GetBool("BrowseOnly")
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
	{
		cveIDs := oleutil.MustGetProperty(item, "CveIDs").ToIDispatch()
		count := int(oleutil.MustGetProperty(cveIDs, "Count").Val)
		up.CveIDs = make([]string, count)
		for i := 0; i < count; i++ {
			cveID := oleutil.MustGetProperty(cveIDs, "Item", i).ToString()
			up.CveIDs[i] = cveID
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
	up.RecommendedCPUSpeed = up.GetInt("RecommendedCPUSpeed")
	up.RecommendedHardDiskSpace = up.GetInt("RecommendedHardDiskSpace")
	up.RecommendedMemory = up.GetInt("RecommendedMemory")
	up.Type = UpdateType(up.GetInt("Type"))
	up.AutoDownload = AutoDownload(up.GetInt("AutoDownload"))
	up.AutoSelection = AutoSelection(up.GetInt("AutoSelection"))
	return up
}

func (up *Update) AcceptEula() (WUError, error) {
	ret, err := oleutil.CallMethod(up.item, "AcceptEula")
	if err == nil {
		up.EulaAccepted = true
	}
	val := WUError(ret.Val)
	return val, err
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
	t := time.Unix(val/10000000-11644473600, 0) // FIXME: I don't know how to interpret this value.
	return t
}

func (up *Update) GetDate(attr string) time.Time {
	val := oleutil.MustGetProperty(up.item, attr).Val
	if val == 0 {
		return time.Time{}
	}
	t := time.Unix(val/10000000-11644473600, 0) // FIXME: I don't know how to interpret this value.
	return t
}
