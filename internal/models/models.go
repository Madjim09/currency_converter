package models

import "time"

type Rates struct {
	Date       string             `json:"time_last_update_utc"`
	Base       string             `json:"base_code"`
	Rates      map[string]float64 `json:"conversion_rates"`
	PausedTime time.Time
}
