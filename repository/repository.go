package repository

// Connector interface is describing the bahaviour of DB connectros
type Connector interface {
	Connect() error
	Close()
	Fetch(uuid string) ([]byte, error)
	Put(key string, data []byte) error
}
