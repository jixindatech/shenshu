package service

import (
	"admin/core/log"
	"admin/server/cache"
	"admin/server/models"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"time"
)

const cacheSiteName string = "site"

type Site struct {
	ID uint

	Name        string
	Host        string
	Path        string
	UpstreamRef uint
	Remark      string

	Page     int
	PageSize int
}

func (r *Site) Save() (err error) {
	data := make(map[string]interface{})
	data["host"] = r.Host
	data["path"] = r.Path
	data["upstreamRef"] = r.UpstreamRef
	data["remark"] = r.Remark

	if r.ID > 0 {
		if r.ID == 1 {
			return fmt.Errorf("%s", "invalid id for site")
		}

		err = models.UpdateSite(r.ID, data)
	} else {
		data["name"] = r.Name
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

	return models.GetSites(data)
}

func (r *Site) Delete() error {
	if r.ID == 1 {
		return fmt.Errorf("%s", "invalid id for site")
	}

	err := models.DeleteSite(r.ID)
	if err != nil {
		return err
	}

	return SetupSites()
}

func SetupGlobalSite() error {
	var id uint = 1
	site, err := models.GetSite(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if err == gorm.ErrRecordNotFound {
		data := make(map[string]interface{})
		data["name"] = "global"
		data["host"] = "global"
		data["path"] = "/*"
		data["remark"] = "do not delete global"

		return models.AddSite(data)
	}

	if site != nil && site.Name != "global" {
		return fmt.Errorf("%s", "default global config is wrong")
	}

	return nil
}

func SetupSites() error {
	site := Site{}
	sites, count, err := site.GetList()
	if err != nil {
		return err
	}
	if count == 0 {
		data := make(map[string]interface{})
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

	routesInfos := []map[string]interface{}{}
	for _, item := range sites {
		route := make(map[string]interface{})
		route["id"] = item.ID
		route["host"] = item.Host
		route["uri"] = item.Path

		/* skip global without upstream */
		if item.ID == 1 && item.Host == "global" {
			continue
		}

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

	return nil
}
