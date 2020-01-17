package visualizer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Server struct {
	srv   *http.Server
	Start chan struct{}
}

type Event struct {
	Type   uint32 `json:"type"`
	Source string `json:"source"`
	Dest   string `json:"dest"`
}

var (
	clients   = make(map[*websocket.Conn]bool)
	wsChan    = make(chan *websocket.Conn, 1)
	eventChan = make(chan *Event, 100000)
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func NewServer() *Server {
	s := &Server{
		Start: make(chan struct{}, 1),
	}

	router := mux.NewRouter()
	router.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) { s.Start <- struct{}{} }).Methods("GET")
	router.HandleFunc("/event", eventHandler).Methods("POST")
	router.HandleFunc("/ws", wsHandler)
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("frontend").HTTPBox()))

	s.srv = &http.Server{Addr: ":8844", Handler: router}

	return s
}

func (s *Server) Run() {
	go echo()

	if err := s.srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %v", err)
	}
}

func (s *Server) Close() {
	close(eventChan)
	if err := s.srv.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
}

func Writer(event *Event) {
	eventChan <- event
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		log.Printf("ERROR: %s", err)
		http.Error(w, "Bad request", http.StatusTeapot)
		return
	}
	defer r.Body.Close()
	go Writer(&event)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// register client
	wsChan <- ws
}

func echo() {
	for val := range eventChan {
		if val.Type <= removeLink {
			time.Sleep(50 * time.Millisecond)
		}
		event := fmt.Sprintf("%d %s %s", val.Type, val.Source, val.Dest)

		// check for new clients
		select {
		case ws := <-wsChan:
			clients[ws] = true
		default:
		}

		// send to every client that is currently connected
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(event))
			if err != nil {
				log.Printf("Websocket error: %s", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
