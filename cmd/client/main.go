package main

import (
	"flag"
	"log"

	"github.com/UfaSemen/testMoneyOperations/pkg/client"
)

func main() {
	servAddr := flag.String("c", "http://localhost:8080", "server path")
	operation := flag.String("o", "", "operation")
	flag.Parse()
	err := client.ExecuteClient(*operation, *servAddr)
	if err != nil {
		log.Fatal(err)
	}
}
