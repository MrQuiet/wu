package wu

import (
	"testing"
)

func Test_Installer(t *testing.T) {
	NewSession().NewInstaller() // cause a crash?
}

func Test_InstallerIsBusy(t *testing.T) {
	t.Logf("Installer IsBusy?  %v", NewSession().NewInstaller().IsBusy())
}
