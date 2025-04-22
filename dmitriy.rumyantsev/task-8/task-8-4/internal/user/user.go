package user

//go:generate go run go.uber.org/mock/mockgen@latest -destination=../../internal/mocks/mock_user.go -package=mocks . User

type User interface {
    GetUser(id int) (string, error)
    SaveUser(name string) error
}
