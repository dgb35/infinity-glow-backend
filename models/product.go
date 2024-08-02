package models

type Product struct {
	Base
	MacAdress string `json:"macAdress"`
	Name      string `json:"name"`
	Users     []User `gorm:"many2many:user_products;"`
}
