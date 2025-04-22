package contact

type Contact struct {
	ID    int    `json:"-"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type ContactResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}