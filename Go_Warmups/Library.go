package main

import (
  "fmt"
)


// Create a library
// a library is a slice of books
// books contain the values of a title, author, and rating
// implement a way to insert a book into the library (if it doesn't exist)
// implement a way to print each book in the library
// implement a way to return a map of books by the author (i.e. the key is the author's name and the value is all books they've written)


type Book struct {
  Title string
  Author string
  Rating uint8 //scale of 1-5
}

var Library []Book


func insertBook ( Title string, Author string, Rating uint8 ) {
  for _, value := range Library {
    if ( value.Title == Title && value.Author == Author ) {
      fmt.Printf("%s by %s is already in the library\n", Title, Author)
      return
    }
  }

  Library = append(Library, Book{ Title: Title, Author: Author, Rating: Rating })
}

func printLibrary() {
  fmt.Printf("The Library contains these books\n----------------------------\n")
  for _, value := range Library {
    fmt.Printf("%s by %s\n", value.Title, value.Author)
  }
}

func groupBooksByAuthor() map[string][]Book {
  var myMap = make( map[string][]Book )

  for _, value := range Library {
    myMap[value.Author] = append( myMap[value.Author], value)
  }

  return myMap
}

func printLibraryByAuthor( mapping map[string][]Book ) {
  for authorName, bookValue := range mapping {
    fmt.Printf("%s has these books in the Library\n----------------------------\n", authorName)
    for _, value := range bookValue {
      fmt.Printf("%s with a rating of %d\n", value.Title, value.Rating)
    }
    fmt.Printf("----------------------------\n")
  }
}

func main() {

  insertBook("Book1", "Logan Cover", 5)
  insertBook("Book2", "Logan Cover", 4)
  insertBook("Book3", "Logan Cover", 4)
  insertBook("Book4", "Logan Cover", 1)

  insertBook("Book5", "Logan Lover", 1)
  insertBook("Book5", "Logan Lover", 1)

  printLibrary()

  printLibraryByAuthor( groupBooksByAuthor() )
}
