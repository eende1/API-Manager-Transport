package tenant

import "strings"

var tenantMap map[string]string = map[string]string{"dev": "dfc3ccb1f", "qa": "d8b3bfb89", "prod": "d9cfb42fa"}

func Get(tenant string) string {
	return tenantMap[strings.ToLower(tenant)]
}
