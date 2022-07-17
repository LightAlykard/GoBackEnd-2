package server

import (
	"context"
	"net/http"
	"time"

	"github.com/LightAlykard/GoLibsTest/logic"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	VersionInfo
	port  string
	Tr    trace.Tracer
	Logic logic.Logic
}

type VersionInfo struct {
	Version string
	Commit  string
	Build   string
}

func New(info VersionInfo, port string, tr trace.Tracer, lg logic.Logic) *Server {
	return &Server{
		VersionInfo: info,
		port:        port,
		Tr:          tr,
		Logic: logic.Logic{
			Tr: tr,
		},
	}
}

func (s Server) Serve(ctx context.Context) error {
	e := echo.New()
	e.Use(otelecho.Middleware("my-server"))
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Recover())
	s.initHandlers(e, ctx)
	go func() {
		e.Logger.Infof("start server on port: %s", s.port)
		err := e.Start(":" + s.port)
		if err != nil {
			e.Logger.Errorf("start server error: %v", err)
		}
	}()
	<-ctx.Done()
	return e.Shutdown(ctx)
}

func (s Server) initHandlers(e *echo.Echo, ctx context.Context) {
	e.GET("/", handler)
	e.GET("/__heartbeat_", heartbeatHandler)
	e.GET("/__version__", s.versionHandler)
	//e.GET("/entities", s.simpleHander2(e.Context, ctx))
	//r.HandleFunc("/create", r.AuthMiddleware(http.HandlerFunc(r.CreateLink)).ServeHTTP)
	//e.HandleFunc("/entities/", s.simpleHander()).Methods(http.MethodPost)

	e.Any("/entities", func(c echo.Context) error {
		ctx, span := s.Tr.Start(ctx, "foo")
		defer span.End()

		data := s.Logic.Example(ctx)

		if time.Now().Second()%2 == 0 {
			return c.String(http.StatusOK, data)
		} else {
			return c.String(http.StatusInternalServerError, data)
		}

		//return c.NoContent(http.StatusNotFound)
	})

	e.Any("/*", func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	})
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World! Welcome to GeekBrains!\n")
}

func heartbeatHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
func (s Server) versionHandler(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		map[string]string{
			"version": s.VersionInfo.Version,
			"commit":  s.VersionInfo.Commit,
			"build":   s.VersionInfo.Build,
		},
	)
}
