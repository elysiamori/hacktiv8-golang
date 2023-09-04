package main

import (
	"fmt"
	"strings"
)

type Book struct {
	title         string
	publishedYear int
	author        string
	publisher     Publisher
}

func (b *Book) PrintInformation(informationBanner string) {

	fmt.Println()
	fmt.Println(strings.ToUpper(informationBanner))
	fmt.Println()

	fmt.Println("author:", b.author)
	fmt.Println("title:", b.title)
	fmt.Println("publishedYear:", b.publishedYear)
	fmt.Println("publisher name:", b.publisher.name)
	fmt.Println("publisher reputation:", b.publisher.reputation)

}

type Publisher struct {
	name       string
	reputation int
}

func sctStruct() {
	var book1 Book = Book{}
	fmt.Println("author:", book1.author)
	fmt.Println("title:", book1.title)
	fmt.Println("publishedYear:", book1.publishedYear)

	book1.author = "dominique"
	book1.title = "story of dominique"
	book1.publishedYear = 2003
	book1.publisher.name = "apa"
	book1.publisher.reputation = 1

	fmt.Println("author:", book1.author)
	fmt.Println("title:", book1.title)
	fmt.Println("publishedYear:", book1.publishedYear)
	fmt.Println("publisher name:", book1.publisher.name)
	fmt.Println("publisher reputation:", book1.publisher.reputation)

	fmt.Println()
	fmt.Println("BOOK 2")
	var book2 Book = Book{
		author:        "domjeff",
		title:         "book of dom",
		publishedYear: 2023,
		publisher: Publisher{
			name:       "alamaro printing",
			reputation: 10,
		},
	}
	fmt.Println("author:", book2.author)
	fmt.Println("title:", book2.title)
	fmt.Println("publishedYear:", book2.publishedYear)
	fmt.Println("publisher name:", book2.publisher.name)
	fmt.Println("publisher reputation:", book2.publisher.reputation)

	var pointerToBook2 *Book
	pointerToBook2 = &book2
	fmt.Println(pointerToBook2)

	var specificBook = struct {
		bookId string
		book   Book
	}{}

	specificBook.bookId = "id-00001"
	specificBook.book = book2

	fmt.Println()
	fmt.Println("SPECIFIC BOOK")
	fmt.Println("book id", specificBook.bookId)
	fmt.Println("author:", specificBook.book.author)
	fmt.Println("title:", specificBook.book.title)
	fmt.Println("publishedYear:", specificBook.book.publishedYear)
	fmt.Println("publisher name:", specificBook.book.publisher.name)
	fmt.Println("publisher reputation:", specificBook.book.publisher.reputation)

	book3 := Book{
		title:         "3rd book of dominique",
		publishedYear: 2023,
		author:        "alamaro",
		publisher: Publisher{
			name:       "some other printing",
			reputation: 7,
		},
	}
	books := []Book{
		book1,
		book2,
		book3,
	}

	for i, book := range books {
		fmt.Println()
		fmt.Printf("BOOK OF INDEX %d\n", i+1)

		fmt.Println("author:", book.author)
		fmt.Println("title:", book.title)
		fmt.Println("publishedYear:", book.publishedYear)
		fmt.Println("publisher name:", book.publisher.name)
		fmt.Println("publisher reputation:", book.publisher.reputation)
	}

	// string di sini berlaku sebagai key yang akan diisi dengan judul buku (karena kita menggunakan asumsi judul buku unik)
	mapOfBooks := map[string]Book{}
	mapOfBooks[book1.title] = book1
	mapOfBooks[book2.title] = book2
	mapOfBooks[book3.title] = book3

	for key, val := range mapOfBooks {
		fmt.Println()
		fmt.Printf("BOOK OF TILE  \"%s\"\n", key)

		fmt.Println("author:", val.author)
		fmt.Println("title:", val.title)
		fmt.Println("publishedYear:", val.publishedYear)
		fmt.Println("publisher name:", val.publisher.name)
		fmt.Println("publisher reputation:", val.publisher.reputation)
	}

	bookFound, _ := findTitleWithStruct("3rd book of dominique", books)
	bookFound.PrintInformation("BOOK FOUND BY ITERATING SLICE")

	bookFound = mapOfBooks["3rd book of dominique"]
	bookFound.PrintInformation("BOOK FOUND BY LOOKING AT HASH")

	bookBaru := &Book{
		title:         "new book",
		publishedYear: 2023,
		author:        "someone",
	}
	bookBaru.PrintInformation("PRINT NEW BOOK")

}

func findTitleWithStruct(in string, books []Book) (Book, bool) {
	for i, book := range books {
		fmt.Println()
		fmt.Printf("ITERASI KE %d\n", i+1)
		fmt.Println()
		if book.title == in {
			return book, true
		}
	}
	return Book{}, false
}

func withPointer(*Book) {

}
