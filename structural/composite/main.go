package main

import "fmt"

// FileComponent - общий интерфейс для файлов и каталогов
type FileComponent interface {
	getName() string
	getSize() int
	add(FileComponent)
	remove(FileComponent)
	display(indent string)
}

// File - конкретный компонент, представляющий файл
type File struct {
	name string
	size int
}

func (f *File) getName() string {
	return f.name
}

func (f *File) getSize() int {
	return f.size
}

func (f *File) add(FileComponent) {
	fmt.Println("Cannot add to a file.")
}

func (f *File) remove(FileComponent) {
	fmt.Println("Cannot remove from a file.")
}

func (f *File) display(indent string) {
	fmt.Println(indent + "File: " + f.getName() + " (" + fmt.Sprint(f.getSize()) + " KB)")
}

// Directory - конкретный компонент, представляющий каталог
type Directory struct {
	name       string
	components []FileComponent
}

func (d *Directory) getName() string {
	return d.name
}

func (d *Directory) getSize() int {
	totalSize := 0
	for _, component := range d.components {
		totalSize += component.getSize()
	}
	return totalSize
}

func (d *Directory) add(component FileComponent) {
	d.components = append(d.components, component)
}

func (d *Directory) remove(component FileComponent) {
	for i, comp := range d.components {
		if comp.getName() == component.getName() {
			d.components = append(d.components[:i], d.components[i+1:]...)
			break
		}
	}
}

func (d *Directory) display(indent string) {
	fmt.Println(indent + "Directory: " + d.getName())
	for _, component := range d.components {
		component.display(indent + "  ")
	}
}

func main() {
	// Создаем файлы
	file1 := &File{name: "file1.txt", size: 120}
	file2 := &File{name: "file2.txt", size: 200}
	file3 := &File{name: "file3.txt", size: 320}

	// Создаем директории
	dir1 := &Directory{name: "Documents"}
	dir2 := &Directory{name: "Projects"}

	// Добавляем файлы в директории
	dir1.add(file1)
	dir1.add(file2)
	dir2.add(file3)

	// Создаем корневую директорию
	rootDir := &Directory{name: "Root"}
	rootDir.add(dir1)
	rootDir.add(dir2)

	// Отображаем структуру файловой системы
	rootDir.display("")

	// Подсчитываем общий размер корневого каталога
	fmt.Printf("\nTotal size of '%s': %d KB\n", rootDir.getName(), rootDir.getSize())
}
