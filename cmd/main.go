package main

import (
	"fmt"
	"github.com/adamsafr/go-chatcord/pkg/chat"
	"github.com/googollee/go-socket.io"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
)

type Config struct {
	ServerPort string `env:"SERVER_PORT"`
}

func loadEnvConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(".env", &cfg)

	return &cfg, err
}

func main()  {
	cfg, err := loadEnvConfig()

	if err != nil {
		log.Fatal(err)
	}

	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	chat.InitChatEndpoints(server)

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./web/static")))

	log.Println(fmt.Sprintf("Serving at localhost:%s...", cfg.ServerPort))
	log.Fatal(http.ListenAndServe(":" + cfg.ServerPort, nil))
}
