package scale

type Code interface {
	Name() string
	Root() *Note
	Notes() []Note
}
