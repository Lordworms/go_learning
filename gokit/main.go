package main

type HelloService interface {
	Greeting(string) (string, error)
}
