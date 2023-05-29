// Package tools
// Copyright 2016-2023 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package tools

// PaginationData 分页结果数据
type PaginationData struct {
	Total        int64   // 总记录数
	Page         int64   // 第几页
	PageSize     int64   // 每页显示多少条记录
	PageNums     int64   // 显示多少个数字
	FirstPage    int64   // 第一页
	LastPage     int64   // 最后一页
	PreviousPage int64   // 上一页,没有上一页时为 0
	NextPage     int64   // 下一页,没有下一页时为 0
	TotalPage    int64   // 总页数
	Pages        []int64 // 分页页码
}

// Pagination 分页组件
// count 记录总数
// page 当前页, 如果小于1会修正为1
// pageSize 每一页显示量,如果小于1会修正为1
// pagesNum 分页时中间显示多少个数字,如果小于1，会修正为1
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
		// 显示的页数超过最大页数，进行修正
		pg.PageNums = pg.TotalPage
	}

	if page > pg.TotalPage {
		// 当前页超过最大页，进行修正
		page = pg.TotalPage
	}

	// 上一页
	pg.PreviousPage = page - 1
	if pg.PreviousPage < 0 { // 最小页数不能是负值
		pg.PreviousPage = 0
	}

	// 下一页
	pg.NextPage = page + 1
	if pg.NextPage > pg.TotalPage { // 下一页页数不能超过总页数
		pg.NextPage = 0
	}

	if pg.PageNums > pg.TotalPage { // 如果页数比总页数还大，下半部分计算会出错
		pg.PageNums = pg.TotalPage
	}

	// 判断一下奇偶
	var halfPagesNum int64
	if pg.PageNums%2 == 0 {
		// 偶
		halfPagesNum = pg.PageNums / 2
	} else {
		// 寄
		halfPagesNum = pg.PageNums/2 + 1
	}

	// 页码前半部分
	var i int64
	for i = halfPagesNum; i > 0; i-- {
		var p = page - i
		if p <= 0 {
			continue
		}
		pg.Pages = append(pg.Pages, p)
	}

	// 页码后半部分
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
