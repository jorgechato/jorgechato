package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	stack := &Stack{
		Languages: []string{"Go", "TS", "Python", "Java"},
		Platforms: []string{"Cloudflare", "Kafka", "AWS"},
		Misc:      []string{"gRPC", "K8s", "Terraform"},
		Ongoing:   []string{"three.js", "NixOs", "日本語"},
		Tools:     []string{"NeoVim", "Jetbrains"},
	}

	jsonData, err := json.Marshal(stack)
	if err != nil {
		log.Fatalf("Failed to marshal stack: %v", err)
	}

	fmt.Println(string(jsonData))
}
