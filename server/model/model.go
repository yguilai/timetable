package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	gorm.Model `gorm:"-" json:"-"`
	ID         uint           `json:"id" gorm:"primarykey"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt" gorm:"index" `
}
