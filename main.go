package main

import (
	_ "github.com/alioth-center/foreseen/router"
	"github.com/alioth-center/infrastructure/exit"
)

func main() {
	exit.BlockedUntilTerminate()
}
