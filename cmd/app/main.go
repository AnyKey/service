package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"time"
)

func init() {
	// init config (viper)
	initConfig()
	// init logger (logrus)
	initLogger()
}

func main() {
	launchApp()
}

func launchApp() {
	var (
		grpcP       = viper.GetString("grpc.port")
		httpAddress = viper.GetString("server.address")
	)

	// creating singleton grpc server
	s := grpc.NewServer()

	// connect redis cient
	redisClient := initRedis()

	router := mux.NewRouter()

	register(router, s, redisClient)

	srv := &http.Server{
		Handler:      router,
		Addr:         httpAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// creating listener on port for http and grpc
	listener, err := net.Listen("tcp", grpcP)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// serve grpc
	log.Infof("Start serve grpc on %s port", grpcP)
	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// serve http
	log.Println("Serve http ON", httpAddress)
	log.Fatal(srv.ListenAndServe())
}
