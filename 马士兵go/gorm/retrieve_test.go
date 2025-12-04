package gorm

import "testing"

func TestGetByPK(t *testing.T) {
	GetByPK()
}

func TestGetOne(t *testing.T) {
	GeOne()
}

func TestGetMap(t *testing.T) {
	GetToMap()
}

func TestGetPluck(t *testing.T) {
	GetPluck()
}

func TestGetPluckExp(t *testing.T) {
	GetPluckExp()
}

func TestGetSelect(t *testing.T) {
	GetSelect()
}

func TestGetDistinct(t *testing.T) {
	GetDistinct()
}

func TestWhereMethod(t *testing.T) {
	WhereMethod()
}

func TestWhereType(t *testing.T) {
	WhereType()
}

func TestPlaceHolder(t *testing.T) {
	PlaceHolder()
}

func TestOrderBy(t *testing.T) {
	OrderBy()
}

func TestPagination(t *testing.T) {
	req := Pager{3, 15}
	Pagination(req)
}

func TestPaginationScope(t *testing.T) {
	req := Pager{4, 5}
	PaginationScope(req)
}

func TestGroupHaving(t *testing.T) {
	GroupHaving()
}

func TestCount(t *testing.T) {
	req := Pager{3, 15}
	Count(req)
}

func TestIterator(t *testing.T) {
	Iterator()
}

func TestLocking(t *testing.T) {
	Locking()
}
