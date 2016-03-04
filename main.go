package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"github.com/gorilla/websocket"
)

const (
	fps = 30
	frameS = 1 / fps
	// frameMs = 1000 / fps
	frameNs = int(1e9) / fps
)

var (
	listenAddr	= ":8080"
	homePage	= "./index.html"
)

var (
	// Declarado en top level para que sea visible para todo el paquete
	mainGas *gas
)

func init() {
	// Declarado en top level para que sea visible para todo el paquete
	mainGas = newGas()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	toByte, _ := ioutil.ReadFile(homePage)
	w.Write(toByte)
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func wsHandler (w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	addr := ws.RemoteAddr()

	log.Printf("Websocket accepted: %s\n", addr)

	// Nuevo jugador
	john := newPlayer(ws)

	mainGas.addPlayer(john)

	go john.writer()

	// se puede mover a addPlayer no?
	john.identity()

	john.reader()

	// Una vez que reader regresa:
	log.Printf("Websocket finalized: %s\n", addr)
}

func main() {

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/ws", wsHandler)

	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))

	log.Printf("Listening on %s.\n", listenAddr)

	// (addr, handler) si handler es nil, se usa default mutex
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
