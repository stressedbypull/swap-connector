package domain

type PaginatedResponse[T any] struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Results  []T `json:"results"`
}
