package main

import (
	"log"
	"net"

	pb "github.com/ishanmadhav/zlack/api"
	"github.com/ishanmadhav/zlack/internal/models"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

type server struct {
	pb.UnimplementedWorkspaceServiceServer
}

//Should include several channels insde
//Users in one workspace should be able to contact each other via dms

func main() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=zlack port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DBConn.AutoMigrate(&models.Workspace{})

	lis, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWorkspaceServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
