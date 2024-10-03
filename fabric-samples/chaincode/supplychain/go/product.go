package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *SmartContract) CreateProduct(ctx contractapi.TransactionContextInterface, productID, name, description, manufacturingDate, batchNumber string) error {
	existingProduct, err := s.QueryProduct(ctx, productID)
	if err == nil && existingProduct != nil {
		return fmt.Errorf("product with ID %s already exists", productID)
	}

	product := Product{
		ProductID:   productID,
		Name:        name,
		ProdDesc:    description,
		MfgDate:     manufacturingDate,
		BatchNumber: batchNumber,
		Status:      "Created",
	}

	return putState(ctx, productID, product)
}

func (s *SmartContract) SupplyProduct(ctx contractapi.TransactionContextInterface, productID, supplyDate, warehouseLocation string) error {
	product, err := s.QueryProduct(ctx, productID)
	if err != nil {
		return err
	}

	product.SupplyDate = supplyDate
	product.StorageLocation = warehouseLocation
	product.Status = "Supplied"

	return putState(ctx, productID, *product)
}

func (s *SmartContract) WholesaleProduct(ctx contractapi.TransactionContextInterface, productID, wholesaleDate, wholesaleLocation string, quantity int) error {
	product, err := s.QueryProduct(ctx, productID)
	if err != nil {
		return err
	}

	product.WholesaleDate = wholesaleDate
	product.WholesaleLocation = wholesaleLocation
	product.Quantity = quantity
	product.Status = "Wholesaled"

	return putState(ctx, productID, *product)
}

func (s *SmartContract) QueryProduct(ctx contractapi.TransactionContextInterface, productID string) (*Product, error) {
	product, err := getState[Product](ctx, productID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *SmartContract) UpdateProductStatus(ctx contractapi.TransactionContextInterface, productID, status string) error {
	product, err := s.QueryProduct(ctx, productID)
	if err != nil {
		return err
	}

	product.Status = status

	return putState(ctx, productID, *product)
}
