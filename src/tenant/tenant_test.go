package tenant

import (
	"testing"
)

func TestInitializeTenantLockMap(t *testing.T) {
	Locks := InitializeTenantLocks()
	lock := Locks.Map["sandbox"]
	(*lock).Lock()
	(*lock).Unlock()
}
