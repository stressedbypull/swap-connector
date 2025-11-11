package domain

type PaginatedResponse[T any] struct {
	Count    int `json:"count"`
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Results  []T `json:"results"`
}
