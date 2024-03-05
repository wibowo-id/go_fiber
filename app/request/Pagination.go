package request

import "go_fiber_wibowo/utils/paginator"

type Pagination struct {
	Pagination *paginator.Pagination `json:"pagination"`
}
