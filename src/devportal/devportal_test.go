package devportal

import (
	"testing"
)

func TestCreateDevPortalProject(t *testing.T) {
	err := createDevPortalProject("test4trace", "description4trace")
	if err != nil {
		t.Errorf("failed with error: %s", err)
	}

	// Calling this twice should work
	err = createDevPortalProject("test4trace", "description4trace")
	if err != nil {
		t.Errorf("failed with error: %s", err)
	}
}
