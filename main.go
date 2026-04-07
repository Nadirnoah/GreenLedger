package main

import (
	"fmt"
	"greenledger" // This points to your folder name
)

func main() {
	fmt.Println("--- GreenLedger Protocol Simulation ---")

	// 1. Simulate a project developer minting a new asset
	myAsset := greenledger.MintAsset("GL-001", "0xYourWalletAddress", "Amazon Rainforest", 500.5)
	fmt.Printf("Asset Created: %s at %s. Verified: %v\n", myAsset.ID, myAsset.Location, myAsset.IsVerified)

	// 2. Simulate a Validator verifying the asset
	verifiedAsset := greenledger.VerifyAsset(myAsset, "0xValidatorAddress")
	fmt.Printf("Asset Updated: %s. Verified: %v\n", verifiedAsset.ID, verifiedAsset.IsVerified)
	
	fmt.Println("--- Simulation Complete ---")
}