package book

import (
	"errors"
	"fmt"
)

type Book struct {
	title  string
	author string
	pages  int
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
		title:  title,
		author: author,
		pages:  pages,
	}, nil
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s, Author: %s, Pages: %d \n", b.title, b.author, b.pages)
}

func (b *Book) Settitle(title string) {
	b.title = title
}

func (b *Book) Gettitle() string {
	return b.title
}

type Textbook struct {
	Book      Book
	editorial string
	niveles   string
}

func NewtextBook(title, author string, pages int, editorial, niveles string) *Textbook {
	return &Textbook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		niveles:   niveles,
	}
}

func (t *Textbook) PrintInfo() {
	fmt.Printf("Title: %s, Author: %s, Pages: %d editorial: %s, niveles%s\n",
		t.Book.title, t.Book.author, t.Book.pages, t.editorial, t.niveles)
}
