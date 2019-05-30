## Audit and Compliance in Service Management using Blockchain

Final year project sponsored by BMC Software to demonstrate the use case of Blockchain in Service Management sector.

### Prerequisites and setup:

* [Docker](https://www.docker.com/products/overview) - v1.12 or higher
* [Docker Compose](https://docs.docker.com/compose/overview/) - v1.8 or higher
* [Git client](https://git-scm.com/downloads) - needed for clone commands
* **Node.js** v8.4.0 or higher
* [Download Docker images](http://hyperledger-fabric.readthedocs.io/en/latest/samples.html#binaries)

```
cd Audit_project
```

Once you have completed the above setup, you will have provisioned a local network with the following docker container configuration:

* 3 CAs
* A SOLO orderer
* 6 peers (2 peers per Org)

#### Artifacts
* Crypto material has been generated using the **cryptogen** tool from Hyperledger Fabric and mounted to all peers, the orderering node and CA containers. More details regarding the cryptogen tool are available [here](http://hyperledger-fabric.readthedocs.io/en/latest/build_network.html#crypto-generator).
* An Orderer genesis block (genesis.block) and channel configuration transaction (mychannel.tx) has been pre generated using the **configtxgen** tool from Hyperledger Fabric and placed within the artifacts folder. More details regarding the configtxgen tool are available [here](http://hyperledger-fabric.readthedocs.io/en/latest/build_network.html#configuration-transaction-generator).

## Running the sample program

The app can be started using the following setup. The chaincode is written in golang.

##### Terminal Window 1

```
cd Audit_project

./runApp.sh

```

* This launches the required network on your local machine
* Installs the fabric-client and fabric-ca-client node modules
* And, starts the node app on PORT 4000

##### Terminal Window 2

```
cd front_end

python app2.py

```

##### User Interface via Browser

After this, access the available functions for license and service management via localhost:5000/frontpage 
