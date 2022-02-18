package server

import (
	"crypto/tls"
	"example-backend/config"
	"log"

	"github.com/go-ldap/ldap/v3"
)

var LdapConn *ldap.Conn

func GetLDAPConn() (*ldap.Conn, error) {
	if LdapConn != nil && !LdapConn.IsClosing() {
		return LdapConn, nil
	}
	LdapConn, _ = ldap.DialURL(config.Cfg.Ldap.Addr)
	err := LdapConn.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Println(err)
	}
	err = LdapConn.Bind(config.Cfg.Ldap.UserName, config.Cfg.Ldap.Password)
	if err != nil {
		return nil, err
	}
	return LdapConn, nil
}
