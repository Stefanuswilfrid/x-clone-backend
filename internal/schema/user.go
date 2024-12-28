package schema

import "time"

type User struct {
	ID              uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            *string    `gorm:"type:varchar(50)" json:"name,omitempty"`
	Username        *string    `gorm:"type:varchar(30);unique" json:"username,omitempty"`
	Bio             *string    `gorm:"type:varchar(250)" json:"bio,omitempty"`
	Email           *string    `gorm:"type:varchar(100);unique" json:"email,omitempty"`
	EmailVerified   *time.Time `json:"email_verified,omitempty"`
	Image           *string    `json:"image,omitempty"`
	CoverImage      *string    `json:"cover_image,omitempty"`
	ProfileImage    *string    `json:"profile_image,omitempty"`
	HashedPassword  *string    `json:"hashed_password,omitempty"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	FollowingIds    []uint     `gorm:"-" json:"following_ids,omitempty"` // Not stored in DB; replace with a junction table if necessary
	HasNotification *bool      `json:"has_notification,omitempty"`

	Posts         []Post         `gorm:"foreignKey:UserID" json:"posts,omitempty"`
	Comments      []Comment      `gorm:"foreignKey:UserID" json:"comments,omitempty"`
	Notifications []Notification `gorm:"foreignKey:UserID" json:"notifications,omitempty"`
}
