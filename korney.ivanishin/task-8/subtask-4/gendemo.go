package gendemo

//go:generate mockgen -destination=artifacts/mock.go -package=gendemo . Doer

type Doer interface {
        DoSomething(int, string) error
}
