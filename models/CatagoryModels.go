package models

import "time"


type CatagoryModel struct {
	ID int
	Name string
	CreatedAt time.Time
	UpdatedAt  time.Time
	Products []ProdectModel
	
}