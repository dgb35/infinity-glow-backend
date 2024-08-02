package models

type User struct {
	Base
	ProxyIds     []string  `gorm:"type:text"`
	PasswordHash string    `json:"-" binding:"required" gorm:"not null"`
	Products     []Product `json:"products,omitempty" gorm:"many2many:user_products;"`
}
