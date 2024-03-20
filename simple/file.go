package simple

import "fmt"

type File struct {
	Name string
}

func NewFile(name string) (*File, func()) {
	f := &File{
		Name: name,
	}
	return f, func() {
		f.Close()
	}
}

func (f *File) Close() {
	fmt.Println("Close file: ", f.Name)
}

type Connection struct {
	File *File
}

func (c *Connection) Close() {
	fmt.Println("Close connection: ", c.File.Name)
}

func NewConnection(file *File) (*Connection, func()) {
	c := &Connection{
		File: file,
	}
	return c, func() {
		c.Close()
	}
}
