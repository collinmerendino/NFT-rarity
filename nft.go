package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type NFTMetadata struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Image       string         `json:"image"`
	Attributes  []NFTAttribute `json:"attributes"`
}

type NFTAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type CollectionStats struct {
	Stats map[string]map[string]float64 `json:"stats"` // Format: { "trait_type": { "value": rarity_percentage } }
}

type RarityScore struct {
	NFTName  string  `json:"nft_name"`
	Score    float64 `json:"score"`
	Rank     int     `json:"rank"`
	Metadata NFTMetadata
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run nft.go <contract_address> <token_id>")
		return
	}

	contractAddress := os.Args[1]
	tokenID := os.Args[2]

	metadata, err := fetchNFTMetadata(contractAddress, tokenID)
	if err != nil {
		fmt.Println("Error fetching NFT metadata:", err)
		return
	}

	collectionStats, err := fetchCollectionStats(contractAddress)
	if err != nil {
		fmt.Println("Error fetching collection stats:", err)
		return
	}

	rarityScore := calculateRarityScore(metadata, collectionStats)

	fmt.Printf("NFT Name: %s\n", rarityScore.NFTName)
	fmt.Printf("Rarity Score: %.2f\n", rarityScore.Score)
	fmt.Printf("Rank: %d\n", rarityScore.Rank)
	fmt.Println("Traits:")
	for _, attr := range rarityScore.Metadata.Attributes {
		rarity := collectionStats.Stats[attr.TraitType][attr.Value]
		fmt.Printf("- %s: %s (Rarity: %.2f%%)\n", attr.TraitType, attr.Value, rarity)
	}
}

// fetchNFTMetadata metadata of the NFT using the OpenSea API
func fetchNFTMetadata(contractAddress, tokenID string) (NFTMetadata, error) {
	url := fmt.Sprintf("https://api.opensea.io/api/v1/asset/%s/%s/", contractAddress, tokenID)

	resp, err := http.Get(url)
	if err != nil {
		return NFTMetadata{}, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return NFTMetadata{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var metadata NFTMetadata
	err = json.Unmarshal(body, &metadata)
	if err != nil {
		return NFTMetadata{}, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	return metadata, nil
}

// fetchCollectionStats trait rarity stats for collection using OpenSea
func fetchCollectionStats(contractAddress string) (CollectionStats, error) {
	url := fmt.Sprintf("https://api.opensea.io/api/v1/asset/%s/stats/", contractAddress)

	resp, err := http.Get(url)
	if err != nil {
		return CollectionStats{}, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CollectionStats{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var stats CollectionStats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		return CollectionStats{}, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	return stats, nil
}

func calculateRarityScore(metadata NFTMetadata, stats CollectionStats) RarityScore {
	score := 0.0
	for _, attr := range metadata.Attributes {
		rarity := stats.Stats[attr.TraitType][attr.Value]
		score += rarity // Use rarity percentage as the score
	}

	return RarityScore{
		NFTName:  metadata.Name,
		Score:    score,
		Rank:     1,
		Metadata: metadata,
	}
}
