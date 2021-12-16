package service

import (
	"admin/server/models"
	"gorm.io/datatypes"
)

type SpecificGroup struct {
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

func (r *SpecificGroup) Save() error {
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
		return models.UpdateSpecificGroup(r.ID, data)
	}

	return models.AddSpecificGroup(data)
}

func (r *SpecificGroup) Delete() error {
	return models.DeleteSpecificGroup(r.ID)
}

func (r *SpecificGroup) Get() (*models.SpecificGroup, error) {
	return models.GetSpecificGroup(r.ID)
}

func (r *SpecificGroup) GetList() ([]*models.SpecificGroup, uint, error) {
	var query = make(map[string]interface{})
	if len(r.Name) > 0 {
		query["name"] = r.Name
	}

	if r.Status != 0 {
		query["status"] = r.Status
	}

	return models.GetSpecificGroups(query, r.Page, r.PageSize)
}
