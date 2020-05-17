package main

import (
	"fmt"
	"github.com/adamsafr/go-chatcord/pkg/chat"
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

const port = "5000"

func main()  {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	chat.InitChatEndpoints(server)

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./web/static")))

	log.Println(fmt.Sprintf("Serving at localhost:%s...", port))
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
