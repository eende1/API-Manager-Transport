package github

import (
	"testing"
)

func TestInitializeGithubRepo(t *testing.T) {
	_, err := initializeGithubRepo()

	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
}



