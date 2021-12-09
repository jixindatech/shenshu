package rbac

import (
	"github.com/casbin/casbin/v2"
)

var enforcer *casbin.Enforcer

func setupCasbin(model, policy string) error {
	var err error
	enforcer, err = casbin.NewEnforcer(model, policy)
	if err != nil {
		return err
	}

	enforcer.AddFunction("role", roleMatchFunc)

	enforcer.EnableLog(true)

	/*
		ok, err := enforcer.Enforce("admin", "/user/abc", "GET")
		if err != nil {
			return err
		}
		return fmt.Errorf("true: ", ok)
	*/

	return nil
}

func roleMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(roleMatch(name1, name2)), nil
}

func roleMatch(key1 string, key2 string) bool {
	return key2 == "*" || key1 == key2
}
