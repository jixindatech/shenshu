package rbac

import (
	"admin/config"
	"fmt"
	"github.com/casbin/casbin/v2"
)

func Setup(cfg *config.Rbac) error {
	var err error

	err = setupCasbin(cfg.Model, cfg.Policy)
	if err != nil {
		return fmt.Errorf("casbin error: %s", err)
	}

	err = setupAuth(cfg.Auth)
	if err != nil {
		return fmt.Errorf("casbin auth error: %s", err)
	}

	return nil
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}

func GetRoleRoutes(role string) []map[string]interface{} {
	if authPaths != nil && authPaths[role] != nil {
		return authPaths[role]
	}

	return nil
}

func GetRoleApi(role string) []string {
	if authApis != nil && authApis[role] != nil {
		return authApis[role]
	}

	return nil
}
