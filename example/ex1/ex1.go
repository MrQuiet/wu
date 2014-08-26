package main

import (
	"gopkg.in/wu.v0"
)

func main() {
	ses := wu.NewSession()
	sea := ses.NewSearcher(true)
	sea.Query(wu.NewSearchQuery().IsInstalled(true).IsHidden(false).String())
	sea.Query(wu.NewSearchQuery().IsInstalled(false).IsHidden(false).String())
}
