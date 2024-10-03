# Setting Up the Test Network

You can easily start a basic Fabric test network using the `./network.sh` script. This network includes two organizations, each with one peer, and a single-node Raft ordering service. The same script can also help you create channels and deploy chaincode. For more details, check out the [Using the Fabric Test Network](https://hyperledger-fabric.readthedocs.io/en/latest/test_network.html) guide. This test network replaces the older `first-network` example starting from Fabric v2.0.

Before you start the test network, make sure to follow the instructions in the [Install the Samples, Binaries, and Docker Images](https://hyperledger-fabric.readthedocs.io/en/latest/install.html) section of the Hyperledger Fabric documentation.

## Using Peer Commands

To set up the environment for your organizations, use the `setOrgEnv.sh` script. This makes it easier to run `peer` commands directly.

First, add the peer binaries to your PATH and set the Fabric configuration path. Make sure you're in the `test-network` directory:

```bash
export PATH=$PATH:$(realpath ../bin)
export FABRIC_CFG_PATH=$(realpath ../config)
```

Now, you can set the environment variables for each organization by running:

```bash
export $(./setOrgEnv.sh Org2 | xargs)
```

(Note: You need Bash version 4 or higher for these scripts.)

Now you can run `peer` commands for Org2. If you want to switch to Org1, just use the same command for that organization.

The `setOrgEnv` script generates a list of `<name>=<value>` strings that you can export to your shell.

## Chaincode as a Service

To learn about the updates in Chaincode-as-a-Service, take a look at this [tutorial](./test-network/../CHAINCODE_AS_A_SERVICE_TUTORIAL.md). This content is expected to enhance the tutorials in the [Hyperledger Fabric ReadTheDocs](https://hyperledger-fabric.readthedocs.io/en/release-2.4/cc_service.html).

## Using Podman

*Note: Podman support is still experimental. It's best to use a Linux VM for this.*

The `install-fabric.sh` script now supports using Podman to manage images instead of Docker. To use Podman, just run the script with the 'podman' argument.

The `network.sh` script has also been updated to work with Podman and `podman-compose`. Before running the script, set the `CONTAINER_CLI` environment variable to `podman`:

```bash
CONTAINER_CLI=podman ./network.sh up
```

Since Podman does not use a Docker daemon, only the `./network.sh deployCCAAS` command will work as intended. You can follow the Chaincode-as-a-Service tutorial mentioned earlier for guidance.