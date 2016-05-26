package main

import (
	"fmt"
	// "golang.org/x/net/context"
	// "log"
	"os"
)

// gcp up
func gcpUp() {
	fmt.Println("GCP UP")
}

// gcp down
func gcpDown() {
	fmt.Println("GCP DOWN")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: devc [up|down] [gcp|aws]")
		return
	}

	// ctx := context.Background()
	// s, err := createService(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	action := os.Args[1]
	provider := os.Args[2]

	if action == "up" {
		switch {
		case provider == "aws":
			awsUp()
		case provider == "gpc":
			gcpUp()
		}
	}

	if action == "down" {
		switch {
		case provider == "aws":
			awsDown()
		case provider == "gpc":
			gcpDown()
		}
	}

	if action == "list" {
		switch {
		case provider == "aws":
			awsList()
		case provider == "gpc":
			gcpList()
		}
	}

	fmt.Println("calling " + provider + " to " + action)
}
