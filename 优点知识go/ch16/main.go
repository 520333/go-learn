package main

import "fmt"

type Book struct {
	id      int
	title   string
	author  string
	subject string
}

func NewBook(id int, title, author, subject string) *Book {
	return &Book{id, title, author, subject}
}
func (b *Book) String() string {
	return fmt.Sprintf("id=%d, title=%s, author=%s, subject=%s", b.id, b.title, b.author, b.subject)
}
func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetTitle() string {
	return b.title
}
func (b *Book) SetTitle(title string) string {
	b.title = title
	return title
}
func main() {
	var book1 *Book
	book1 = new(Book)
	book1.id = 1
	book1.title = "go入门"
	book1.author = "宝哥"
	book1.subject = "学好指针"
	fmt.Println(book1)

	book2 := Book{
		id:      2,
		title:   "python",
		author:  "dawn",
		subject: "学好算法",
	}
	// fmt.Println(&book2)
	fmt.Println(book2.String())

	book3 := NewBook(3, "Java", "宝哥", "数据类型")
	fmt.Println(book3.String())

}
