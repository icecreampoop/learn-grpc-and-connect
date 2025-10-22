package main

import (
	"context"
	"log"
	"net/http"
	v1 "testconnect/gen/test/v1"
	testv1connect "testconnect/gen/test/v1/testv1connect"

	"connectrpc.com/connect"
)

type simpleGreetServer struct {
	testv1connect.UnimplementedSimpleGreetServiceHandler
}

func (s *simpleGreetServer) SayHello(ctx context.Context, req *connect.Request[v1.ClientRequest]) (*connect.Response[v1.ServerResponse], error) {
	return connect.NewResponse(&v1.ServerResponse{
		Name: "Hello, " + req.Msg.Name,
	}), nil
}

func main() {
	mux := http.NewServeMux()
	path, handler := testv1connect.NewSimpleGreetServiceHandler(&simpleGreetServer{})
	mux.Handle(path, handler)

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
