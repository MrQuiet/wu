package main

import (
	"fmt"
	"gopkg.in/wu.v0"
	"strings"
)

func main() {
	ses := wu.NewSession()
	sea := ses.NewSearcher(false)
	updates := sea.Query(wu.NewSearchQuery().IsInstalled(true).IsHidden(false).String())
	fmt.Printf("Updates: %d\n", len(updates.Updates))
	for _, update := range updates.Updates {
		securityBulletinIDs := "\t" + strings.Join(update.SecurityBulletinIDs, "\n\t")
		supersededUpdateIDs := "\t" + strings.Join(update.SupersededUpdateIDs, "\n\t")
		fmt.Printf("[%s]\n%s\nIsUninstallable[%t]\nRebootRequired[%t]\nUpdate URL[%s]\nSupersededUpdateIDs:\n%s\nSecurityBulletinIDs:\n%s\n\n", update.UpdateID, update.Title, update.IsUninstallable, update.RebootRequired, update.SupportUrl, supersededUpdateIDs, securityBulletinIDs)
	}
}
