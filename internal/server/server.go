package server

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	api := s.webserver.Engine().Group("/api", logger.New())
	router.ApiRoutes(api.Group("/"), s.db)
	router.AuthRoutes(api.Group("/auth"), s.db)
	router.UserRoutes(api.Group("/user"), s.db)
	router.ProductRoutes(api.Group("/product"), s.db)
	router.CategoryRoutes(api.Group("/category"), s.db)
	router.OrderRoutes(api.Group("/order"), s.db)
}