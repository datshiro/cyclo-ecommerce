package server

import (
	"fmt"
	"net/http"

	"github.com/datshiro/cyclo-ecommerce/internal/consts"
	"github.com/datshiro/cyclo-ecommerce/internal/domain"
	"github.com/datshiro/cyclo-ecommerce/internal/infras/server/config"
	"github.com/datshiro/cyclo-ecommerce/internal/infras/server/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func New(config config.Config) Server {
	return Server{Echo: echo.New(), config: config}
}

type Server struct {
	*echo.Echo
	config config.Config
}

func (s *Server) Start() error {
	s.Config()
	handlers.RegisterHandlers(s.Echo, s.config.ApiPath, s.config)

	return s.Echo.Start(s.config.Address())
}

func (s *Server) Config() error {
	s.configMiddleware()
	s.configErrHandler()
	return nil
}

func (s *Server) configMiddleware() {
	// Middleware
	s.configLoggerMiddleware()
	s.Echo.Use(middleware.Logger())
	s.Echo.Use(middleware.Recover())
}

func (s *Server) configLoggerMiddleware() {
	log := logrus.New()
	s.Echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")

			return nil
		},
	}))
}

//This function will add associated response error
func (s *Server) configErrHandler() {
	defaultHandler := s.Echo.HTTPErrorHandler
	s.Echo.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.IsWebSocket() {
			return // connection is hijacked, can't write to response anymore
		}

		if he, ok := err.(*echo.HTTPError); ok {
			if he.Internal != nil {
				err = fmt.Errorf("%v, %v", err, he.Internal)
			}
			c.JSON(he.Code, domain.ErrorResponse{Message: he.Message})
			return
		} else if he, ok := err.(consts.ServerErr); ok {
			err := &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: domain.ErrorResponse{Message: he.Error()},
			}
			c.JSON(err.Code, err.Message)
			return
		}

		// c.NoContent(http.StatusInternalServerError)
		defaultHandler(err, c)
	}
}
