package plutopure

import "testing"

func TestVersion(t *testing.T) {
	t.Logf("version num: %d", VersionNumber())
}

func TestVersionString(t *testing.T) {
	t.Logf("version string: %s", Version())
}
