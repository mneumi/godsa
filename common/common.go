package common

type Element[T any] interface {
	Less(e T) bool
	Equal(e T) bool
	Grater(e T) bool
}
