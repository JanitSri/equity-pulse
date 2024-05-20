package net

// TODO:
//     - implement cache interface and structs
//     - maybe create a response object that stores the response and time retrieved, can be used for caching and rate limiting
//       the cache should be checked at the net level (data provider) not at the service level (still needs to think about this)
//     - current option is use LevelDB
//     - private struct and fields because will be used at net layer

type ConnectionConfig struct {
	ConnectionString string
	Timeout          int
	MaxRetries       int
}

type Cache interface {
	connect(c ConnectionConfig) error
	get(key string) (string, error)
	put(key string) (string, error)
	delete(key string) error
}
