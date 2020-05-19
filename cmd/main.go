package main

import (
	"github.com/phuwn/clia/cmd/fine"
	"github.com/phuwn/tools/log"
)

func main() {
	err := fine.New().Execute()
	if err != nil {
		log.Fatal(err)
	}
}
