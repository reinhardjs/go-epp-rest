package infrastructure

type Context interface {
	BindQuery(obj any) error
	Query(key string) string
	AbortWithError(code int, fatalErr error)
}
