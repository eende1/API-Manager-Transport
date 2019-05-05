package okta

import (
	"os"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	secret := os.Getenv("UNAUTHORIZEDID_SECRET")
	if secret == "" {
		t.Error("UNAUTHORIZEDID_SECRET not in environment")
	}
	token, err := GenerateToken("nike.sapae.unauthorizedid", secret, false)
	t.Log(token)
	if err != nil {
		t.Errorf("returned an error: %s", err)
	}
}
