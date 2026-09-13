package book

import (
	"errors"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func NewBook(title, author string, pages int) (*Book, error) {
	if title == "" {
		return nil, errors.New("el titulo no peude ser vacioo")
	}

	if author == "" {
		return nil, errors.New("el author no puede ser cadena vacia")
	}

	if pages <= 0 {
		return nil, errors.New("el libro tiene que tener paginas amigue :v")
	}

	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}, nil
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s, Author: %s, Pages: %d \n", b.Title, b.Author, b.Pages)
}
