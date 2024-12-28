package schema

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Body      string    `gorm:"type:text" json:"body"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	PostID    uint      `gorm:"not null" json:"post_id"`

	User User `gorm:"constraint:OnDelete:CASCADE;" json:"user,omitempty"`
	Post Post `gorm:"constraint:OnDelete:CASCADE;" json:"post,omitempty"`
}
