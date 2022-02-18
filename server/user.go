package server

import (
	"context"
	"fmt"
	"net/http"
	"tats-backend/api"
	"tats-backend/config"
	"time"

	"github.com/go-ldap/ldap/v3"
)

func LdapSearchUser(name, password string) (bool, error) {
	conn, err := GetLDAPConn()
	if err != nil {
		return false, err
	}

	if err := conn.Bind(config.Cfg.Ldap.UserName, config.Cfg.Ldap.Password); err != nil {
		return false, err
	}
	searchRequest := ldap.NewSearchRequest("ou=People,dc=domain,dc=sensetime,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 10, 10, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(cn=%s))", name),
		[]string{"dn"}, nil)
	sr, err := conn.Search(searchRequest)
	if err != nil {
		return false, fmt.Errorf(err.Error())
	}
	if len(sr.Entries) != 1 {
		return false, fmt.Errorf("user does not exist or too many entries returned")
	}

	userDN := sr.Entries[0].DN
	time.Sleep(time.Second)

	//Bind as the user to verify their password
	err = conn.Bind(userDN, password)
	if err != nil {
		return false, fmt.Errorf(err.Error())
	}
	return true, nil
}

func HttpLogin(w http.ResponseWriter, r *http.Request) {

	succ(w, "")
}

func HttpLogout(w http.ResponseWriter, r *http.Request) {

	succ(w, nil)
}

func (s *Ser) CreateUser(context.Context, *api.CreateUserReq) (*api.CreateUserResp, error) {
	return nil, nil
}
func (s *Ser) UpdateUser(context.Context, *api.UpdateUserReq) (*api.UpdateUserResp, error) {
	return nil, nil
}
func (s *Ser) DeleteUser(context.Context, *api.DeleteUserReq) (*api.SuccessResp, error) {
	return nil, nil
}
func (s *Ser) Login(context.Context, *api.LoginReq) (*api.LoginResp, error) {
	return nil, nil
}
func (s *Ser) Logout(context.Context, *api.LogoutReq) (*api.SuccessResp, error) {
	return nil, nil
}
func (s *Ser) ListUsers(context.Context, *api.ListUsersReq) (*api.ListUsersResp, error) {
	return nil, nil
}
