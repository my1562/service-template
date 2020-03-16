package main

import (
	"fmt"

	"github.com/my1562/service-template/config"
	"go.uber.org/dig"
)

func main() {

	c := dig.New()
	c.Provide(config.NewConfig)

	err := c.Invoke(func(config *config.Config) {
		fmt.Println("Hello world")
		fmt.Println("Replace 'github.com/my1562/service-template' to your URL")

	})
	if err != nil {
		panic(err)
	}
}
