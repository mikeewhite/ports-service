package ports

// PortRepository defines a repository for managing ports
type PortRepository interface {
	Add([]Port) error
	Get(key string) (*Port, error)
	List() ([]Port, error)
}
