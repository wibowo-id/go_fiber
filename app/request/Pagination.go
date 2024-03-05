package request

import "dlh_oss_be/utils/paginator"

type Pagination struct {
	Pagination *paginator.Pagination `json:"pagination"`
}
