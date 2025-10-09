package store

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/vault/api"
)

func GetMongoDBConnectionString() string {
	token := "s.zYuzrx9WevJhjbubN4lGAT5J"
	vaultAddr := "http://localhost:8200"

	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: &http.Client{Timeout: 100 * time.Second}})
	if err != nil {
		fmt.Println("Error creating Vault client:", err)
		return ""
	}
	client.SetToken(token)
	println("Connected to vault")

	kv := client.KVv2("secret")

	// name is "mysqlsecret" (the entry you see in UI)
	secret, err := kv.Get(context.Background(), "mongodb")
	if err != nil {
		log.Fatalf("vault read: %v", err)
	}
	if secret == nil || secret.Data == nil {
		log.Fatal("vault: empty secret or no data")
	}

	uri, _ := secret.Data["uri"].(string)

	fmt.Println("uri:", uri)

	return uri
}
