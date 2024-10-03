// routes/productRoute.js

const express = require('express');
const router = express.Router();
const productController = require('../controllers/productController');

// Define Routes
router.post('/create', productController.createProduct);
router.post('/supply', productController.supplyProduct);
router.post('/wholesale', productController.wholesaleProduct);
router.get('/query/:productID', productController.queryProduct);
router.post('/sell', productController.sellProduct);

module.exports = router;
