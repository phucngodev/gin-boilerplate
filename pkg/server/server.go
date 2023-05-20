package server

import (
	"apiserver/pkg/config"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HttpServer struct {
	config     *config.Config
	logger     *zap.Logger
	onShutdown func()
}

func NewHttpServer(config *config.Config, logger *zap.Logger) *HttpServer {
	return &HttpServer{
		config: config,
		logger: logger,
	}
}

type Router interface {
	RegisterRouter(engine *gin.Engine)
}

type AppOptions struct {
	PrintVersion   bool
	ConfigFilePath string
}

func ResolveAppOptions(opt *AppOptions) {
	var printVersion bool
	var configFilePath string

	flag.BoolVar(&printVersion, "v", false, "-v print app version")
	flag.StringVar(&configFilePath, "c", "", "-c app configuration file")
	flag.Parse()

	opt.PrintVersion = printVersion
	opt.ConfigFilePath = configFilePath
}

func (s *HttpServer) registerRouter(g *gin.Engine, routers ...Router) {
	for _, r := range routers {
		r.RegisterRouter(g)
	}
}

func (s *HttpServer) RegisterOnShutdown(f func()) {
	s.onShutdown = f
}

func Ping(port string, maxCount int) error {
	seconds := 1
	if len(port) == 0 {
		log.Fatal("Please specify the service port")
	}
	if !strings.HasPrefix(port, ":") {
		port += ":"
	}
	url := fmt.Sprintf("http://localhost%s/ping", port)
	for i := 0; i < maxCount; i++ {
		resp, err := http.Get(url)

		if nil == err && resp != nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		log.Printf("Ping after %d second，%d times", seconds, maxCount)

		time.Sleep(time.Second * 1)
		seconds++
	}

	return fmt.Errorf("service is not start on port %s", port)
}

func (s *HttpServer) Run(routers ...Router) {
	var wg sync.WaitGroup
	wg.Add(1)
	gin.SetMode(s.config.Mode)
	g := gin.New()
	s.registerRouter(g, routers...)

	// TODO: register validator

	go func() {
		if err := Ping(s.config.Port, s.config.MaxPingCount); err != nil {
			log.Fatal("server no response")
		}
		log.Printf("server started success! port: %s", s.config.Port)
	}()

	srv := http.Server{
		Addr:    s.config.Port,
		Handler: g,
	}

	if s.onShutdown != nil {
		srv.RegisterOnShutdown(s.onShutdown)
	}

	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		defer wg.Done()
		<-sgn
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("server shutdown err %v \n", err)
		}
	}()

	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Printf("server started failed on port: %s", s.config.Port)
			return
		}
	}

	wg.Wait()
	log.Printf("server stop on port: %s", s.config.Port)
}
