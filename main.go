package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

const (
	baseURL = "http://127.0.0.1:8200/v1/kv/data/"
)

var secrets = []string{"my-secret", "my-secret2", "my-secret3", "my-secret4", "my-secret5"}

func main() {
	var wg sync.WaitGroup
	requestCounts := make(map[string]int)

	// Get the request counts from command line arguments
	for _, secret := range secrets {
		count := getRequestCount(secret)
		requestCounts[secret] = count
	}

	// Launch requests concurrently
	for _, secret := range secrets {
		wg.Add(1)
		go func(secret string, count int) {
			defer wg.Done()
			for i := 0; i < count; i++ {
				makeRequest(secret)
			}
		}(secret, requestCounts[secret])
	}

	wg.Wait()
}

func getRequestCount(secret string) int {
	count := 1 // Default count
	if len(os.Args) > 1 {
		// Example: go run main.go my-secret=5 my-secret2=10
		for _, arg := range os.Args[1:] {
			var s string
			var c int
			_, err := fmt.Sscanf(arg, "%[^=]=%d", &s, &c)
			if err == nil && s == secret {
				count = c
			}
		}
	}
	return count
}

func makeRequest(secret string) {
	vaultToken := os.Args[1]
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL+secret, nil)
	if err != nil {
		fmt.Printf("Failed to create request for %s: %v\n", secret, err)
		return
	}
	req.Header.Add("X-Vault-Token", vaultToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Request to %s failed: %v\n", secret, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Request to %s succeeded with status: %s\n", secret, resp.Status)
}
