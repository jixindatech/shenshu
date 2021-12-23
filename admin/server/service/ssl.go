package service

import (
	"admin/core/log"
	"admin/server/cache"
	"admin/server/models"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"time"
)

const cacheSSLName string = "ssl"

type SSL struct {
	ID uint

	Name   string
	Pub    string
	Pri    string
	Remark string

	Page     int
	PageSize int
}

func (c *SSL) Save() (err error) {
	var dnsNames []string
	cert, err := tls.X509KeyPair([]byte(c.Pub), []byte(c.Pri))
	if err != nil {
		return err
	}

	leaf := cert.Leaf
	if leaf != nil {
		for _, san := range leaf.DNSNames {
			if len(san) == 0 {
				return fmt.Errorf("%s", "invalid cert dnsname")
			}
			dnsNames = append(dnsNames, san)
		}
	} else {
		certificat, err := x509.ParseCertificate(cert.Certificate[0])
		if err != nil {
			return err
		}
		for _, san := range certificat.DNSNames {
			if len(san) == 0 {
				return fmt.Errorf("%s", "invalid cert dnsname")
			}
			dnsNames = append(dnsNames, san)
		}
	}
	data := make(map[string]interface{})
	data["name"] = c.Name
	data["pub"] = c.Pub
	data["pri"] = c.Pri
	data["remark"] = c.Remark

	names, err := json.Marshal(dnsNames)
	if err != nil {
		return err
	}
	data["server"] = names

	if c.ID > 0 {
		err = models.PutSSL(c.ID, data)
	} else {
		err = models.AddSSL(data)
	}
	if err != nil {
		return err
	}

	return nil
}

func (c *SSL) Get() (*models.SSL, error) {
	return models.GetSSL(c.ID)
}

func (c *SSL) GetList() ([]*models.SSL, int, error) {
	data := make(map[string]interface{})
	data["name"] = c.Name
	data["page"] = c.Page
	data["pagesize"] = c.PageSize

	return models.GetSSLs(data)
}

func (c *SSL) Delete() error {
	err := models.DeleteSSL(c.ID)
	if err != nil {
		return err
	}
	return nil
}

func SetupSSLs() error {
	cert := SSL{
		Page: 0,
	}
	certs, count, err := cert.GetList()
	if err != nil {
		return err
	}
	if count == 0 {
		serverData := make(map[string]interface{})
		serverData["timestamp"] = time.Now().Unix()
		serverData["values"] = [][]struct{}{}

		certData, err := json.Marshal(serverData)
		if err != nil {
			log.Logger.Error("cert", zap.String("err", err.Error()))
			return err
		}

		err = cache.Set(cache.CONFIG, cacheSSLName, string(certData), 0)
		if err != nil {
			log.Logger.Error("cert", zap.String("err", err.Error()))
			return err
		}

		return nil
	}

	servers := []map[string]interface{}{}
	for _, item := range certs {
		pub := item.Pub
		key := item.Pri
		var dnsNames []string
		err := json.Unmarshal(item.Server, &dnsNames)
		if err != nil {
			return err
		}
		for _, name := range dnsNames {
			if len(name) > 0 {
				ssl := make(map[string]interface{})
				ssl["name"] = name
				ssl["timestamp"] = item.UpdatedAt.Unix()
				ssl["config"] = map[string]interface{}{
					"sni":  name,
					"cert": pub,
					"key":  key,
				}
				servers = append(servers, ssl)
			} else {
				return fmt.Errorf("%s", "invalid dnsname")
			}
		}
	}

	serverData := make(map[string]interface{})
	serverData["timestamp"] = time.Now().Unix()
	serverData["values"] = servers

	certData, err := json.Marshal(serverData)
	if err != nil {
		log.Logger.Error("ssl", zap.String("err", err.Error()))
		return err
	}

	err = cache.Set(cache.CONFIG, cacheSSLName, string(certData), 0)
	if err != nil {
		log.Logger.Error("ssl", zap.String("err", err.Error()))
		return err
	}

	return nil
}
