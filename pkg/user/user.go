package main

import (
	"context"
	"fmt"
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

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedUserServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Println("Create User function is running")
	//Function to create a user
	db := DBConn
	newUser := models.User{
		Username:  in.Username,
		Email:     in.Email,
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Password:  in.Password}
	result := db.Create(&newUser)

	if result.Error != nil {
		return &pb.CreateUserResponse{}, result.Error
	}

	return &pb.CreateUserResponse{Username: newUser.Username, Email: newUser.Email, Firstname: newUser.Firstname, Lastname: newUser.Lastname}, nil
}

func (s *server) GetAllUsers(ctx context.Context, in *pb.Empty) (*pb.GetAllUsersResponse, error) {
	db := DBConn
	var users []models.User
	result := db.Find(&users)
	if result.Error != nil {
		return &pb.GetAllUsersResponse{}, result.Error
	}
	var pbUsers []*pb.GetUserResponse
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.GetUserResponse{Username: user.Username, Email: user.Email, Firstname: user.Firstname, Lastname: user.Lastname})
	}
	return &pb.GetAllUsersResponse{Users: pbUsers}, nil
}

func (s *server) GetUserByID(ctx context.Context, in *pb.GetUserByIDRequest) (*pb.GetUserResponse, error) {
	db := DBConn
	var user models.User
	result := db.First(&user, in.Id)
	if result.Error != nil {
		return &pb.GetUserResponse{}, result.Error
	}
	return &pb.GetUserResponse{Username: user.Username, Email: user.Email, Firstname: user.Firstname, Lastname: user.Lastname}, nil
}

func (s *server) GetUserByUsername(ctx context.Context, in *pb.GetUserByUsernameRequest) (*pb.GetUserResponse, error) {
	db := DBConn
	var user models.User
	result := db.First(&user, "username = ?", in.Username)
	if result.Error != nil {
		return &pb.GetUserResponse{}, result.Error
	}
	return &pb.GetUserResponse{Username: user.Username, Email: user.Email, Firstname: user.Firstname, Lastname: user.Lastname}, nil
}

func (s *server) DeleteUserByID(ctx context.Context, in *pb.DeleteUserByIDRequest) (*pb.DeleteUserResponse, error) {
	db := DBConn
	var user models.User
	result := db.Delete(&user, in.Id)
	if result.Error != nil {
		return &pb.DeleteUserResponse{}, result.Error
	}
	return &pb.DeleteUserResponse{Id: in.Id}, nil
}

func (s *server) DeleteAllUsers(ctx context.Context, in *pb.Empty) (*pb.DeleteAllUsersResponse, error) {
	db := DBConn
	var users []models.User
	result := db.Delete(&users)
	if result.Error != nil {
		return &pb.DeleteAllUsersResponse{}, result.Error
	}
	return &pb.DeleteAllUsersResponse{}, nil
}

// func (s *server) GetUserByID(ctx context.Context, in *pb.GetUserByIDRequest) (*pb.GetUserResponse, error) {
// 	//Function to get a user by ID
// }

// func (s *server) GetUserByUsername(ctx context.Context, in *pb.GetUserByUsernameRequest) (*pb.GetUserResponse, error) {
// 	//Function to get a user by Username
// }

func main() {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=zlack port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DBConn.AutoMigrate(&models.User{})

	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
