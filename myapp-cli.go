package main

import (
//	"io"
	"log"
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "echotest/echotest"
)

const (
	address     = "192.168.34.72:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)
	//客户端的流
	stream,_ := c.Echo(context.Background())

	//并发线程接收并打印echo
	go func(){
		for true{	
		   echoreply,_ := stream.Recv()
		   if echoreply.Output != ""{
			fmt.Println(echoreply.Output)
		   }
		   
		   if 	echoreply.Nowtime != ""{
			   fmt.Println(echoreply)
		   }
	    }
    }()

	for true{
		//发送echo请求流
		var input string
		fmt.Scanln(&input)

		echorequest := &pb.EchoRequest{}
		echorequest.Input=input
		stream.Send(echorequest)
	}

}