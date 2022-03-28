package main

import (
	"flag"
	"log"

	"github.com/UfaSemen/testMoneyOperations/pkg/server"
)

func main() {
	confPath := flag.String("c", "config.toml", "configuration file path")
	flag.Parse()
	config, err := server.ReadConfig(*confPath)
	if err != nil {
		log.Fatal("decode of config file:", err)
	}

	pbc, err := server.NewPostgreBalansController(config.Postgre)
	if err != nil {
		log.Fatal("can't connect to database:", err)
	}
	defer pbc.Close()
	server.StartServer(config.Port, pbc)
}
