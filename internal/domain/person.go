package domain

import "time"

type Person struct {
	Name   string    `json:"name"`
	Mass   int       `json:"mass"`
	Create time.Time `json:"created"`
	Films  []string  `json:"films"`
}
