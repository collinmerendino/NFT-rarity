# NFT-rarity

A Go program to check the rarity of an NFT based on its metadata and trait rarity. 
It fetches NFT metadata and trait rarity data from the OpenSea API and calculates a rarity score.

---

## Features
- Fetches NFT metadata using the OpenSea API.
- Fetches trait rarity data for the NFT collection.
- Calculates a rarity score based on trait rarity.
- Displays the NFT's rarity score, rank, and traits with their rarity percentages.

---

## Prerequisites
- Go installed on your machine.
- An OpenSea API key (optional)

---

## Usage
```
go run nft.go <contract_address> <token_id>
```

Example
```
go run nft.go 0xContractAddress 384
```

## Output
```
NFT Name: NFT #384
Rarity Score: 15.75
Rank: 1
Traits:
- Background: Blue (Rarity: 5.25%)
- Hat: Red (Rarity: 3.50%)
- Eyes: Green (Rarity: 7.00%)
```

## Contributing
Contributions are welcome! If you have suggestions for improvements, please open an issue or submit a pull request.

This project is licensed under the MIT License.
