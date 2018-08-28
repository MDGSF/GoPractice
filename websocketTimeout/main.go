package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/MDGSF/utils/log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func ServeIndex(w http.ResponseWriter, r *http.Request) {

	log.Info("new client comming")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic("upgrade failed")
	}

	go func() {

		conn.SetPongHandler(func(string) error {
			log.Info("server get pong from client")
			return nil
		})

		for {
			mt, message, err := conn.ReadMessage()
			log.Info("server read message", mt, message, err)
		}
	}()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			log.Info("server send one ping to client")
			conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
			conn.WriteMessage(websocket.PingMessage, nil)
		}
	}
}

func main() {

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", ServeIndex)

		httpServer := &http.Server{}
		httpServer.Addr = "127.0.0.1:11111"
		httpServer.Handler = mux
		httpServer.ListenAndServe()
	}()

	time.Sleep(time.Second)

	u := url.URL{
		Scheme: "ws",
		Host:   "127.0.0.1:11111",
		Path:   "/",
	}
	d := websocket.Dialer{}
	conn, _, err := d.Dial(u.String(), nil)
	if err != nil {
		panic("client dial failed")
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	log.Info("client connect to server success")

	go func() {

		conn.SetPongHandler(func(string) error {
			log.Info("client get pong from server")
			return nil
		})

		for {
			mt, message, err := conn.ReadMessage()
			log.Info("client read message", mt, message, err)
		}
	}()

	ticker := time.NewTicker(7 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			log.Info("client send one ping to server")
			conn.SetWriteDeadline(time.Now().Add(3 * time.Second))
			conn.WriteMessage(websocket.PingMessage, nil)
		}
	}
}
