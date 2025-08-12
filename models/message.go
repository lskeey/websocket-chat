package models

import "time"

type Message struct {
	ID          uint   `gorm:"primaryKey"`
	Content     string `gorm:"not null"`
	SenderID    uint   `gorm:"not null"`
	RecipientID uint   `gorm:"not null"`

	Sender    User `gorm:"foreignKey:SenderID"`
	Recipient User `gorm:"foreignKey:RecipientID"`
	CreatedAt time.Time
}
