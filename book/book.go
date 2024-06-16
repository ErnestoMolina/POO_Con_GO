package book

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
}

type PrintTable interface {
	PrintInfo()
}

func Print(p PrintTable) {
	p.PrintInfo()
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func (b *Book) GetPages() int {
	return b.pages
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s,\nAutor: %s,\nPaginas: %d.\n", b.title, b.author, b.pages)
}

type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title, author string, pages int, editorial, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Titulo: %s,\nAutor: %s,\nPaginas: %d,\nEditorial: %s,\nNivel: %s.\n",
		b.title, b.author, b.pages, b.editorial, b.level)
}
