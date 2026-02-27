package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GenerateShortLink(originalURL string) string {
	secretSalt := "DKDN,KTB,ANGLETHEDREAMGIRL"
	dataToHash := originalURL + secretSalt
	hash := sha256.Sum256([]byte(dataToHash))
	fullHashStr := hex.EncodeToString(hash[:])
	shortCode := fullHashStr[:8]

	return shortCode
}
func main() {
	targetURL := "THIS IS NOT HOME"
	fmt.Println("start")
	shortLink := GenerateShortLink(targetURL)

	fmt.Println("shortCode:%s\n", shortLink)
}
