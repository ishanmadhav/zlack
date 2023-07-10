package api

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	pb "github.com/ishanmadhav/zlack/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserBody struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

func SetupUserRoutes(app *fiber.App) {
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("Hello, User!")
	})

	//Get User Routes
	app.Get("/users", GetAllUsers)
	app.Get("/user/:id", GetUserByID)
	app.Get("/user/:username", GetUserByUsername)

	//Post Routes
	app.Post("/user", CreateUser)

	//Put Routes
	app.Put("/user/:id", UpdateUser)

	//Delete Routes
	app.Delete("/user/:id", DeleteUserByID)
	//Utility Route, should be delete later
	app.Delete("/users", DeleteAllUsers)

}

func GetAllUsers(c *fiber.Ctx) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Printf("So it connecfted")
	defer conn.Close()
	gc := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := gc.GetAllUsers(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r)
	return c.JSON(r)

}

// // Function will make an gRPC call to the User service to get all users
// func GetAllUsersAlt(c *fiber.Ctx) error {
// 	// Set up a connection to the server.
// 	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	fmt.Printf("So it connecfted")
// 	defer conn.Close()
// 	gc := pb.NewGreeterClient(conn)

// 	// Contact the server and print out its response.
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	r, err := gc.SayHello(ctx, &pb.HelloRequest{Name: "User"})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}
// 	log.Printf("Greeting: %s", r.GetMessage())

// 	return c.SendString("Got Users")
// }

// Function will make an gRPC call to the User service to get a user by ID
func GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")
	fmt.Println(userID)
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	gc := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := gc.GetUserByID(ctx, &pb.GetUserByIDRequest{Id: userID})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Got user from gRPC")
	log.Printf("User: %s", r.GetUsername())
	return c.JSON(r)
}

// Function will make an gRPC call to the User service to create a user
func CreateUser(c *fiber.Ctx) error {
	var user UserBody
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(user)
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	gc := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := gc.CreateUser(ctx, &pb.CreateUserRequest{Username: user.Username, Email: user.Email, Password: user.Password, Firstname: user.Firstname, Lastname: user.Lastname})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Got user from gRPC")
	log.Printf("User: %s", r.GetUsername())

	return c.JSON(r)
}

// Function will make an gRPC call to the User service to update a user
func UpdateUser(c *fiber.Ctx) error {
	return c.SendString("Update User")
}

// Function will make an gRPC call to the User service to delete a user
func DeleteUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")
	fmt.Println(userID)
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	gc := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := gc.DeleteUserByID(ctx, &pb.DeleteUserByIDRequest{Id: userID})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Deleted user via gRPC")
	log.Printf("UserId: %s", r.GetId())
	return c.JSON(r)
}

// Function will make an gRPC call to the User service to get a user by username
func GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	fmt.Println(username)
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	gc := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := gc.GetUserByUsername(ctx, &pb.GetUserByUsernameRequest{Username: username})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println("Got user from gRPC")
	log.Printf("User: %s", r.GetUsername())
	return c.JSON(r)
}

// Function will make an gRPC call to the User service to delete all users
func DeleteAllUsers(c *fiber.Ctx) error {
	return c.SendString("Delete All Users")
}
