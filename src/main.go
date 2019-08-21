package main

import (
	"math/rand"
	"time"

	"github.com/agzuniverse/go-testcases/src/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cmd.Execute()
}
