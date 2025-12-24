package common

type Pagination struct {
	TotalData   int  `json:"total_data"`
	Limit       int  `json:"limit"`
	CurrentPage int  `json:"current_page"`
	TotalPages  int  `json:"total_pages"`
	HasPrev     bool `json:"has_prev"`
	HasNext     bool `json:"has_next"`
}

type PaginationQuery struct {
	Limit  int
	Offset int
}
