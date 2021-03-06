package models

import (
	"admin/core/log"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
	"time"
)

type Site struct {
	Model

	Name   string `json:"name" gorm:"column:name;not null;unique"`
	Host   string `json:"host" gorm:"column:host;not null;"`
	Path   string `json:"path" gorm:"column:path;not null"`
	Status int    `json:"status" gorm:"column:status;not null"`
	Remark string `json:"remark" gorm:"column:remark"`

	UpstreamTimestamp int64 `json:"upstream_timestamp" gorm:"column:upstream_timestamp;default 0"`
	IPTimestamp       int64 `json:"ip_timestamp" gorm:"column:ip_timestamp;default 0"`
	CCTimestamp       int64 `json:"cc_timestamp" gorm:"column:cc_timestamp;default 0"`
	RuleTimestamp     int64 `json:"rule_timestamp" gorm:"column:rule_timestamp;default 0"`

	Upstreams      []*Upstream      `json:"upstreamRef" gorm:"many2many:site_upstream;"`
	IPs            []IP             `json:"ips"`
	CCs            []*CC            `json:"ccs"`
	BatchGroups    []*BatchGroup    `json:"batchgroup" gorm:"many2many:site_batchgroup;"`
	SpecificGroups []*SpecificGroup `json:"specificgroup" gorm:"many2many:site_specificgroup;"`
}

func changeSiteTimestamp(id uint, column string) {
	site := Site{}
	site.Model.ID = id

	data := make(map[string]interface{})
	data[column] = time.Now().Unix()
	err := db.Model(&site).Update(data).Error
	if err != nil {
		log.Logger.Error("site", zap.String("err", err.Error()))
	}
}

func AddSite(data map[string]interface{}) error {
	site := &Site{
		Name:   data["name"].(string),
		Host:   data["host"].(string),
		Path:   data["path"].(string),
		Status: data["status"].(int),
		Remark: data["remark"].(string),
	}

	err := db.Create(&site).Error
	if err != nil {
		return err
	}

	/* for global site */
	_, ok := data["upstreamRef"]
	if !ok {
		return nil
	}

	var upstreams []*Upstream
	temp := Upstream{}
	temp.Model.ID = data["upstreamRef"].(uint)
	upstreams = append(upstreams, &temp)

	err = db.Debug().Model(&site).Association("Upstreams").Replace(upstreams).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteSite(id uint) error {
	site := Site{}
	site.Model.ID = id
	/* clear Associations with upstreams
	db.Model(&site).Association("Upstreams").Clear()
	db.Model(&site).Association("IPs").Clear()
	*/
	err := db.Select(clause.Associations).Delete(&site).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateSiteBatchRuleGroup(id uint, ids []uint) error {
	site, err := GetSite(id)
	if err != nil {
		return err
	}
	var rulegroups []*BatchGroup
	for _, item := range ids {
		temp := BatchGroup{}
		temp.Model.ID = item
		rulegroups = append(rulegroups, &temp)
	}

	err = db.Model(&site).Association("BatchGroups").Replace(rulegroups).Error
	if err != nil {
		return err
	}

	changeSiteTimestamp(site.ID, "RuleTimestamp")
	return nil
}

func GetSiteBatchGroup(id uint) ([]*BatchGroup, error) {
	site := Site{}
	site.Model.ID = id
	var rulegroups []*BatchGroup
	err := db.Model(&site).Association("BatchGroups").Find(&rulegroups).Error
	if err != nil {
		return nil, err
	}
	return rulegroups, nil
}

func UpdateSiteSpecificRuleGroup(id uint, ids []uint) error {
	site, err := GetSite(id)
	if err != nil {
		return err
	}

	var rulegroups []*SpecificGroup
	for _, item := range ids {
		temp := SpecificGroup{}
		temp.Model.ID = item
		rulegroups = append(rulegroups, &temp)
	}

	err = db.Model(&site).Association("SpecificGroups").Replace(rulegroups).Error
	if err != nil {
		return err
	}

	changeSiteTimestamp(site.ID, "RuleTimestamp")

	return nil
}

func GetSiteSpecificGroup(id uint) ([]*SpecificGroup, error) {
	site := Site{}
	site.Model.ID = id
	var rulegroups []*SpecificGroup
	err := db.Model(&site).Association("SpecificGroups").Find(&rulegroups).Error
	if err != nil {
		return nil, err
	}
	return rulegroups, nil
}

func UpdateSite(id uint, data map[string]interface{}) error {
	site := Site{}
	site.Model.ID = id

	var upstreams []*Upstream
	temp := Upstream{}
	temp.Model.ID = data["upstreamRef"].(uint)
	upstreams = append(upstreams, &temp)

	delete(data, "upstreamRef")

	err := db.Model(&site).Update(data).Error
	if err != nil {
		return err
	}

	err = db.Debug().Model(&site).Association("Upstreams").Replace(upstreams).Error
	if err != nil {
		return err
	}

	return nil
}

func GetSite(id uint) (*Site, error) {
	var site Site

	err := db.Preload("Upstreams").Where("id = ?", id).Find(&site).Error
	if err != nil {
		return &site, err
	}

	return &site, nil
}

/*
func GetSiteConfig(id uint) (map[string]interface{}, error) {
	var site Site
	site.Model.ID = id
	data := make(map[string]interface{})

	var err error
	var ips []*IP
	err = db.Model(&site).Association("IPs").Find(&ips).Error
	if err != nil {
		return nil, err
	}

	var ipsAllow []string
	var ipsDeny []string
	for _, ip := range ips {
		fmt.Println(ip.ID, ip.Type, ip.IP)
		var tmpIPs []string
		err := json.Unmarshal(ip.IP, &tmpIPs)
		if err != nil {
			return nil, err
		}

		if ip.Type == util.IP_ACCEPT {
			ipsAllow = append(ipsAllow, tmpIPs...)
		} else if ip.Type == util.IP_DENY {
			ipsDeny = append(ipsDeny, tmpIPs...)
		} else {
			return nil, fmt.Errorf("%s", "invalid ip type")
		}
	}
	/*
		err := db.Preload("BatchGroups").Preload("Upstreams").Where("id = ?", id).Find(&site).Error
		if err != nil {
			return &site, err
		}

	return data, nil
}
*/

func GetSites(data map[string]interface{}) ([]*Site, int, error) {
	var sites []*Site
	name := data["name"].(string)
	page := data["page"].(int)
	pageSize := data["pagesize"].(int)

	search := make(map[string]interface{})
	if data["status"].(int) != 0 {
		search["status"] = data["status"].(int)
	}

	var count int
	if page > 0 {
		offset := (page - 1) * pageSize
		if len(name) > 0 {
			name = "%" + name + "%"
			err := db.Preload("Upstreams").Where(search).Where("name LIKE ?", name).Offset(offset).Limit(pageSize).Find(&sites).Count(&count).Error
			if err != nil {
				return nil, 0, err
			}
		} else {
			err := db.Preload("Upstreams").Where(search).Offset(offset).Limit(pageSize).Find(&sites).Count(&count).Error
			if err != nil {
				return nil, 0, err
			}
		}
	} else {
		err := db.Preload("Upstreams").Where(search).Find(&sites).Count(&count).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0, err
		}
	}

	return sites, count, nil
}
