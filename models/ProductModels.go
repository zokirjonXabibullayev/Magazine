package models

import (
	"time"
)

type ProdectModel  struct {

	ID int
	
	ProdectType string
	Name string
	Quantity int
	Available bool
	Price int
	CreatedAt time.Time
	UptadedAt time.Time

}