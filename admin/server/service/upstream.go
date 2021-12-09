package service

import (
	"admin/core/log"
	"admin/server/models"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"time"
)

const cacheUpstreamName string = "upstream"

type Node struct {
	IP     string `json:"ip" valid:"Required;IP;MaxSize(254)"`
	Port   int    `json:"port" valid:"Required;Min(1);Max(65535)"`
	Weight int    `json:"weight" valid:"Required;Min(1);Max(65535)"`
}

type Upstream struct {
	ID uint

	Name    string
	Lb      string
	Key     string
	Backend []byte

	Retry          int
	TimeoutConnect int
	TimeoutSend    int
	TimeoutReceive int

	Remark string

	Page     int
	PageSize int
}

func (p *Upstream) Save() (err error) {
	data := make(map[string]interface{})
	data["name"] = p.Name
	data["lb"] = p.Lb
	data["key"] = p.Key
	data["backend"] = p.Backend
	data["retry"] = p.Retry
	data["timeoutConnect"] = p.TimeoutConnect
	data["timeoutSend"] = p.TimeoutSend
	data["timeoutReceive"] = p.TimeoutReceive

	data["remark"] = p.Remark

	if p.ID > 0 {
		err = models.PutUpstream(p.ID, data)
	} else {
		err = models.AddUpstream(data)
	}

	if err != nil {
		return err
	}

	return SetupUpstreams()
}

func (p *Upstream) Get() (*models.Upstream, error) {
	return models.GetUpstream(p.ID)
}

func (p *Upstream) GetList() ([]*models.Upstream, int, error) {
	data := make(map[string]interface{})
	data["name"] = p.Name
	data["page"] = p.Page
	data["pagesize"] = p.PageSize

	return models.GetUpstreams(data)
}

func (p *Upstream) Delete() error {
	err := models.DeleteUpstream(p.ID)
	if err != nil {
		return err
	}

	return SetupUpstreams()
}

func SetupUpstreams() error {
	upstream := Upstream{}
	upstreams, count, err := upstream.GetList()
	if err != nil {
		return err
	}
	if count == 0 {
		data := make(map[string]interface{})
		data["values"] = [][]struct{}{}
		data["timestamp"] = time.Now().Unix()

		upstreamStr, err := json.Marshal(data)
		if err != nil {
			log.Logger.Error("upstream", zap.String("err", err.Error()))
			return err
		}
		fmt.Println(upstreamStr)
		/*
			err = cache.Set(cacheUpstreamName, string(upstreamStr))
			if err != nil {
				golog.Error("upstream", zap.String("err", err.Error()))
				return err
			}
		*/
		return nil
	}

	upstreamInfos := []map[string]interface{}{}
	for _, upstream := range upstreams {
		item := make(map[string]interface{})
		item["id"] = upstream.ID
		item["type"] = upstream.Lb
		item["name"] = upstream.Name
		item["retry"] = upstream.Retry

		timeout := make(map[string]interface{})
		timeout["connect"] = upstream.TimeoutConnect
		timeout["send"] = upstream.TimeoutSend
		timeout["receive"] = upstream.TimeoutReceive

		item["timeout"] = timeout

		nodes := []Node{}
		err := json.Unmarshal(upstream.Backend, &nodes)
		if err != nil {
			return err
		}

		nodeInfos := make(map[string]int)
		for _, node := range nodes {
			port := strconv.FormatInt(int64(node.Port), 10)
			nodeInfo := node.IP + ":" + port
			nodeInfos[nodeInfo] = node.Weight
		}

		item["nodes"] = nodeInfos

		upstreamInfos = append(upstreamInfos, item)
	}

	data := make(map[string]interface{})
	data["values"] = upstreamInfos
	data["timestamp"] = time.Now().Unix()

	upstreamStr, err := json.Marshal(data)
	if err != nil {
		log.Logger.Error("upstream", zap.String("err", err.Error()))
		return err
	}
	fmt.Println(upstreamStr)
	/*
		err = cache.Set(cacheUpstreamName, string(upstreamStr))
		if err != nil {
			golog.Error("upstream", zap.String("err", err.Error()))
			return err
		}
	*/
	return nil
}
