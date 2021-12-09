package rbac

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"os"
)

var ROLES map[string]bool
var authPaths map[string][]map[string]interface{}
var authApis map[string][]string

func setupAuth(authFile string) error {
	file, err := os.Open(authFile)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	data, err := simplejson.NewJson(bytes)
	if err != nil {
		return err
	}

	roles, err := data.Get("roles").StringArray()
	if err != nil {
		return err
	}

	authPaths = make(map[string][]map[string]interface{})
	authApis = make(map[string][]string)
	ROLES = make(map[string]bool)
	for _, role := range roles {
		ROLES[role] = true
		auth := data.Get(role).Interface()
		path, apis, err := parserData(auth.([]interface{}))
		if err != nil {
			return err
		}
		authPaths[role] = path
		authApis[role] = apis
	}

	return nil
}

func parserData(items []interface{}) ([]map[string]interface{}, []string, error) {
	data := make([]map[string]interface{}, 0)
	var apis []string
	for _, item := range items {
		info := make(map[string]interface{})
		apisInf := item.(map[string]interface{})["api"]
		if apisInf != nil {
			apisInfs := item.(map[string]interface{})["api"].([]interface{})
			for _, apisInfItem := range apisInfs {
				apis = append(apis, apisInfItem.(string))
			}
		}
		pathInf := item.(map[string]interface{})["path"]
		if pathInf != nil {
			info["path"] = pathInf.(string)
		} else {
			return data, nil, fmt.Errorf("%s", "invalid format, not found path")
		}

		childrenInf := item.(map[string]interface{})["children"]
		if childrenInf != nil {
			childPath, childApis, err := parserData(childrenInf.([]interface{}))
			if err != nil {
				return data, apis, err
			}
			info["children"] = childPath
			apis = append(apis, childApis...)
		}
		data = append(data, info)
	}

	return data, apis, nil
}
