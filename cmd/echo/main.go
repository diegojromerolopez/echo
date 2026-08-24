package main

import (
	"fmt"
	"os"

	"github.com/noctifab/echo/pkg/echoer"
)

func main() {
	res := echoer.Echo(os.Args[1:])
	fmt.Println(res)
}
