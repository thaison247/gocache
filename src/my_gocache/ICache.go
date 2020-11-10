package my_gocache

type ICache interface {
	// Connect to cache server
	Connect() error

	//
	Set(key string, value interface{}) error

	//
	Get(key string) (interface{}, error)

	//
	Delete(key string) error
}
