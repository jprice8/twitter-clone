package server

import (
	"github.com/jprice8/twitter-clone/internal/router"
	"github.com/jprice8/twitter-clone/internal/shared/database"
	"github.com/jprice8/twitter-clone/internal/shared/webserver"
)

type Server struct {
	db database.Database
	webserver webserver.WebServer
}

// New creates a new instance of Fiber web server
func New(webserver webserver.WebServer, db database.Database) *Server {
	return &Server {
		webserver: webserver,
		db: db,
	}
}

func (s *Server) Listen() {
	s.initRouteGroups()
	s.webserver.Listen(":8080")
}

func (s *Server) initRouteGroups() {
	// router.authRouter(s.webserver.Engine().Group("/auth"), s.db)
	router.ApiRoutes(s.webserver.Engine().Group("/api"), s.db)
	router.AuthRoutes(s.webserver.Engine().Group("/auth"), s.db)
	router.UserRoutes(s.webserver.Engine().Group("/user"), s.db)
}