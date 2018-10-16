package router

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/sysu-go-online/service-end/controller"
	"github.com/urfave/negroni"
)

var upgrader = websocket.Upgrader{}

// GetServer return web server
func GetServer() *negroni.Negroni {
	r := mux.NewRouter()

	r.Handle("", controller.ErrorHandler(controller.CreateUserHandler)).Methods("POST")
	r.Handle("/", controller.ErrorHandler(controller.CreateUserHandler)).Methods("POST")
	r.Handle("/{username}", controller.ErrorHandler(controller.GetUserMessageHandler)).Methods("GET")

	// Use classic server and return it
	handler := cors.Default().Handler(r)
	s := negroni.Classic()
	s.UseHandler(handler)
	return s
}
