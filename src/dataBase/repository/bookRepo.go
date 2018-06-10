package repository

type BookRepository interface {
	GetAll() (books []Book, err error)
	GetByCategory(categoryID int) (books []Book, err error)
	GetByID(bookID int) (books Book, err error)
	GetMostPopularBooks(quantity int) ([]Book, error)
	InsertNewBook(b Book) (err error)
	GetAllTakenBooks() (books []Book, err error)
	UpdateBookState(bookId int, state string) (err error)
	UpdateBookStateAndUsersBookIdByUserEmail(email string, state string, bookId int) (err error)
	GetFreeBooks(books []Book)([]Book, error)
}