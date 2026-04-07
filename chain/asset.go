package greenledger

import "time"

// GreenAsset represents a real-world environmental project on the blockchain
type GreenAsset struct {
    ID              string    `json:"id"`               // Unique ID for the project
    Owner           string    `json:"owner"`            // Wallet address of the owner
    Location        string    `json:"location"`         // e.g., "Amazon Rainforest"
    CarbonTons      float64   `json:"carbon_tons"`      // Amount of CO2 sequestered
    IsVerified      bool      `json:"is_verified"`      // Has a validator signed off?
    CreatedAt       time.Time `json:"created_at"`       // When it was put on-chain
}