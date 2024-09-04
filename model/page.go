package model

type Pagination struct {
	TotalPages   int
	VisiblePages int
	CurrentPage  int
	PageSize     int
}
