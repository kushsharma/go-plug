package transformer

type Interface interface {
	Name() (string, error)
	Description() (string, error)
	GenerateDependencies(string) ([]string, error)
}
