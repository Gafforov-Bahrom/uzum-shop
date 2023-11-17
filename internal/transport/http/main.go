package http

import (
	"context"
	"net/http"
	"sync"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/transport/http/shop"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	server *http.Server
	engine *gin.Engine
}

func NewServer(port string, engine *gin.Engine, shopRouter *shop.Router) *Server {
	server := &Server{
		engine: engine,
		server: &http.Server{
			Addr: port,
		},
	}
	server.setShopRoutes(shopRouter)
	return server
}

func (s *Server) Start(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	s.server.Handler = s.engine
	go func() {
		<-ctx.Done()
		logrus.Error("shutting down http server")
		logrus.Error(s.server.Shutdown(ctx))
	}()
	return s.server.ListenAndServe()
}

func (s *Server) setShopRoutes(shopRouter *shop.Router) {
	shop := s.engine.Group("/shop")
	shop.GET("/v1/product/:product_id", shopRouter.GetProduct)
	shop.GET("/v1/products", shopRouter.GetProducts)
	shop.POST("/v1/basket", shopRouter.AddProduct)
}
