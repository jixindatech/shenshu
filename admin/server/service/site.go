package service

import (
	"admin/core/log"
	"admin/server/cache"
	"admin/server/models"
	"admin/server/util"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"time"
)

const cacheSiteName string = "site"

type Site struct {
	ID uint

	Name        string
	Host        string
	Path        string
	Status      int
	UpstreamRef uint
	Remark      string

	Type int
	Ids  []uint

	Page     int
	PageSize int
}

func (r *Site) Save() (err error) {
	data := make(map[string]interface{})
	data["name"] = r.Name
	data["host"] = r.Host
	data["path"] = r.Path
	data["status"] = r.Status
	data["upstreamRef"] = r.UpstreamRef
	data["remark"] = r.Remark

	if r.ID > 0 {
		err = models.UpdateSite(r.ID, data)
	} else {
		err = models.AddSite(data)
	}

	if err != nil {
		return err
	}

	return SetupSites()
}

func (r *Site) Get() (*models.Site, error) {
	return models.GetSite(r.ID)
}

func (r *Site) GetList() ([]*models.Site, int, error) {
	data := make(map[string]interface{})
	data["name"] = r.Name
	data["page"] = r.Page
	data["pagesize"] = r.PageSize
	data["status"] = r.Status

	return models.GetSites(data)
}

func (r *Site) Delete() error {
	err := models.DeleteSite(r.ID)
	if err != nil {
		return err
	}

	return SetupSites()
}

func (r *Site) UpdatRuleGroup() error {
	var err error
	if r.Type == util.TYPE_BATCH_GROUP {
		err = models.UpdateSiteBatchRuleGroup(r.ID, r.Ids)
	} else if r.Type == util.TYPE_SPECIFIC_GROUP {
		err = models.UpdateSiteSpecificRuleGroup(r.ID, r.Ids)
	}

	return err
}

func (r *Site) GetRuleGroup() ([]uint, error) {
	var ids []uint
	if r.Type == util.TYPE_BATCH_GROUP {
		rulegroups, err := models.GetSiteBatchGroup(r.ID)
		if err != nil {
			return nil, err
		}

		for _, item := range rulegroups {
			ids = append(ids, item.ID)
		}

	} else if r.Type == util.TYPE_SPECIFIC_GROUP {
		rulegroups, err := models.GetSiteSpecificGroup(r.ID)
		if err != nil {
			return nil, err
		}

		for _, item := range rulegroups {
			ids = append(ids, item.ID)
		}

	}

	return ids, nil
}

/*
{
	"host": "*.test.com",
	"path": "/*",
	"upstream_id": 1,
	"ip" : {
		"accept": [],
		"deny": []
	},
	"cc": {
		[
			"path": "/login1",
			"period": 1,
			"time": 1
		],
		[
			"path": "/login2",
			"period": 1,
			"time": 1
		],
	},
	"rule": {
		"config": {
			"action": deny,
			"content-type": ["json", "xml"]
		},
		"rules": {
			"batch": [1, 2, 3],
			"speicific": [4, 5]
		}
	}
}
*/

func getIPConfig(id uint) (map[string]interface{}, error) {
	ipSrv := IP{
		Site:     id,
		Type:     util.IP_ACCEPT,
		Page:     0,
		PageSize: 0,
	}

	list, _, err := ipSrv.GetList()
	if err != nil {
		return nil, err
	}

	var ipsAccept []string
	for _, ip := range list {
		var tmpIPs []string
		err := json.Unmarshal(ip.IP, &tmpIPs)
		if err != nil {
			return nil, err
		}
		ipsAccept = append(ipsAccept, tmpIPs...)
	}

	ipSrv.Type = util.IP_DENY
	list, _, err = ipSrv.GetList()
	if err != nil {
		return nil, err
	}

	var ipsDeny []string
	for _, ip := range list {
		var tmpIPs []string
		err := json.Unmarshal(ip.IP, &tmpIPs)
		if err != nil {
			return nil, err
		}
		ipsDeny = append(ipsDeny, tmpIPs...)
	}

	data := make(map[string]interface{})
	data["accept"] = ipsAccept
	data["deny"] = ipsDeny

	return data, nil
}

func getCCsConfig(id uint) ([]map[string]interface{}, error) {
	var err error
	ccSrv := CC{
		Site:     id,
		Page:     0,
		PageSize: 0,
	}
	list, _, err := ccSrv.GetList()
	if err != nil {
		return nil, err
	}

	var data []map[string]interface{}
	for _, item := range list {
		tmp := make(map[string]interface{})
		tmp["uri"] = item.URI
		tmp["match"] = item.Match
		tmp["mode"] = item.Mode
		tmp["method"] = item.Method
		tmp["threshold"] = item.Threshold
		tmp["duration"] = item.Duration
		tmp["action"] = item.Action

		data = append(data, tmp)
	}

	return data, err
}

func setCache(key string, data interface{}) error {
	siteStr, err := json.Marshal(data)
	if err != nil {
		log.Logger.Error("site", zap.String("err", err.Error()))
		return err
	}

	err = cache.Set(cache.CONFIG, key, string(siteStr), 0)
	if err != nil {
		log.Logger.Error("site", zap.String("err", err.Error()))
		return err
	}

	return nil
}

func getIPData(id uint) (map[string]interface{}, error) {
	res, err := getIPConfig(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func setupIPs(key string, ids []uint) error {
	var data []interface{}
	for _, id := range ids {
		ipConfig, err := getIPData(id)
		if err != nil {
			return err
		}

		item := make(map[string]interface{})
		item["id"] = id
		item["config"] = ipConfig

		data = append(data, item)
	}

	if len(data) == 0 {
		data = []interface{}{}
	}

	return setCache(key, data)
}

func setupCCs(key string, ids []uint) error {
	var data []interface{}
	for _, id := range ids {
		ccConfig, err := getCCsConfig(id)
		if err != nil {
			return err
		}

		item := make(map[string]interface{})
		item["id"] = id
		item["config"] = ccConfig
		data = append(data, item)
	}

	if len(data) == 0 {
		data = []interface{}{}
	}

	return setCache(key, data)
}

func (r *Site) Enable() error {
	return SetupSites()
}

func setupBatchRules(key string) error {
	ruleBatchSrv := &RuleBatch{
		Status: util.RULE_ENABLE,
	}
	batchList, _, err := ruleBatchSrv.GetList()
	if err != nil {
		return err
	}

	err = setCache(key, batchList)
	if err != nil {
		return err
	}

	return nil
}

func setupSpecificRules(key string) error {
	ruleSpecifcSrv := &RuleSpecific{
		Status: util.RULE_ENABLE,
	}
	specificList, _, err := ruleSpecifcSrv.GetList()
	if err != nil {
		return err
	}

	err = setCache(key, specificList)
	if err != nil {
		return err
	}

	return nil
}

func getRuleConfig(id uint) (map[string]interface{}, error) {
	var err error
	decoders := make(map[string]interface{})
	batchgroup := BatchGroup{
		ID:       id,
		Status:   util.RULE_ENABLE,
		Page:     0,
		PageSize: 0,
	}
	batchList, _, err := batchgroup.GetList()
	if err != nil {
		return nil, err
	}

	action := 255
	var batch []uint
	var specific []uint
	for _, item := range batchList {
		if action < item.Action {
			action = item.Action
		}

		var tmpDecoders []string
		err = json.Unmarshal(item.Decoder, &tmpDecoders)
		if err != nil {
			return nil, err
		}
		for _, decoder := range tmpDecoders {
			decoders[decoder] = true
		}

		ruleSrv := RuleBatch{
			RuleGroup: item.ID,
			Status:    util.RULE_ENABLE,
			Page:      0,
			PageSize:  0,
		}
		rules, _, err := ruleSrv.GetList()
		if err != nil {
			return nil, err
		}

		var ids []uint
		for _, rule := range rules {
			ids = append(ids, rule.ID)
		}

		batch = append(batch, ids...)
	}

	specificgroup := SpecificGroup{
		ID:       id,
		Status:   util.RULE_ENABLE,
		Page:     0,
		PageSize: 0,
	}

	specificList, _, err := specificgroup.GetList()
	if err != nil {
		return nil, err
	}

	for _, item := range specificList {
		if action > item.Action {
			action = item.Action
		}

		ruleSrv := RuleSpecific{
			RuleGroup: item.ID,
			Status:    util.RULE_ENABLE,
			Page:      0,
			PageSize:  0,
		}
		rules, _, err := ruleSrv.GetList()
		if err != nil {
			return nil, err
		}

		var ids []uint
		for _, rule := range rules {
			ids = append(ids, rule.ID)
		}

		specific = append(specific, ids...)
	}

	data := make(map[string]interface{})
	data["decoders"] = decoders
	data["action"] = action
	data["batch"] = batch
	data["specific"] = specific

	return data, err
}

func setupRulesConfig(key string, ids []uint) error {
	var data []interface{}
	for _, id := range ids {
		res, err := getRuleConfig(id)
		if err != nil {
			return nil
		}
		item := make(map[string]interface{})
		item["id"] = id
		item["config"] = res

		data = append(data, item)
	}

	if len(data) == 0 {
		data = []interface{}{}
	}

	return setCache(key, data)
}

func SetupSites() error {
	site := Site{
		Status: util.SITE_ENABLE,
	}
	sites, count, err := site.GetList()
	if err != nil {
		return err
	}
	if count == 0 {
		data := make(map[string]interface{})
		data["id"] = 0
		data["values"] = [][]struct{}{}
		data["timestamp"] = time.Now().Unix()

		siteStr, err := json.Marshal(data)
		if err != nil {
			log.Logger.Error("site", zap.String("err", err.Error()))
			return err
		}

		err = cache.Set(cache.CONFIG, cacheSiteName, string(siteStr), 0)
		if err != nil {
			log.Logger.Error("site", zap.String("err", err.Error()))
			return err
		}

		return nil
	}

	var ids []uint
	routesInfos := []map[string]interface{}{}
	for _, item := range sites {
		ids = append(ids, item.ID)

		route := make(map[string]interface{})
		route["id"] = item.ID
		route["host"] = item.Host
		route["uri"] = item.Path

		if len(item.Upstreams) != 1 {
			return fmt.Errorf("%s", "invalid site upstream")
		}

		route["upstream_id"] = item.Upstreams[0].ID
		routesInfos = append(routesInfos, route)
	}

	data := make(map[string]interface{})
	data["values"] = routesInfos
	data["timestamp"] = time.Now().Unix()

	siteStr, err := json.Marshal(data)
	if err != nil {
		log.Logger.Error("site", zap.String("err", err.Error()))
		return err
	}

	err = cache.Set(cache.CONFIG, cacheSiteName, string(siteStr), 0)
	if err != nil {
		log.Logger.Error("site", zap.String("err", err.Error()))
		return err
	}

	err = setupIPs("ip", ids)
	if err != nil {
		log.Logger.Error("site", zap.String("err", err.Error()))
		return err
	}

	err = setupCCs("cc", ids)
	if err != nil {
		log.Logger.Error("site", zap.String("err", err.Error()))
		return err
	}

	err = setupBatchRules("batch_rules")
	if err != nil {
		log.Logger.Error("site", zap.String("err", err.Error()))
		return err
	}

	err = setupSpecificRules("specific_rules")
	if err != nil {
		log.Logger.Error("site", zap.String("err", err.Error()))
		return err
	}

	err = setupRulesConfig("rules", ids)
	if err != nil {
		log.Logger.Error("site", zap.String("err", err.Error()))
		return err
	}

	return nil
}
