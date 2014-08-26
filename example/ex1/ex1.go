package main

import (
	"gopkg.in/wu.v0"
)

func main() {
	ses := wu.NewSession()
	_ = ses.NewSearcher(true)
}
