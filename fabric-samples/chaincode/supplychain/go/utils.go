package main

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func putState(ctx contractapi.TransactionContextInterface, key string, data interface{}) error {
    dataAsBytes, err := json.Marshal(data)
    if err != nil {
        return fmt.Errorf("failed to marshal data for key %s: %w", key, err)
    }

    if err := ctx.GetStub().PutState(key, dataAsBytes); err != nil {
        return fmt.Errorf("failed to put state for key %s: %w", key, err)
    }

    return nil
}

func getState[T any](ctx contractapi.TransactionContextInterface, key string) (*T, error) {
    dataAsBytes, err := ctx.GetStub().GetState(key)
    if err != nil {
        return nil, fmt.Errorf("failed to read from world state for key %s: %w", key, err)
    }
    if dataAsBytes == nil {
        return nil, fmt.Errorf("product %s does not exist", key)
    }

    var data T
    if err := json.Unmarshal(dataAsBytes, &data); err != nil {
        return nil, fmt.Errorf("failed to unmarshal data for key %s: %w", key, err)
    }

    return &data, nil
}
