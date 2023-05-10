// pagination 分页组件
// count 记录总数
// page 当前页
// pageSize 每一页显示量
// pagesNum 分页时中间显示多少个数字
func pagination(count, page, pageSize int64, pagesNum int64) (pg *Pagination) {

	pg = &Pagination{}

	if count <= 0 || page <= 0 || pageSize <= 0 || pagesNum < 0 {
		// 没有记录
		return
	}

	//if page*pageSize > count {
	// 超过总条数，参数有问题
	//return
	//}

	var (
		previousPage, nextPage, pageCount int64
		pages                             []int64
	)

	if count%pageSize == 0 {
		pageCount = count / pageSize
	} else {
		pageCount = count/pageSize + 1
	}

	if pagesNum > pageCount {
		// 显示的页数超过最大页数，进行修正
		pagesNum = pageCount
	}

	if page > pageCount {
		// 当前页超过最大页，进行修正
		page = pageCount
	}

	//a, _ := json.Marshal(pagination(4, 10, 100, 5))
	//fmt.Println(string(a))
	//{"PreviousPage":9,"NextPage":0,"PageCount":1,"Pages":[9]}

	// 上一页
	previousPage = page - 1
	if previousPage < 0 { // 最小页数不能是负值
		previousPage = 0
	}

	// 下一页
	nextPage = page + 1
	if nextPage > pageCount { // 下一页页数不能超过总页数
		nextPage = 0
	}

	if pagesNum > pageCount { // 如果页数比总页数还大，下半部分计算会出错
		pagesNum = pageCount
	}

	// 判断一下奇偶
	var halfPagesNum int64
	if pagesNum%2 == 0 {
		// 偶
		halfPagesNum = pagesNum / 2
	} else {
		// 寄
		halfPagesNum = pagesNum/2 + 1
	}

	// 页码前半部分
	var i int64
	for i = halfPagesNum; i > 0; i-- {
		var p = page - i
		if p <= 0 {
			continue
		}
		pages = append(pages, p)
	}

	// 页码后半部分
	var n int64
	var maxN = pagesNum - int64(len(pages))
	for n = 0; n < maxN; n++ {
		var p = page + n
		if p > pageCount {
			break
		}
		pages = append(pages, p)
	}

	pg.PreviousPage = previousPage
	pg.NextPage = nextPage
	pg.PageCount = pageCount
	pg.Pages = pages

	return
}
