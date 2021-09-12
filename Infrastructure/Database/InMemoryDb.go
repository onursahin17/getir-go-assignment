package Database

// Interface for in memory database
type InMemoryDb interface {
	Get(key string) (*string, error)
	Set(key, val string) error
}