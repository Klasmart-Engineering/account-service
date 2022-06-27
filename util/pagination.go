package util

const DefaultPageSize = 10

type OffsetPaginationArgs struct {
	Limit  int `form:"limit" binding:"omitempty,gte=1,lte=25" default:"10" minimum:"1" maximum:"25"`
	Offset int `form:"offset" binding:"omitempty,gte=0" default:"0" minimum:"0"`
}
