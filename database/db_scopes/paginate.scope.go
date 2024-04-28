package db_scopes

import (
	"math"

	"gorm.io/gorm"
)

type PaginateMetadata struct {
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
	Page       int   `json:"page"`
	Take       int   `json:"take"`
}

func Paginate(page, take int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * take
		return db.Offset(offset).Limit(take)
	}
}

func PaginateMeta(rows int64, page, take int) PaginateMetadata {
	var meta PaginateMetadata

	meta.Page = page
	meta.Take = take
	meta.TotalRows = rows
	meta.TotalPages = int(math.Ceil(float64(rows) / float64(take)))

	return meta
}
