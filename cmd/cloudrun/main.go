package main

import (
	"os"

	"github.com/ethanquix/logbalancer/pkg/lbdestinations"
	"github.com/ethanquix/logbalancer/pkg/logbalancer"
)

func main() {
	lb := logbalancer.New(logbalancer.WithPort(os.Getenv("PORT")), logbalancer.WithPassword(os.Getenv("password")))

	lb.On("/", lbdestinations.StdoutSend)

	err := lb.Run()
	if err != nil {
		panic(err)
	}
}
