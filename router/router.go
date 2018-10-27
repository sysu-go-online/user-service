package router

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/sysu-go-online/public-service/types"
	"github.com/sysu-go-online/user-service/controller"
	"github.com/urfave/negroni"
)

var upgrader = websocket.Upgrader{}

// GetServer return web server
func GetServer() *negroni.Negroni {
	r := mux.NewRouter()

	r.Handle("/users", types.ErrorHandler(controller.CreateUserHandler)).Methods("POST")
	r.Handle("/users/{username}", types.ErrorHandler(controller.GetUserMessageHandler)).Methods("GET")
	r.Handle("/users/{username}/files", types.ErrorHandler(controller.GetUserHomeStructureHandler)).Methods("GET")

	// Use classic server and return it
	handler := cors.Default().Handler(r)
	s := negroni.Classic()
	s.UseHandler(handler)
	return s
}
