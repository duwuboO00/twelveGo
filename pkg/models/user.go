
package models

import (
	"github.com/google/uuid"
	"time"
)

// User 定義用戶模型
type User struct {
	ID          uuid.UUID `json:"id" db:"id" gorm:"type:uuid;primaryKey"`
	UserName    string    `json:"user_name" db:"user_name" gorm:"type:varchar(100);not null"`
	Email       string    `json:"email" db:"email" gorm:"type:varchar(255);unique;not null"`
	LineID      *string   `json:"line_id,omitempty" db:"line_id" gorm:"type:varchar(30);unique"`
	GoogleID    *string   `json:"google_id,omitempty" db:"google_id" gorm:"type:varchar(30);unique"`
	AppleID     *string   `json:"apple_id,omitempty" db:"apple_id" gorm:"type:varchar(30);unique"`
	DateOfBirth time.Time `json:"date_of_birth" db:"date_of_birth" gorm:"type:date;not null"`
	Gender      string    `json:"gender" db:"gender" gorm:"type:char(1);check:gender IN ('M','F','U')"`
	Address     *string   `json:"address,omitempty" db:"address" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at" db:"created_at" gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" gorm:"type:timestamp with time zone"`

}
