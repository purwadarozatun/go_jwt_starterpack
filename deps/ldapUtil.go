package deps

import (
	"{project_package}/dto"

	"github.com/jtblin/go-ldap-client"
)

func AuthenticateLdapUser(loginRequest dto.LoginRequest) (bool, map[string]string, error) {

	config, err := LoadConfig(".")
	if err != nil {
		return false, nil, err
	}
	client := &ldap.LDAPClient{
		Base:         config.LDAP_BASE_DN,
		Host:         config.LDAP_URL,
		Port:         config.LDAP_PORT,
		UseSSL:       false,
		BindDN:       config.LDAP_BIND_DN,
		BindPassword: config.LDAP_BIND_PASSWORD,
		UserFilter:   "(uid=%s)",
		GroupFilter:  "(memberUid=%s)",
		Attributes:   []string{"givenName", "sn", "mail", "uid"},
	}
	// It is the responsibility of the caller to close the connection
	defer client.Close()

	return client.Authenticate(loginRequest.Username, loginRequest.Password)
}
