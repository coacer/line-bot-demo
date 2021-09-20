package model

import (
	"time"
)

type DbBaseModel struct {
	id        string    `json:"id"`
	createdAt time.Time `json:"createdAt"`
	updatedAt time.Time `json:"updatedAt"`
}

func (m *DbBaseModel) GetId() string {
	return m.id
}

func (m *DbBaseModel) GetCreatedAt() time.Time {
	return m.createdAt
}

func (m *DbBaseModel) GetUpdatedAt() time.Time {
	return m.updatedAt
}

func (m *DbBaseModel) SetId(id string) {
	m.id = id
}

func (m *DbBaseModel) SetCreatedAt(timestamp time.Time) {
	m.createdAt = timestamp
}

func (m *DbBaseModel) SetUpdatedAt(timestamp time.Time) {
	m.updatedAt = timestamp
}
