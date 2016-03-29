package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/websocket"
)

const (
	fps = 30
	// frameMs = 1000 / fps // Duration of a frame in milli seconds
	frameNs = int64(1e9) / fps // Duration of a frame in nano seconds
)

var (
	listenAddr	= ":8080"
	homePage	= "./index.html"
)

var (
	// Declared in top level so that it is visible for the whole package
	mainGas *gas
)

func init() {
	// Declared in top level so that it is visible for the whole package
	mainGas = newGas()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Handler to /
	toByte, _ := ioutil.ReadFile(homePage)
	w.Write(toByte)
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func wsHandler (w http.ResponseWriter, r *http.Request) {
	// Handler to /ws
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	addr := ws.RemoteAddr()

	log.Printf("Websocket accepted: %s\n", addr)

	john := newPlayer(ws)

	// New player is added to the main gas
	mainGas.addPlayer(john)

	go john.writer()

	// se puede mover a addPlayer no?
	john.identity()

	john.reader()

	// Once reader returns the connection is finalized
	log.Printf("Websocket finalized: %s\n", addr)
}

func main() {

	// Register the handlers
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/ws", wsHandler)

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	// Note the trailing slash in /js/, it's a directory

	log.Printf("Listening on %s.\n", listenAddr)

	// (addr, handler) if handler is nil, default mutex is used
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
