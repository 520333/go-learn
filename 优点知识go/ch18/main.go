package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Speaker interface {
	Say(string)
}

func SpeakAlphabet(s Speaker) {
	s.Say("abcdegfhijk...xzy")
}

type Person struct {
	name string
}

func (p *Person) Say(msg string) {
	fmt.Println(p.name, ":", msg)
}

type SpeakWriter struct {
	w io.Writer
}

func (sw *SpeakWriter) Say(msg string) {
	io.WriteString(sw.w, msg)
}

type FileWriter struct {
	filename string
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	err = ioutil.WriteFile(fw.filename, p, 0644)
	n = 0
	return
}

func main() {
	james := new(Person)
	james.name = "james"
	SpeakAlphabet(james)

	console := new(SpeakWriter)
	console.w = os.Stdout
	SpeakAlphabet(console)

	fileWriter := new(FileWriter)
	fileWriter.filename = "abc.txt"
	fileSpeakWriter := new(SpeakWriter)
	fileSpeakWriter.w = fileWriter
	SpeakAlphabet(fileSpeakWriter)
}
