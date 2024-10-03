package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Product struct {
	ProductID         string `json:"productID"`
	Name              string `json:"name"`
	ProdDesc          string `json:"description"`
	MfgDate           string `json:"manufacturingDate"`
	BatchNumber       string `json:"batchNumber"`
	Status            string `json:"status"`
	SupplyDate        string `json:"supplyDate,omitempty"`
	StorageLocation   string `json:"storageLocation,omitempty"`
	WholesaleDate     string `json:"wholesaleDate,omitempty"`
	WholesaleLocation string `json:"wholesaleLocation,omitempty"`
	Quantity          int    `json:"quantity,omitempty"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	initialProducts := []Product{
		{
			ProductID:   "101",
			Name:        "NOTEBOOK",
			ProdDesc:    "300 PAGES RULED",
			MfgDate:     "2020-12-29",
			BatchNumber: "122020",
			Status:      "SUPPLIED",
		},
	}

	for _, product := range initialProducts {
		if err := putState(ctx, product.ProductID, product); err != nil {
			return fmt.Errorf("failed to initialize ledger for product %s: %w", product.ProductID, err)
		}
	}

	return nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error creating supply chain smart contract: %s\n", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting supply chain smart contract: %s\n", err.Error())
	}
}
