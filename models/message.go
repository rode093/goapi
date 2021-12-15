package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UUID string `gorm:"primary_key"`
	Content string `gorm:"not null"`
	Attachment string
	Sender string `gorm:"not null"`
	ConversationRef uint 
}