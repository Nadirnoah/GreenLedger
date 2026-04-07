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

// MintAsset creates a new GreenAsset and saves it to the Canopy state
func MintAsset(id string, owner string, location string, tons float64) GreenAsset {
    // 1. Create the new asset object
    newAsset := GreenAsset{
        ID:         id,
        Owner:      owner,
        Location:   location,
        CarbonTons: tons,
        IsVerified: false, // Starts as unverified until a validator checks it
        CreatedAt:  time.Now(),
    }

    // 2. In a real Canopy app, we would now save 'newAsset' to the database
    // For your demo, this logic shows how the data is structured!
    return newAsset
}

// VerifyAsset allows an authorized validator to confirm the asset's legitimacy
func VerifyAsset(asset GreenAsset, validatorAddress string) GreenAsset {
	// In a real app, we would check if 'validatorAddress' is on an "Approved List"
	// For the demo, we will assume the validator is authorized.
	if validatorAddress != "" {
		asset.IsVerified = true
	}
	return asset
}