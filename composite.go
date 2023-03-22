package main

import "fmt"

// Component interface
type Component interface {
	GetName() string
	Add(Component)
	Remove(Component)
	Display(int)
}

// File Leaf object
type File struct {
	name string
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) Add(_ Component) {
	fmt.Println("Cannot add to a file")
}

func (f *File) Remove(_ Component) {
	fmt.Println("Cannot remove from a file")
}

func (f *File) Display(indent int) {
	fmt.Println(fmt.Sprintf("%s%s", getIndent(indent), f.GetName()))
}

// Folder Composite object
type Folder struct {
	name     string
	children []Component
}

func (f *Folder) GetName() string {
	return f.name
}

func (f *Folder) Add(c Component) {
	f.children = append(f.children, c)
}

func (f *Folder) Remove(c Component) {
	for i, child := range f.children {
		if child == c {
			f.children = append(f.children[:i], f.children[i+1:]...)
			return
		}
	}
}

func (f *Folder) Display(indent int) {
	fmt.Println(fmt.Sprintf("%s%s", getIndent(indent), f.GetName()))

	for _, child := range f.children {
		child.Display(indent + 2)
	}
}

func getIndent(indent int) string {
	return fmt.Sprintf("%"+fmt.Sprintf("%d", indent)+"s", "")
}

func main() {
	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}
	file3 := &File{name: "file3"}
	folder1 := &Folder{name: "folder1"}
	folder2 := &Folder{name: "folder2"}

	folder1.Add(file1)
	folder1.Add(folder2)
	folder2.Add(file2)
	folder2.Add(file3)

	folder1.Display(0)
}
