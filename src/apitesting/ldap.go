package apitesting

import (
	"gopkg.in/ldap.v3"
	"fmt"
)

func LDAPAuthenticationTest(c chan TestResult, email, password, name string) {
	result, err := LDAPAuthorization(email, password)
	c <- TestResult{name, result, err}
}


func LDAPAuthorization(email, password string) (bool, error) {
	l, err := ldap.Dial("tcp", "ldap.nike.com:389")
	if err != nil {
		return false, err
	}
	err = l.Bind(email, password)
	if err != nil {
		return false, err
	}
	
	defer l.Close()
	searchRequest := ldap.NewSearchRequest(
		"DC=ad,DC=nike,DC=com", // The base dn to search
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 20, 20, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(userPrincipalName=%s)(memberOf=CN=Lst-ERP.DDA.Portal&Integration,OU=Lists,OU=BEAVERTN,OU=OR,OU=USA,DC=ad,DC=nike,DC=com))", email),
		[]string{"memberOf","co"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return false, nil
	}
	
	if (len(sr.Entries) < 1) {
		return false, nil
	}
	
	return true, nil
}
