package domain

import "time"

type Planet struct {
	Name     string    `json:"name"`
	Resident []string  `json:"residents"`
	Created  time.Time `json:"created"`
	Films    []string  `json:"films"`
}
