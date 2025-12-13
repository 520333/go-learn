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

myfunc (s *SreaWriter) Write(string) {
	fmt.Println("write")

}
myfunc (s *SreaWriter) Read() string {
	fmt.Println("read")
	return ""
}
myfunc (s *SreaWriter) ReadWrite() {
	fmt.Println("read and write")
}
myfunc main() {
	var mrw MyReaderWriter = &SreaWriter{}
	mrw.Read()
}
*/
