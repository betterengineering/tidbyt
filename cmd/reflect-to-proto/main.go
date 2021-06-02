package main

import (
	"context"
	"log"

	"github.com/betterengineering/tidbyt/pkg/tidbyt"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoprint"
	"github.com/jhump/protoreflect/grpcreflect"
	rpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

// reflect-to-proto uses reflection on the Tidbyt gRPC interface to get the proto definitions for the services it
// exposes. These are used in scripts/gen-proto.sh to generate the Tidbyt client. If there is ever an official client,
// we can simply remove this functionality.
func main() {
	cfg, err := tidbyt.NewConfigFromEnv()
	if err != nil {
		log.Fatalf("could not get config: %s\n", err)
	}

	conn, err := tidbyt.NewTidbytAPIConn(cfg.Token)
	if err != nil {
		log.Fatalf("could not create connection: %s\n", err)
	}
	defer conn.Close()

	stub := rpb.NewServerReflectionClient(conn)
	client := grpcreflect.NewClient(context.Background(), stub)

	services, err := client.ListServices()
	if err != nil {
		log.Fatalf("could not list services: %s\n", err)
	}

	descriptors := []*desc.FileDescriptor{}
	for _, service := range services {
		descr, err := client.ResolveService(service)
		if err != nil {
			log.Fatalf("could not resolve service: %s\n", err)
		}
		descriptors = append(descriptors, descr.GetFile())
	}

	p := protoprint.Printer{}
	err = p.PrintProtosToFileSystem(descriptors, "api/")
	if err != nil {
		log.Fatalf("could not write files: %s\n", err)
	}
}
