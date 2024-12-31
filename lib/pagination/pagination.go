package pagination

import (
	"app-news/internal/core/domain/entity"
	"math"
)

type PagginationInterface interface {
	AddPagination(totalData, page, perPage int) (*entity.Page, error)
}

type Options struct{}

// AddPagination implements PagginationInterface.
func (o *Options) AddPagination(totalData int, page int, perPage int) (*entity.Page, error) {
	newPage := page

	if newPage <= 0 {
		return nil, ErrorPage
	}

	limitData := 10
	if perPage > 0 {
		limitData = perPage
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limitData)))

	last := (newPage * limitData)
	first := last - limitData

	if totalData < last {
		last = totalData
	}

	zeroPage := &entity.Page{PageCount: 1, Page: newPage}

	if totalData == 0 && newPage == 1 {
		return zeroPage, nil
	}

	if newPage > totalPage {
		return nil, ErrorHexPage
	}

	pages := &entity.Page{
		Page:       newPage,
		PerPage:    perPage,
		PageCount:  totalPage,
		TotalCount: totalData,
		First:      first,
		Last:       last,
	}
	return pages, nil
}

func NewPagination() PagginationInterface {
	pagination := new(Options)

	return pagination
}
