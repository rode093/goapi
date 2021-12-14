package models

import "gorm.io/gorm"
type Conversation struct{
	ID        string           `gorm:"primaryKey;uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	}

type Message struct {
	gorm.Model()

}