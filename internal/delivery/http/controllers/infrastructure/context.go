package infrastructure

type Context interface {
	BindQuery(obj any) error
	String(code int, format string, values ...any)
}
