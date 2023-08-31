package main

import (
	"fmt"
	arithmetic "hex/internal/adapters/core"
	"hex/internal/ports"
	"log"
	"time"
)

// "internal/adapters/core/arithmetic"

func main() {
	t := time.Now()
	defer func() {
		log.Println("Executed in ", time.Since(t))
	}()
	var core ports.ArithmeticPort

	core = arithmetic.NewAdapter()
	fmt.Println(core.Addition(1, 2))
}
