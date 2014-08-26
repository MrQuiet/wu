package main

import (
	"fmt"
	"gopkg.in/wu.v0"
)

func main() {
	ses := wu.NewSession()
	sea := ses.NewSearcher(true)
	updates := sea.Query(wu.NewSearchQuery().IsInstalled(true).IsHidden(false).String())
	fmt.Printf("Updates: %d\n", len(updates.Updates))
	for i := 0; i < len(updates.Updates); i++ {
		update := updates.Updates[i]
		fmt.Printf("[%s] [%d] %s\nIsUninstallable[%t]\nUpdate URL[%s]\n\n", update.UpdateID, update.RevisionNumber, update.Title, update.IsUninstallable, update.SupportUrl)
	}
}
