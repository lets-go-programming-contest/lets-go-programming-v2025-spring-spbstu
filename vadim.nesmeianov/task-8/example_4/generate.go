package generate

//go:generate mockgen -destination=mock_foo.go -package=generate . Doer

type Doer interface {
	DoSomething(int, string) error
}
