package domain

type PaginatedResponse[T any] struct {
	Count    int  `json:"count"`
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	NextPage *int `json:"next_page,omitempty"`
	PrevPage *int `json:"prev_page,omitempty"`
	Results  []T  `json:"results"`
}
