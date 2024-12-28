package schema

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Body      string    `gorm:"type:text" json:"body"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	LikedIds  []uint    `gorm:"-" json:"liked_ids,omitempty"` // Not stored in DB; replace with a junction table if necessary
	Image     *string   `json:"image,omitempty"`

	User     User      `gorm:"constraint:OnDelete:CASCADE;" json:"user,omitempty"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}
