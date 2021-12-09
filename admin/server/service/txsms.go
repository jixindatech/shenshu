package service

import (
	"admin/server/models"
)

type Txsms struct {
	ID uint

	Host      string
	SecretId  string
	SecretKey string
}

var txsmsCache *Txsms

func (t *Txsms) Save() error {
	data := make(map[string]interface{})
	data["host"] = t.Host
	data["secretId"] = t.SecretId
	data["secretKey"] = t.SecretKey

	if t.ID > 0 {
		txsmsCache = nil
		return models.UpdateTxsms(t.ID, data)
	}

	return models.AddTxsms(data)
}

func (t *Txsms) Get() (*models.Txsms, error) {
	return models.GetTxsms()
}

func getTxsmsConfig() (*Txsms, error) {
	txsmsCache = new(Txsms)
	if txsmsCache == nil {
		txsms, err := txsmsCache.Get()
		if err != nil {
			return nil, err
		}
		txsmsCache.Host = txsms.Host
		txsmsCache.SecretKey = txsms.SecretKey
		txsmsCache.SecretId = txsms.SecretID
		return txsmsCache, nil
	}

	return txsmsCache, nil
}
