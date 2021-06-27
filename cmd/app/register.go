package main

import (
	"github.com/AnyKey/service/client"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	userGrpcDelivery "github.com/AnyKey/service/user/delivery/grpc"
	userHttpDelivery "github.com/AnyKey/service/user/delivery/http"
	userRediceRepository "github.com/AnyKey/service/user/repository/redis"
	userUseCase "github.com/AnyKey/service/user/usecase"
)

// register usecase, delivery, repository for each entity
func register(router *mux.Router,s *grpc.Server, rdb *redis.Client) {
	emailRegister(router, s, rdb)
	client.Template(router)
}

// register user entity
func emailRegister(router *mux.Router,s *grpc.Server, rdb *redis.Client) {
	RedisRepo := userRediceRepository.New(rdb)
	HttpDelivery := userHttpDelivery.New()

	userUCase := userUseCase.New(RedisRepo, HttpDelivery)

	go userGrpcDelivery.Launch(s, userUCase)

	userHttpDelivery.Launch(router, userUCase)

}
