package tenant

import (
	"strings"
	"sync"
)

var tenantMap map[string]string = map[string]string{"dev": "dfc3ccb1f", "qa": "d8b3bfb89", "prod": "d9cfb42fa", "sandbox": "d6c83d68e"}

func Get(tenant string) string {
	return tenantMap[strings.ToLower(tenant)]
}

func Advance(tenant string) string {
	tenant = strings.ToLower(tenant)
	switch tenant {
	case "sandbox":
		return "dev"
	case "dev":
		return "qa"
	case "qa":
		return "prod"
	case "prod":
		return "sandbox"
	}
	return ""
}

type Lock struct {
	Map map[string]*sync.Mutex
}

func InitializeTenantLocks() Lock {
	lockMap := make(map[string]*sync.Mutex)
	for k, _ := range(tenantMap) {
		tmp := sync.Mutex{}
		lockMap[k] = &tmp
	}	 
	return Lock{lockMap}
}
