package stringsfmt

type NameCreator interface {
	CreateName(name string) string
}
