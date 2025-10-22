package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	v1 "testconnect/gen/test/v1"
	svc "testconnect/gen/test/v1/testv1connect"

	"connectrpc.com/connect"
)

func main() {
	// Use gRPC transport option
	client := svc.NewSimpleGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		connect.WithGRPC(), // <— tells the client to speak gRPC protocol
	)

	fmt.Println("Enter a name to send")
	fmt.Println("Enter quit to quit")

	var input string

	for strings.ToLower(input) != "quit" {
		fmt.Scanln(&input)
		callgRPC(input, client)
	}

}

func callgRPC(name string, client svc.SimpleGreetServiceClient) {
	req := connect.NewRequest(&v1.ClientRequest{Name: name})
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response:", resp.Msg.Name)
}
