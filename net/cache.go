package net

// TODO:
//     - implement cache interface and structs
//     - maybe create a response object that stores the response and time retrieved, can be used for caching and rate limiting
//       the cache should be checked at the net level (data provider) not at the service level (still needs to think about this)
//     - current option is use LevelDB
//     - private struct and fields because will be used at net layer

type connectionConfig struct {
	connectionString string
	timeout          int
	maxRetries       int
}

type cache interface {
	connect(c connectionConfig) error
	get(key string) (string, error)
	put(key string) (string, error)
	delete(key string) error
}
