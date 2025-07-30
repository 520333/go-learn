package example

import "fmt"

type Course struct {
	Title   string
	SubJect string
}

func (c *Course) String() string {
	return fmt.Sprintf("[course]{Title=%s,Subject=%s}", c.Title, c.SubJect)
}
