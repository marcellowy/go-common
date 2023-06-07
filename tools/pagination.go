// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

// PaginationData pagination data
type PaginationData struct {
	Total        int64   // total record
	Page         int64   // current page
	PageSize     int64   // show size every page
	PageNums     int64   // show digital numbers on page
	FirstPage    int64   // first page
	LastPage     int64   // last page
	PreviousPage int64   // previous page
	NextPage     int64   // next page
	TotalPage    int64   // total page
	Pages        []int64 // pages
}

// Pagination calculate
// total record
// page  current page
// pageSize show size every page
// pagesNum show digital numbers on page
func Pagination(total, page, pageSize int64, pageNum int64) (pg *PaginationData) {

	pg = &PaginationData{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		PageNums: pageNum,
	}

	if pg.Page < 1 {
		pg.Page = 1
	}

	if pg.PageSize < 1 {
		pg.PageSize = 1
	}

	if pg.PageNums < 1 {
		pg.PageNums = 1
	}

	if pg.Total <= 0 {
		// 没有记录
		return
	}

	if total%pageSize == 0 {
		pg.TotalPage = total / pageSize
	} else {
		pg.TotalPage = total/pageSize + 1
	}

	if pg.PageNums > pg.TotalPage {
		// max not than total page
		pg.PageNums = pg.TotalPage
	}

	if page > pg.TotalPage {
		// max not than total page
		page = pg.TotalPage
	}

	// previous page
	pg.PreviousPage = page - 1
	if pg.PreviousPage < 0 {
		pg.PreviousPage = 0
	}

	// next page
	pg.NextPage = page + 1
	if pg.NextPage > pg.TotalPage { // next page max value is total page
		pg.NextPage = 0
	}

	if pg.PageNums > pg.TotalPage { // if page > total page, then page = total page
		pg.PageNums = pg.TotalPage
	}

	// check even or odd
	var halfPagesNum int64
	if pg.PageNums%2 == 0 {
		// even
		halfPagesNum = pg.PageNums / 2
	} else {
		// odd
		halfPagesNum = pg.PageNums/2 + 1
	}

	// the part 1
	var i int64
	for i = halfPagesNum; i > 0; i-- {
		var p = page - i
		if p <= 0 {
			continue
		}
		pg.Pages = append(pg.Pages, p)
	}

	// the part 2
	var n int64
	var maxN = pg.PageNums - int64(len(pg.Pages))
	for n = 0; n < maxN; n++ {
		var p = page + n
		if p > pg.TotalPage {
			break
		}
		pg.Pages = append(pg.Pages, p)
	}

	return
}
