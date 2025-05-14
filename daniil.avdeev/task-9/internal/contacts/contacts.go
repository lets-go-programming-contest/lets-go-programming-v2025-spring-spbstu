package contacts

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type DBContact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
