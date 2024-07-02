package models

import (
	"fmt"
	"strings"

	"github.com/v1bh475u/LibMan_MVC/pkg/types"
)

func FetchBooks(title, author, genre string, ID int) []types.Book {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return nil
	}
	defer db.Close()

	query := `SELECT * FROM books`
	conditions := []string{}
	params := []interface{}{}
	if ID != 0 {
		conditions = append(conditions, `BookID = ?`)
		params = append(params, ID)
	}
	if title != "" {
		conditions = append(conditions, `Title LIKE %?%`)
		params = append(params, title)
	}
	if author != "" {
		conditions = append(conditions, `Author = ?`)
		params = append(params, author)
	}
	if genre != "" {
		conditions = append(conditions, `Genre = ?`)
		params = append(params, genre)
	}
	if len(conditions) > 0 {
		query += ` WHERE ` + strings.Join(conditions, ` AND `)
	}
	result, err := db.Query(query, params...)
	if err != nil {
		fmt.Printf("Error querying database: %v", err)
		return nil
	}
	var books []types.Book
	for result.Next() {
		var book types.Book
		err := result.Scan(&book.BookID, &book.Title, &book.Author, &book.Genre, &book.Quantity)
		if err != nil {
			fmt.Printf("Error scanning database: %v", err)
			return nil
		}
		books = append(books, book)
	}
	return books
}

func UpdateBook(Quantity, ID int) error {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return err
	}
	defer db.Close()

	updateBook := `UPDATE books SET Quantity = ? WHERE BookID = ?`
	_, err = db.Exec(updateBook, Quantity, ID)
	if err != nil {
		fmt.Printf("Error updating database: %v", err)
		return err
	}
	return nil
}

func InsertBook(book types.Book) error {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return err
	}
	defer db.Close()

	insertBook := `INSERT INTO books (Title, Author, Genre, Quantity) VALUES (?, ?, ?, ?)`
	_, err = db.Exec(insertBook, book.Title, book.Author, book.Genre, book.Quantity)
	if err != nil {
		fmt.Printf("Error inserting into database: %v", err)
		return err
	}
	return nil
}

func FetchUniqueitems(property string) []string {
	db, err := connection()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		return nil
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT %s FROM books GROUP BY %s", property, property)
	result, err := db.Query(query)
	if err != nil {
		fmt.Printf("Error querying database: %v", err)
		return nil
	}
	var items []string
	for result.Next() {
		var item string
		err := result.Scan(&item)
		if err != nil {
			fmt.Printf("Error scanning database: %v", err)
			return nil
		}
		items = append(items, item)
	}
	return items
}
