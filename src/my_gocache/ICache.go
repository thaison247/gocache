package my_gocache

type ICache interface {

	// Connect to cache server
	Connect() error

	// Close cache server connection
	Close()

	// set a pair of key-value
	// params: key (string)
	//		   value (interface{})
	//		   expireTime (int) - [optional]
	// return: error
	Set(key string, value interface{}, exprireTime ...int) error

	// get a pair of key-value
	// param: key (string)
	// returns: value in pair of key-value (interface{})
	//			error (error)
	Get(key string) (interface{}, error)

	// delete a pair of key-value
	// param: key (string)
	// returns: number of deleted key-value (1 | 0): int64
	//			error: error
	Delete(key string) (int64, error)

	// set expire time on a key
	// params: key (string)
	//		   expireTime (int)
	// returns: number of affected key (1 | 0):
	Expire(key string, expireTime int) (int64, error)

	// get remain time life (expiretime)
	// param: key (string)
	// returns: remaining lifetime (int64)
	//			error
	GetRemainLifeTime(key string) (int64, error)

	// set expire
	// ExpireV2(key string, expireTime int) error
}
