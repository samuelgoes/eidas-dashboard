package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	service "github.com/samuelgoes/eidas-dashboard/app/eidas"
)

type PublicKeyResponse struct {
	PublicKey string `json:"public_key"`
}

func getPublicKey() {
	// Replace these with actual values
	apiEndpoint := "https://example.com/api/get_public_key"
	apiKey := "your_api_key"
	tspID := "europe_tsp_id"

	// Create HTTP client
	client := http.Client{}

	// Prepare request
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Query parameters
	q := req.URL.Query()
	q.Add("tsp_id", tspID)
	req.URL.RawQuery = q.Encode()

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected response status:", resp.Status)
		fmt.Println("Response body:", string(body))
		return
	}

	// Parse response JSON
	var publicKeyResp PublicKeyResponse
	if err := json.Unmarshal(body, &publicKeyResp); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Print public key
	fmt.Println("Public key of Europe TSP:", publicKeyResp.PublicKey)
}

func trustedList() {
	// Get the trusted list from the EFDA API
	trustedList, err := service.GetTrustedList()
	if err != nil {
		fmt.Println("Error GetTrustedList:", err)
		return
	}
	fmt.Println("Trusted List: ", trustedList)
}

func main() {
	trustedList()
}
