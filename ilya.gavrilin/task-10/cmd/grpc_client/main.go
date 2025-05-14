package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	contactpb "task-10/gen/proto/api/proto/contact_manager/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := flag.String("server", "localhost:50051", "The server address in the format of host:port")
	flag.Parse()

	conn, err := grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := contactpb.NewContactServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if len(flag.Args()) < 1 {
		fmt.Println("Available commands: create, get, list, update, delete")
		os.Exit(1)
	}

	cmd := flag.Args()[0]

	switch cmd {
	case "create":
		createCmd := flag.NewFlagSet("create", flag.ExitOnError)
		name := createCmd.String("name", "", "Contact name")
		phone := createCmd.String("phone", "", "Contact phone")
		email := createCmd.String("email", "", "Contact email")
		createCmd.Parse(flag.Args()[1:])

		resp, err := client.CreateContact(ctx, &contactpb.CreateContactRequest{
			Name:  *name,
			Phone: *phone,
			Email: *email,
		})
		if err != nil {
			log.Fatalf("Could not create contact: %v", err)
		}
		fmt.Printf("Created contact: %+v\n", resp.Contact)

	case "get":
		getCmd := flag.NewFlagSet("get", flag.ExitOnError)
		id := getCmd.String("id", "", "Contact ID")
		getCmd.Parse(flag.Args()[1:])

		resp, err := client.GetContact(ctx, &contactpb.GetContactRequest{Id: *id})
		if err != nil {
			log.Fatalf("Could not get contact: %v", err)
		}
		fmt.Printf("Contact: %+v\n", resp.Contact)

	case "list":
		resp, err := client.ListContacts(ctx, &contactpb.ListContactsRequest{})
		if err != nil {
			log.Fatalf("Could not list contacts: %v", err)
		}
		fmt.Printf("Contacts: %+v\n", resp.Contacts)

	case "update":
		updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := updateCmd.String("id", "", "Contact ID")
		name := updateCmd.String("name", "", "Contact name")
		phone := updateCmd.String("phone", "", "Contact phone")
		email := updateCmd.String("email", "", "Contact email")
		updateCmd.Parse(flag.Args()[1:])

		resp, err := client.UpdateContact(ctx, &contactpb.UpdateContactRequest{
			Id:    *id,
			Name:  *name,
			Phone: *phone,
			Email: *email,
		})
		if err != nil {
			log.Fatalf("Could not update contact: %v", err)
		}
		fmt.Printf("Updated contact: %+v\n", resp.Contact)

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.String("id", "", "Contact ID")
		deleteCmd.Parse(flag.Args()[1:])

		resp, err := client.DeleteContact(ctx, &contactpb.GetContactRequest{Id: *id})
		if err != nil {
			log.Fatalf("Could not delete contact: %v", err)
		}
		fmt.Printf("Result: %s\n", resp.Message)

	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}
