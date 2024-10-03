# Supply Chain Management System on Hyperledger Fabric 2.4.9

## Project Overview

This project aims to develop a decentralized Supply Chain Management System using Hyperledger Fabric 2.4.9. The system will incorporate three organizations—Producer, Supplier, and Wholeseller—each with two peers. It will include three Raft ordering nodes, three channels for communication, and smart contracts to track the complete journey of products. Additionally, Fabric Gateway APIs will be developed for easy interaction via Postman.

## Objective

The objective is to create a permissioned blockchain network that facilitates the tracking of products across the supply chain, ensuring transparency and efficiency. Each organization will be able to manage their roles through smart contracts and interact through well-defined APIs.

## Problem Statement

You are tasked with developing a Supply Chain Management System involving three organizations: Producer, Supplier, and Wholeseller. Each organization will participate in a blockchain network with two peers and use three Raft-based ordering nodes. The project will involve creating channels for communication, deploying smart contracts to manage the product lifecycle, and providing APIs for interaction.

## Project Tasks

### 1. Network Setup

- **Hyperledger Fabric Configuration**:
  - Set up three organizations: Producer, Supplier, and Wholeseller.
  - Each organization will have 2 peers.
  - Implement Raft consensus with 3 ordering nodes for transaction management.

- **Channel Setup**:
  - **Channel1**: Shared by all three organizations.
  - **Channel2**: Between Producer and Supplier.
  - **Channel3**: Between Producer and Wholeseller.

- **Certificate Authorities**:
  - Establish CAs for each organization to manage identities.
  - Define anchor peers and ensure correct network topology for communication.

### 2. Chaincode (Smart Contract)

Create a smart contract to manage the entire product journey:

- **Product Lifecycle Functions**:
  - **Create Product**: Producer adds a new product with details like ID, name, description, manufacturing date, and batch number.
  - **Supply Product**: Supplier updates product details upon receipt with supply date and location.
  - **Wholesale Product**: Wholeseller updates product details upon receipt with wholesale date, location, and quantity.
  - **Query Product**: Allows any organization to trace product history from creation to sale.
  - **Update Product Status**: Updates status at each stage (created, supplied, wholesaled, sold).

- **Channel Interactions**:
  - **Channel1**: Used for querying product history and making final sales visible to all.
  - **Channel2**: Manages transactions between Producer and Supplier for product creation and supply.
  - **Channel3**: Handles transactions between Producer and Wholeseller for wholesale activities.
  
- **Endorsement Policies**: Ensure that the smart contract is deployed on all channels with appropriate endorsement policies.

### 3. Fabric Gateway & APIs

- **Fabric Gateway Implementation**: Set up a Fabric Gateway to interact with the network and smart contracts across channels.
  
- **API Development**: Create REST APIs for key operations:
  - **POST /createProduct**: Allows the Producer to add a new product (Channel2 & Channel3).
  - **POST /supplyProduct**: Updates product status when received by Supplier (Channel2).
  - **POST /wholesaleProduct**: Updates product status when received by Wholeseller (Channel3).
  - **GET /queryProduct/{productID}**: Retrieves the complete product history (Channel1).
  - **POST /sellProduct**: Marks the product as sold, updating the ledger (Channel1).

- **Testing with Postman**: Ensure all APIs function correctly and provide meaningful responses.

### 4. Raft Ordering Service

- **Implementation**: Set up the Raft ordering service with 3 nodes to ensure fault tolerance.
  
- **Consensus Configuration**: Configure the network to require consensus from at least two orderer nodes to maintain operation in case one fails.

### 5. Additional Requirements

- **Membership Service Providers (MSP)**: Set up MSP for each organization.
  
- **Endorsement Policies**: Ensure policies reflect the need for organizational endorsements in transactions.
  
- **Chaincode Event Handling**: Implement event notifications for relevant transactions (e.g., product creation, supply, sale).
