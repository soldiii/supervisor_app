package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/soldiii/supervisor_app/internal/server"
)

var (
	serverConfPath string
)

func init() {
	flag.StringVar(&serverConfPath, "serv-conf-path", "configs/server.toml", "path to server config file")
}

func main() {
	flag.Parse()
	serverConfig := server.NewServerConfig()
	_, err := toml.DecodeFile(serverConfPath, serverConfig)
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(serverConfig)
	if err := server.RunServer(); err != nil {
		log.Fatal(err)
	}
}
