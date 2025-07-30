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

func (fw *fileWriter) Write(string) error {
	fmt.Println("write string to file...")
	return nil
}
func (dc *databaseWriter) Write(string) error {
	fmt.Println("write string to db...")
	return nil
}
func (wc *writeCloser) Close() error {
	fmt.Println("write string...")
	return nil
}

func main() {
	var mw MyWriter = &writeCloser{
		// &fileWriter{},
		&databaseWriter{},
	}
	mw.Write("dawn")
	// var mc MyCloser = &writeCloser{}

}

// type writeCloser struct{}
*/
