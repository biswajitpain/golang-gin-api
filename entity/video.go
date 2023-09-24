package entity

import "time"

type Person struct {
	ID        uint64 `gorm:"primary_key;auto_incriment" json:"id"`
	FristName string `json:"firstname" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(32)"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" validate:"required,email" gorm:"type:varchar(256)"`
}

type Video struct {
	ID          uint64    `gorm:"primary_key;auto_incriment" json:"id"`
	Title       string    `json:"title" binding:"min=2,max=10" gorm:"type:varcahr(100)"`
	Description string    `json:"description" binding:"max=20" gorm:"type:varcahr(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varcahr(100);UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignkey:PersonID"`
	PersonID    uint64    `json:"-"`
	CreateedAt  time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
