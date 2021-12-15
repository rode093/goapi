package models

import (
	"time"

	"gorm.io/gorm"
)
type Conversation struct{
	ID        string           `gorm:"primaryKey;uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Messages []Message `gorm:"foreignKey:ConversationRef"`
}
