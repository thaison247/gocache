package my_gocache

type ICache interface {
	// Connect to cache server
	Connect() error

	//
	Set(key string, value interface{}, exprireTime ...int) error

	//
	Get(key string) (interface{}, error)

	// delete a key
	Delete(key string) (int64, error)

	//
	Close()

	// set expire time on a key
	Expire(key string, expireTime int) (interface{}, error)

	// set expire
	ExpireV2(key string, expireTime int) error

	// get remain time life (expiretime)
	GetRemainLifeTime(key string) (int64, error)
}
