package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)


type msg struct {
	clientKey string
	text string
}

type newClientEvent struct {
	clientKey string
	msgChan chan *msg
}

var dirPath string
var clientRequests = make(chan *newClientEvent, 100)
var clientDisconnects = make(chan string, 100)
var messages = make(chan *msg, 100)

func getFile(w http.ResponseWriter, req *http.Request, filename string) {
	fp, err := os.Open(dirPath + "/" + filename)

	if err != nil {
		log.Println("Could not open file", err.Error())
		w.Write([]byte("500 internal server error"))
		return
	}

	defer fp.Close()

    _, err = io.Copy(w, fp)

	if err != nil {
		log.Println("Could not open file", err.Error())
		w.Write([]byte("500 internal server error"))
		return
	}

}

func router() {
	clients := make(map[string]chan *msg)
	for {
		select {
			case req := <-clientRequests:
				fmt.Println(req)
				clients[req.clientKey] = req.msgChan
				log.Println("Websocket connected " +  req.clientKey)
			case clientKey:= <-clientDisconnects:
				close(clients[clientKey])
				delete(clients,clientKey)
				log.Println("Websocket disconnected " +  clientKey)
			case msg:= <-messages:
				fmt.Println(msg)
				for _, msgChan := range clients {
					if len(msgChan) < cap(msgChan) {
						msgChan <- msg
					}
				}

		}
	}
}

func chatServer(ws *websocket.Conn){
	var lenBuf = make([]byte, 5 )

	msgChan := make(chan *msg, 100)
	clientKey:= ws.RemoteAddr().String()
	clientRequests <-&newClientEvent{clientKey, msgChan}
	defer func() {clientDisconnects <-clientKey}()
	go func() {
		for msg:= range msgChan {
			ws.Write([]byte(msg.text))
		}
	}()

	for {
		_,err := ws.Read(lenBuf[:])
		if err != nil {
			log.Println("Error: ", err.Error())
			return
		}
		length,_ := strconv.Atoi(strings.TrimSpace(string(lenBuf)))

		if length > 65536 {
			log.Println("Error; too big length: ", length)
			return
		}

		if length <= 0 {
			log.Println("Empty length: ", length)
			return
		}

		buf := make([]byte, length)
		_,err = ws.Read(buf)

		if err != nil {
			log.Println("could not read", length, "bytes: ", err.Error())
			return
		}

		messages<-&msg{clientKey,string(buf)}
	}
}

func init(){

	if len(os.Args) < 2 {
		log.Fatal("Usage: chat <srcdir>")
	}

	dirPath = os.Args[1]
	fmt.Println("Starting...")

	go router()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			getFile(writer, request, "index.html")
	})

	http.HandleFunc("/index.js", func(writer http.ResponseWriter, request *http.Request) {
		getFile(writer, request, "index.js")
	})

	http.Handle("/ws",websocket.Handler(chatServer))
}

func main() {
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("listend and serve: ", err)
	}
}