package ApiBooks

type Book struct {
	ID    int
	Title string
}

type UsersBooks struct {
	ID     int
	UserID int
	BookID int
}
