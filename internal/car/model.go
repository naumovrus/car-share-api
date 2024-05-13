package car

import "time"

type AdditionalInfo struct {
	Mark    string    `json:"mark"`
	Uuid    string    `json:"uuid"`
	DateBuy time.Time `json:"date_buy"`
}

