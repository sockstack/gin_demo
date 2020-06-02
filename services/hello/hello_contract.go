package hello

type HelloContract interface {
	SayHello(username string) string
}