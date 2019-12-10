package main
//go get -u google.golang.org/grpc
import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"strings"
	"sunjinfu/api/proto"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("Can not connect gRPC server: %v", err)
	}
	defer conn.Close()
	c := proto.NewStudentServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := c.GetAllStudents(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Failed to invoke GetAllStudents, err: %v", err)
	}
	log.Printf("GetAllStudents response message: %v", *response)

	student, err := c.SearchStudent(ctx, &proto.SearchRequest{Name: "test"})
	if err != nil {
		log.Printf("Failed to invoke SearchStudent, err: %v", err)
	}
	log.Printf("SearchStudent response message: %v", *student)

	_, err = c.AddStudent(ctx, &proto.Student{
		Id: 10,
		Name: "darlyhounty",
		Address: &proto.Address{
			Street: "Load",
			Postcode: "100020",
		},
	})
	if err != nil {
		log.Printf("Failed to invoke AddStudent, err: %v", err)
	}

	stream, err := c.GetStudent(ctx)
	if err != nil {
		log.Printf("Failed to invoke  GetStudent, err: %v", err)
	} else {
		//发送10个消息
		for a := 0; a < 10; a++ {
			sr := &proto.SearchRequest{
				Id: strconv.Itoa(a),
				Name: strings.Join([]string{"name", strconv.Itoa(a)}, "")}
			err = stream.Send(sr)
			if err != nil {
				log.Printf("Failed to send message, err: %v", err)
			}
		}
		reply, err := stream.CloseAndRecv()
		if err != nil {
			log.Printf("Failed to get reply message from server, err: %v", reply)
		} else {
			log.Printf("GetStudent reply message: %v", *reply)
		}

	}
}