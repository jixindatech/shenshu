package service

import (
	"admin/server/models"
)

type Ldap struct {
	ID uint

	Type     string
	Host     string
	Port     int
	DN       string
	BaseDN   string
	Password string
}

var ldapCache *Ldap

func (l *Ldap) Save() error {
	data := make(map[string]interface{})
	data["type"] = l.Type
	data["host"] = l.Host
	data["port"] = l.Port
	data["dn"] = l.DN
	data["basedn"] = l.BaseDN
	data["password"] = l.Password

	if l.ID > 0 {
		ldapCache = nil
		return models.UpdateLdap(l.ID, data)
	}

	return models.AddLdap(data)
}

func (l *Ldap) Get() (*models.Ldap, error) {
	return models.GetLdap()
}

func getLdapConfig() (*Ldap, error) {
	if ldapCache == nil {
		ldapCache = new(Ldap)
		ldap, err := ldapCache.Get()
		if err != nil {
			return nil, err
		}
		ldapCache.Type = ldap.Type
		ldapCache.Host = ldap.Host
		ldapCache.Port = ldap.Port
		ldapCache.BaseDN = ldap.BaseDN
		ldapCache.DN = ldap.DN
		ldapCache.Password = ldap.Password
		return ldapCache, nil
	}

	return ldapCache, nil
}
