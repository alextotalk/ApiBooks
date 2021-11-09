package ApiBooks

type Book struct {
	ID    int
	Title string
}
type BooksStorage struct {
	ID    int
	Books []Book
}

type UsersBooksStorage struct {
	ID             int
	UserID         int
	BooksStorageID int
}
