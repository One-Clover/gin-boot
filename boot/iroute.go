package boot

type IRoute interface {
	Build(boot *Boot)
	Name() string
}
