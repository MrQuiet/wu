package wu

import (
	"testing"
)

func Test_SystemInfo(t *testing.T) {
	NewSystemInfo() // cause a crash?
}

func Test_RebootRequired(t *testing.T) {
	t.Logf("Reboot required?  %v", NewSystemInfo().RebootRequired())
}
