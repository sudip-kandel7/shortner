package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sudip-kandel7/shortner/package/dbConfig"
)

type User struct {
	gorm.Model

	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`

	URLs []URL `gorm:"foreignKey:UserID" json:"urls,omitempty"`
}

type URL struct {
	gorm.Model

	UserID *uint `gorm:"index" json:"user_id,omitempty"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL;" json:"-"`

	OriginalURL string `gorm:"type:text;not null" json:"original_url"`

	ShortCode string `gorm:"type:varchar(20);unique;not null" json:"short_code"`

	Clicks uint `gorm:"default:0" json:"clicks"`

	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

func init(){
	
}