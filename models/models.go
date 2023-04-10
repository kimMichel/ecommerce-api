package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id          string  `json:"id" gorm:"text;not null;default:null`
	Name        string  `json:"name" gorm:"text;not null;default:null`
	Description string  `json:"description" gorm:"text;not null;default:null`
	Price       float32 `json:"price" gorm:"decimal(8,2);not null;default:null`
	Image       string  `json:"image" gorm:"text;not null;default:null`
	Rating      float32 `json:"rating" gorm:"decimal(8,2);not null;default:null`
	Quantity    int32   `json:"quantity" gorm:"int;not null;default:null`
}
