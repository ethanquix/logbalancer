package logbalancer

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
	"github.com/labstack/echo/v4"
	"github.com/ucarion/urlpath"
)

type logBalancerTarget struct {
	fn      func(l *pb_logs.RuntimeLogs) error
	rawPath string
	path    urlpath.Path
}

type LogBalancer struct {
	port          string
	password      string
	middleware    http.HandlerFunc
	customHandles []func(*echo.Echo)
	listeners     map[string][]logBalancerTarget
	echoServer    *echo.Echo
}

type Opts func(*LogBalancer)

func WithPassword(password string) Opts {
	return func(balancer *LogBalancer) {
		balancer.password = password
	}
}

func WithPort(port string) Opts {
	return func(balancer *LogBalancer) {
		balancer.port = port
	}
}

func WithHandle(fn func(e *echo.Echo)) Opts {
	return func(balancer *LogBalancer) {
		balancer.customHandles = append(balancer.customHandles, fn)
	}
}

func New(options ...Opts) *LogBalancer {
	lb := &LogBalancer{
		listeners: make(map[string][]logBalancerTarget),
	}

	for _, o := range options {
		o(lb)
	}

	return lb
}

func (lb *LogBalancer) On(path string, fn ...func(incomingLog *pb_logs.RuntimeLogs) error) *LogBalancer {
	if strings.HasSuffix(path, "/") {
		path += "*"
	}
	uPath := urlpath.New(path)
	for _, f := range fn {
		var target logBalancerTarget
		target.path = uPath
		target.rawPath = path
		target.fn = f
		lb.listeners[path] = append(lb.listeners[path], target)
	}
	return lb
}

func (lb *LogBalancer) Run() error {
	// create server
	e := echo.New()
	lb.echoServer = e
	e.HideBanner = true

	if lb.password != "" {
		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				header := c.Request().Header.Get(echo.HeaderAuthorization)
				if header == lb.password {
					return next(c)
				} else {
					return c.String(http.StatusUnauthorized, "unauthorized. Please provide a password in the Authorization header. Example: Authorization: password")
				}
			}
		})
	}

	// Routes
	e.POST("/json", func(c echo.Context) error {
		return HandleJson(lb, c)
	})
	e.POST("/proto", func(c echo.Context) error {
		return HandleProto(lb, c)
	})
	e.Any("/connect/*", echo.WrapHandler(http.StripPrefix("/connect", HandleConnect(lb))))

	e.Any("/", echo.HandlerFunc(func(c echo.Context) error {
		return c.String(http.StatusOK, "logbalancer is running. Try sending a log to /json or /proto. You can also send logs to /connect/ with the connect protocol")
	}))

	// Custom
	for _, handle := range lb.customHandles {
		handle(e)
	}

	if lb.port != "" {
		return e.Start(fmt.Sprintf(":%s", lb.port))
	}
	return e.Start(":8080")
}

func (lb *LogBalancer) Stop() error {
	if lb.echoServer == nil {
		return fmt.Errorf("server was never started. Please call Run() first")
	}
	return lb.echoServer.Close()
}
