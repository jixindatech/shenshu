package service

import (
	"admin/server/models"
)

type Msg struct {
	ID uint

	Name    string
	Content string

	Remark string

	Page     int
	PageSize int
}

func (m *Msg) Save() error {
	data := map[string]interface{}{
		"name":    m.Name,
		"content": m.Content,
		"remark":  m.Remark,
	}

	if m.ID > 0 {
		return models.UpdateMsg(m.ID, data)
	}

	return models.AddMsg(data)
}

func (m *Msg) Delete() error {
	return models.DeleteMsg(m.ID)
}

func (m *Msg) Get() (*models.Msg, error) {
	return models.GetMsg(m.ID)
}

func (m *Msg) GetList() ([]*models.Msg, uint, error) {
	var query = make(map[string]interface{})
	if len(m.Name) > 0 {
		query["name"] = m.Name
	}

	return models.GetMsgs(query, m.Page, m.PageSize)
}

func (m *Msg) SendMsgs(ids []uint) error {
	_, err := m.Get()
	if err != nil {
		return err
	}

	users, err := models.GetBatchUser(ids)
	if err != nil {
		return err
	}

	var phones []string
	for _, user := range users {
		phones = append(phones, user.Phone)
	}

	return nil
}
