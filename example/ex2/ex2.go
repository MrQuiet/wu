package main

import (
	"fmt"
	"gopkg.in/wu.v0"
)

func main() {
	ses := wu.NewSession()
	sea := ses.NewSearcher(true)
	updates := sea.Query(wu.NewSearchQuery().IsInstalled(false).IsHidden(false).String())
	fmt.Printf("Updates: %d\n", len(updates.Updates))
	for _, update := range updates.Updates {
		fmt.Printf("[%s] %s\n", update.Type, update.Title)
		if !update.EulaAccepted {
			code, err := update.AcceptEula()
			fmt.Printf("ACCEPTEULA: [%s] [%s]\n", code, err)
		}
	}
	downloader := ses.NewDownloader(updates)
	code, err := downloader.Download()
	fmt.Printf("DOWNLOAD: [%s] [%s]\n", code, err)
}
