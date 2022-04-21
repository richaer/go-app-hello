package main

import (
	"fmt"

	go_module "github.com/richaer/go-module/v2"
)

func main() {
	name := "Quamala"
	fmt.Println(go_module.SayHello(name))
}
