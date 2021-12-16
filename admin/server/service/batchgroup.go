package service

import (
	"admin/server/models"
	"gorm.io/datatypes"
)

type BatchGroup struct {
	ID uint

	Name     string
	Action   int
	Priority int
	Status   int
	Level    int
	Decoder  datatypes.JSON

	Remark string

	Page     int
	PageSize int
}

func (r *BatchGroup) Save() error {
	data := map[string]interface{}{
		"name":     r.Name,
		"priority": r.Priority,
		"action":   r.Action,
		"status":   r.Status,
		"level":    r.Level,
		"decoder":  r.Decoder,

		"remark": r.Remark,
	}

	if r.ID > 0 {
		return models.UpdateBatchGroup(r.ID, data)
	}

	return models.AddBatchGroup(data)
}

func (r *BatchGroup) Delete() error {
	return models.DeleteBatchGroup(r.ID)
}

func (r *BatchGroup) Get() (*models.BatchGroup, error) {
	return models.GetBatchGroup(r.ID)
}

func (r *BatchGroup) GetList() ([]*models.BatchGroup, uint, error) {
	var query = make(map[string]interface{})
	if len(r.Name) > 0 {
		query["name"] = r.Name
	}

	if r.Status != 0 {
		query["status"] = r.Status
	}

	return models.GetBatchGroups(query, r.Page, r.PageSize)
}
