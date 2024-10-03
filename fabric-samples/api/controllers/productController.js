const { connectToNetwork } = require('../fabric/gateway');
const { validationResult } = require('express-validator');

/**
 * Handles the creation of a new product.
 */
exports.createProduct =
    body('productID').notEmpty(),
    body('name').notEmpty(),
    async (req, res, next) => {
        const { productID, name, description, manufacturingDate, batchNumber } = req.body;


        try {
            const { contract, gateway } = await connectToNetwork('channel2', 'supplychain', 'Producer');
            await contract.submitTransaction('CreateProduct', productID, name, description, manufacturingDate, batchNumber);
            await gateway.disconnect();
            res.status(201).json({ message: 'Product created successfully.' });
        } catch (error) {
            next(error);
        }
    };

/**
 * Handles supplying an existing product.
 */
exports.supplyProduct = async (req, res, next) => {
    const { productID, supplyDate, warehouseLocation } = req.body;

    try {
        const { contract, gateway } = await connectToNetwork('channel2', 'supplychain', 'Supplier');
        await contract.submitTransaction('SupplyProduct', productID, supplyDate, warehouseLocation);
        await gateway.disconnect();
        res.status(200).json({ message: 'Product supplied successfully.' });
    } catch (error) {
        next(error);
    }
};

/**
 * Handles wholesaling a product.
 */
exports.wholesaleProduct = async (req, res, next) => {
    const { productID, wholesaleDate, wholesaleLocation, quantity } = req.body;

    try {
        const { contract, gateway } = await connectToNetwork('channel3', 'supplychain', 'Wholeseller');
        await contract.submitTransaction('WholesaleProduct', productID, wholesaleDate, wholesaleLocation, quantity.toString());
        await gateway.disconnect();
        res.status(200).json({ message: 'Product wholesaled successfully.' });
    } catch (error) {
        next(error);
    }
};

/**
 * Handles querying a product by its ID.
 */
exports.queryProduct = async (req, res, next) => {
    const { productID } = req.params;

    try {
        const { contract, gateway } = await connectToNetwork('channel1', 'supplychain', 'Producer');
        const result = await contract.evaluateTransaction('QueryProduct', productID);
        await gateway.disconnect();
        res.status(200).json(JSON.parse(result.toString()));
    } catch (error) {
        next(error);
    }
};

/**
 * Handles selling a product by updating its status.
 */
exports.sellProduct = async (req, res, next) => {
    const { productID, buyerInfo } = req.body;

    try {
        const { contract, gateway } = await connectToNetwork('channel1', 'supplychain', 'Wholeseller');
        await contract.submitTransaction('UpdateProductStatus', productID, 'Sold');
        // Assuming 'buyerInfo' needs to be handled separately or stored differently
        // If 'buyerInfo' is part of the chaincode, ensure the chaincode supports it.
        await gateway.disconnect();
        res.status(200).json({ message: 'Product sold successfully.' });
    } catch (error) {
        next(error);
    }
};
