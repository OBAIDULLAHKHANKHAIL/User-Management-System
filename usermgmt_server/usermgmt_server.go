package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "example.com/go-usermgmt-grpc/pb"
	"google.golang.org/grpc"
)

//////////////////////////////////////////////////////////
// Defining port where we had to run this service
//////////////////////////////////////////////////////////

const (
	port = ":50051"
)

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		user_list: &pb.UserList{},
	}
}

//////////////////////////////////////////////////////////
// Embed protobuf unimplemented user management server
// This user management server which is implementation of grpc service
// And to register this type with grpc
//////////////////////////////////////////////////////////

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	user_list *pb.UserList
}

func (server *UserManagementServer) run() error {
	
}

//////////////////////////////////////////////////////////
// Service method
// Reciever function of rpc service method create new user
// Which we had defined in the usermgmt.proto file
// Reciever function in the user management server file
//////////////////////////////////////////////////////////

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(100))
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})

	// Server is listening on local host
	log.Printf("server listening at %v", lis.Addr())

	// start server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// func main(){
// 	lis, err := net.Listen("tcp", port)
// 	if err !=nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	s := grpc.NewServer()
// 	pb.RegisterUserManagementServer(s, &UserManagementServer{})

// 	// Server is listening on local host
// 	log.Printf("Server listening at %v", lis.Addr())

// 	// start server
// 	if err := s.Serve(lis); err != nil{
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
