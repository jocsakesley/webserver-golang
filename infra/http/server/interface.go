package server

type IContext interface {
	Param(key string) string
	JSON(code int, object any) error
	BodyParser(out interface{}) error
}

type IServer interface {
	Get(path string, handler interface{})
	Post(path string, handler interface{})
	Run(port string)
}
