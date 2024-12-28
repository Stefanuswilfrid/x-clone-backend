package schema

import "time"

type Notification struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Body      string    `gorm:"type:text" json:"body"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	User User `gorm:"constraint:OnDelete:CASCADE;" json:"user,omitempty"`
}
