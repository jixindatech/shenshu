package service

import (
	"admin/core/log"
	"admin/server/cache"
	"admin/server/models"
	"admin/server/util"
	"encoding/json"
	"go.uber.org/zap"
	"time"
)

const cacheGlobalIPName = "shenshu_globalip"

type GlobalIP struct {
	ID uint

	Name   string
	Type   int
	IP     []string
	Remark string

	Sites []uint

	Page     int
	PageSize int
}

func (i *GlobalIP) Save() (err error) {
	data := make(map[string]interface{})
	data["name"] = i.Name
	data["type"] = i.Type

	ip, err := json.Marshal(&i.IP)
	if err != nil {
		return err
	}
	data["ip"] = ip
	data["remark"] = i.Remark

	if i.ID > 0 {
		return models.UpdateGlobalIP(i.ID, data)
	}

	return models.AddGlobalIP(data)
}

func (i *GlobalIP) GetList() ([]*models.GlobalIP, int, error) {
	data := make(map[string]interface{})
	data["name"] = i.Name
	data["type"] = i.Type
	data["page"] = i.Page
	data["pagesize"] = i.PageSize

	return models.GetGlobalIPs(data)
}

func (i *GlobalIP) Get() (*models.GlobalIP, error) {
	return models.GetGlobalIP(i.ID)
}

func (i *GlobalIP) Delete() error {
	err := models.DeleteGlobalIP(i.ID)
	if err != nil {
		return err
	}

	return nil
}

func SetupGlobalIPs() error {
	ipSrv := GlobalIP{
		Type:     util.IP_ACCEPT,
		Page:     0,
		PageSize: 0,
	}
	list, _, err := ipSrv.GetList()
	if err != nil {
		return err
	}

	var ipsAccept []string
	for _, ip := range list {
		var tmpIPs []string
		err := json.Unmarshal(ip.IP, &tmpIPs)
		if err != nil {
			return err
		}
		ipsAccept = append(ipsAccept, tmpIPs...)
	}
	if len(ipsAccept) == 0 {
		ipsAccept = []string{}
	}

	ipSrv.Type = util.IP_DENY
	list, _, err = ipSrv.GetList()
	if err != nil {
		return err
	}

	var ipsDeny []string
	for _, ip := range list {
		var tmpIPs []string
		err := json.Unmarshal(ip.IP, &tmpIPs)
		if err != nil {
			return err
		}
		ipsDeny = append(ipsDeny, tmpIPs...)
	}
	if len(ipsDeny) == 0 {
		ipsDeny = []string{}
	}

	data := make(map[string]interface{})
	data["values"] = []map[string]interface{}{
		0: map[string]interface{}{
			"config": map[string]interface{}{
				"accept": ipsAccept,
				"deny":   ipsDeny,
			},
			"name":      "globalip",
			"timestamp": time.Now().Unix(),
		},
	}
	data["timestamp"] = time.Now().Unix()

	jsonStr, err := json.Marshal(data)
	if err != nil {
		log.Logger.Error("globalip", zap.String("err", err.Error()))
		return err
	}

	err = cache.Set(cache.CONFIG, cacheGlobalIPName, string(jsonStr), 0)
	if err != nil {
		log.Logger.Error("globalip", zap.String("err", err.Error()))
		return err
	}

	return nil
}
