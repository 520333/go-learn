package main

/*
import "fmt"

type MyWriter interface {
	Write(string) error
}
type MyCloser interface {
	Close() error
}

type writeCloser struct {
	MyWriter // interface也是一个类型
}
type fileWriter struct {
	filePath string
}
type databaseWriter struct {
	host string
	port string
	db   string
}

myfunc (fw *fileWriter) Write(string) error {
	fmt.Println("write string to file...")
	return nil
}
myfunc (dc *databaseWriter) Write(string) error {
	fmt.Println("write string to db...")
	return nil
}
myfunc (wc *writeCloser) Close() error {
	fmt.Println("write string...")
	return nil
}

myfunc main() {
	var mw MyWriter = &writeCloser{
		// &fileWriter{},
		&databaseWriter{},
	}
	mw.Write("dawn")
	// var mc MyCloser = &writeCloser{}

}

// type writeCloser struct{}
*/
