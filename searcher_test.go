package wu

import (
	"testing"
)

func Test_NewSearcher(t *testing.T) {
	ses := NewSession()
	ses.NewSearcher(true)  // cause a crash?
	ses.NewSearcher(false) // cause a crash?
}
