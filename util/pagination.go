package util

const DefaultPageSize = 10

type OffsetPaginationArgs struct {
	Limit  int `form:"limit" binding:"omitempty,gte=1,lte=25"`
	Offset int `form:"offset" binding:"omitempty,gte=0"`
}
