package api

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/order/service/pkg/api/handler"
	"github.com/order/service/pkg/pb"
	"google.golang.org/grpc"
)

type ServerHttp struct {
	Engine *gin.Engine
}

func NewGrpcServer(orderHandler *handler.OrderHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalln("Failed to listen to the grpc port: ", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterOrderManagementServer(grpcServer, orderHandler)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("Could not serve the grpc server", err)
	}

}

func NewServerHttp(orderHandler *handler.OrderHandler) *ServerHttp {
	engine := gin.New()

	engine.Use(gin.Logger())

	go NewGrpcServer(orderHandler, "8892")

	return &ServerHttp{
		Engine: engine,
	}
}

func (s *ServerHttp) Start() {
	s.Engine.Run(":7776")
}
