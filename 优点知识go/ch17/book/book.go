package book

import (
	"fmt"
	"reflect"
)

type book struct {
	Id      int    "this is a book id"
	Title   string "this is a book title"
	Author  string "this is a book author"
	Subject string "this is a book subject"
}

type techBook struct {
	Cat string
	int
	book
}

func NewBook(Id int, Title, Author, Subject string) *book {
	return &book{Id, Title, Author, Subject}
}

func (b *book) String() string {
	return fmt.Sprintf("id=%d, title=%s, author=%s, subject=%s", b.Id, b.Title, b.Author, b.Subject)
}

func RefTag(b book, i int) {
	bType := reflect.TypeOf(b)
	iField := bType.Field(i)
	fmt.Printf("%v\n", iField.Tag)
}
func InitTechBook() {
	bk := NewBook(1000, "golang", "dawn", "pointer")
	tb := new(techBook)
	tb.Cat = "tech"
	tb.int = 10
	tb.book = *bk

	fmt.Printf("techBook Cat=%s\n", tb.Cat)
	fmt.Printf("techBook int=%d\n", tb.int)
	fmt.Printf("techBook book=%s\n", tb.book.String())
}
