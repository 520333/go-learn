package main

// interface结构嵌套
/*
import "fmt"

type MyWriter interface{ Write(string) }
type MyReader interface{ Read() string }
type MyReaderWriter interface {
	MyWriter
	MyReader
	ReadWrite()
}

type SreaWriter struct{}

func (s *SreaWriter) Write(string) {
	fmt.Println("write")

}
func (s *SreaWriter) Read() string {
	fmt.Println("read")
	return ""
}
func (s *SreaWriter) ReadWrite() {
	fmt.Println("read and write")
}
func main() {
	var mrw MyReaderWriter = &SreaWriter{}
	mrw.Read()
}
*/
