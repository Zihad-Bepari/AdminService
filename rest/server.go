package rest

import (
	"admin-service/config"
	"admin-service/rest/handlers/user"
	"admin-service/rest/middleware"
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	cnf         *config.Config
	userHandler *user.Handler
}

func Newserver(cnf *config.Config, userHandler *user.Handler) *Server {
	return &Server{
		cnf:         cnf,
		userHandler: userHandler,
	}
}

func (server *Server) Start() {

	mux := http.NewServeMux()

	manager := middleware.Newmanager()

	wrappedmux := manager.Wrapmux(mux,
		middleware.Logger,
	)

	server.userHandler.RegisterRoute(mux, manager)

	//InitRoute(mux, manager)
	port := ":" + string(server.cnf.Httpport)
	println("🚀 Server is running at http://localhost" + port)
	err := http.ListenAndServe(port, wrappedmux)
	if err != nil {
		fmt.Println("Error starting the server")
		os.Exit(1)
	}
}
