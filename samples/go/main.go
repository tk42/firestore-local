package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

const HostEnvKey = "FIRESTORE_EMULATOR_HOST"
const Project = "test-project"

func main() {
	// Checking ENV
	_, exist := os.LookupEnv(HostEnvKey)
	if !exist {
		log.Fatalf("required env %s", HostEnvKey)
	}

	// Store
	ctx := context.Background()
	client, _ := firestore.NewClient(ctx, Project)
	doc, _, err := client.
		Collection("developer").
		Add(ctx, map[string]interface{}{
			"name":    "john",
			"age":     "20",
			"message": "Hello World",
		})

	if err != nil {
		log.Fatal(err)
	}

	// Success Message
	successMsg := fmt.Sprintf("Store is success, 'doc.ID' is %s, 'doc.Path' is %s", doc.ID, doc.Path)
	log.Println(successMsg)
}
