// Package tools
// Copyright 2023 marcello<volibearw@gmail.com>. All rights reserved.
package tools

import (
	"fmt"
	"testing"
)

func TestPagination(t *testing.T) {
	//p := Pagination(100, 1, 2, 5)
	//if p.TotalPage != 50 {
	//	t.Error("total page err")
	//	return
	//}
	//if len(p.Pages) != 5 {
	//	t.Error("pages err")
	//	return
	//}
	//if p.NextPage == 0 {
	//	t.Error("next page err")
	//	return
	//}

	p := Pagination(7, 7, 1, 5)
	fmt.Println(JSONMarshalString(p))
	fmt.Println(p.Pages)

	p = Pagination(7, 5, 1, 3)
	fmt.Println(JSONMarshalString(p))
	fmt.Println(p.Pages)

	p = Pagination(7, 7, 1, 7)
	fmt.Println(JSONMarshalString(p))
	fmt.Println(p.Pages)

	p = Pagination(6, 1, 1, 3)
	fmt.Println(JSONMarshalString(p))
	fmt.Println(p.Pages)

	p = Pagination(5, 1, 1, 5)
	fmt.Println(JSONMarshalString(p))
	fmt.Println(p.Pages)
}
