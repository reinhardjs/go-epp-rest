package infrastructure

type Context interface {
	String(code int, format string, values ...any)
}
