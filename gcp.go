package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
)

func gcpList() {

	// Authentication is provided by the gcloud tool when running locally, and
	// by the associated service account when running on Compute Engine.
	client, err := google.DefaultClient(context.Background())
	if err != nil {
		log.Fatalf("Unable to get default client: %v", err)
	}
	fmt.Printf("Google compute client is OK ")
	log.Fatalf("Unable to get default client: %v", client)
}
