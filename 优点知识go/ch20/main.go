package main

import (
	"fmt"
	"learn/ch20/example"
)

type Appender interface {
	Append(int)
}
type Lener interface {
	Len() int
}
type List []int

func (l List) Len() int {
	return len(l)
}
func (l *List) Append(val int) {
	*l = append(*l, val)
}

func Counter(a Appender, start, end int) {
	for i := start; i < end; i++ {
		a.Append(i)
	}
}
func IsLarge(l Lener) bool {
	return l.Len() > 50
}
func main() {
	var list List
	Counter(&list, 1, 10)
	fmt.Println(list)
	plst := new(List)
	Counter(plst, 1, 52)
	fmt.Println(*plst)
	fmt.Println(IsLarge(plst))

	course := new(example.Course)
	course.Title = "golang"
	course.SubJect = "golang实战"
	fmt.Println(course)

	data := []int{23, 50, 78, 11, 19, 60, 100, 1000}
	ia := example.IntArray(data)
	example.Sort(ia)
	fmt.Println(ia)

}
