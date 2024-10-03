const { Gateway, Wallets } = require('fabric-network');
const path = require('path');
const fs = require('fs');

/**
 * Connects to the Hyperledger Fabric network and returns the contract and gateway instances.
 * @param {string} channelName - The name of the channel to connect to.
 * @param {string} contractName - The name of the smart contract.
 * @param {string} orgName - The name of the organization.
 * @returns {Object} - An object containing the contract and gateway instances.
 */
const connectToNetwork = async (channelName, contractName, orgName) => {
    // Load network configuration
    const ccpPath = path.resolve(__dirname, '..', 'connection', `connection-${orgName}.json`);
    const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

    // Initialize wallet
    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);

    // Check for identity
    const identity = await wallet.get(orgName);
    if (!identity) {
        throw new Error(`An identity for the organization "${orgName}" does not exist in the wallet`);
    }

    // Create gateway and connect
    const gateway = new Gateway();
    await gateway.connect(ccp, {
        wallet,
        identity: orgName,
        discovery: { enabled: true, asLocalhost: true },
    });

    // Get network and contract
    const network = await gateway.getNetwork(channelName);
    const contract = network.getContract(contractName);

    return { contract, gateway };
};

module.exports = { connectToNetwork };
