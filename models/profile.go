package models

import "time"

type Profile struct{
	Id string `gorm:"type:uuid;primaryKey"`
	Username string `gorm:"column:username;unique"`
	Created_At time.Time `gorm:"column:created_at"`
}