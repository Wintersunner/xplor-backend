package api

import (
	"fmt"
	db "github.com/Wintersunner/xplor/db/sqlc"
	"github.com/Wintersunner/xplor/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Server struct {
	store  db.Store
	router *gin.Engine
	config util.Config
}

func NewServer(store db.Store, config util.Config) *Server {
	server := &Server{
		store:  store,
		config: config,
	}

	server.setupRouter()
	return server
}

func (server *Server) Start() error {
	return server.router.Run(fmt.Sprintf("%s:%s", server.config.Host, server.config.Port))
}

func (server *Server) setupRouter() {
	gin.SetMode(server.config.GinMode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(server.config.AllowedOrigins, ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Content-Type", "Accept", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))

	router.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "not found 404!",
		})
	})

	router.GET("/", healthHandler)

	router.GET("fizzbuzz", server.listFizzBuzzes)
	router.GET("fizzbuzz/:id", server.getFizzBuzz)
	router.POST("fizzbuzz", server.createFizzBuzz)

	server.router = router
}

func healthHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "healthy as a horse!",
	})
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
