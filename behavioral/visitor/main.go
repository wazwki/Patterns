package main

import (
	"fmt"
	"strings"
)

// Element - интерфейс для элементов, которые могут быть посещены
type Element interface {
	Accept(visitor Visitor)
}

// Visitor - интерфейс посетителя, который реализует различные операции
type Visitor interface {
	VisitFile(file *File)
	VisitFolder(folder *Folder)
}

// File - конкретный элемент (файл)
type File struct {
	name string
	size int
}

func (f *File) Accept(visitor Visitor) {
	visitor.VisitFile(f)
}

// Folder - конкретный элемент (папка)
type Folder struct {
	name     string
	contents []Element
}

func (f *Folder) Add(element Element) {
	f.contents = append(f.contents, element)
}

func (f *Folder) Accept(visitor Visitor) {
	visitor.VisitFolder(f)
}

// SizeVisitor - посетитель, который подсчитывает размер файлов и папок
type SizeVisitor struct {
	totalSize int
}

func (s *SizeVisitor) VisitFile(file *File) {
	s.totalSize += file.size
}

func (s *SizeVisitor) VisitFolder(folder *Folder) {
	for _, element := range folder.contents {
		element.Accept(s)
	}
}

func (s *SizeVisitor) GetTotalSize() int {
	return s.totalSize
}

// NameVisitor - посетитель, который отображает структуру имен файлов и папок
type NameVisitor struct {
	level int
}

func (n *NameVisitor) VisitFile(file *File) {
	fmt.Printf("%sFile: %s\n", n.getIndentation(), file.name)
}

func (n *NameVisitor) VisitFolder(folder *Folder) {
	fmt.Printf("%sFolder: %s\n", n.getIndentation(), folder.name)
	n.level++
	for _, element := range folder.contents {
		element.Accept(n)
	}
	n.level--
}

func (n *NameVisitor) getIndentation() string {
	return strings.Repeat("\t", n.level)
}

func main() {
	// Создаем файловую систему
	file1 := &File{name: "file1.txt", size: 100}
	file2 := &File{name: "file2.txt", size: 200}

	folder1 := &Folder{name: "folder1"}
	folder1.Add(file1)

	rootFolder := &Folder{name: "root"}
	rootFolder.Add(folder1)
	rootFolder.Add(file2)

	// Используем SizeVisitor для подсчета размера
	sizeVisitor := &SizeVisitor{}
	rootFolder.Accept(sizeVisitor)
	fmt.Printf("Total size: %d bytes\n", sizeVisitor.GetTotalSize())

	// Используем NameVisitor для отображения структуры
	nameVisitor := &NameVisitor{}
	rootFolder.Accept(nameVisitor)
}
