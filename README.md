# Lottery Raffle

## Overview
This project is a smart contract-based lottery raffle system built using Solidity and Foundry. Participants can enter the raffle by purchasing tickets, and a winner is selected randomly based on the Chainlink VRF.

## Features
- Decentralized lottery system
- Random winner selection using Chainlink VRF
- Secure and transparent transactions
- Automated execution with smart contracts

## Tech Stack
- **Solidity**: Smart contract development
- **Foundry**: Smart contract testing and deployment
- **Chainlink VRF**: Random number generation
- **Ethereum**: Blockchain platform

## Installation
Ensure you have Foundry installed. If not, install it using:
```sh
curl -L https://foundry.paradigm.xyz | bash
foundryup
```
Clone the repository:
```sh
git clone https://github.com/rajat-sharma-Dev/Lottery-Raffle-2.git
cd Lottery-Raffle-2
```
Install dependencies:
```sh
forge install
```

## Usage
### Run Tests
```sh
forge test
```

### Deploy Contract
Update the environment variables in a `.env` file and deploy:
```sh
forge create --rpc-url <RPC_URL> --private-key <PRIVATE_KEY> src/Lottery.sol:Lottery
```

## Makefile
This project includes a `Makefile` to streamline development. Below are the available commands:

### Setup and Cleanup
- **Install dependencies**: `make install`
- **Remove unnecessary files**: `make remove`
- **Clean compiled artifacts**: `make clean`
- **Update dependencies**: `make update`

### Development Commands
- **Compile smart contracts**: `make build`
- **Run tests**: `make test`
- **Format Solidity files**: `make format`
- **Take a snapshot**: `make snapshot`

### Blockchain Interaction
- **Run Anvil local blockchain**: `make anvil`
- **Deploy contract**: `make deploy ARGS="--network sepolia"`
- **Create Chainlink subscription**: `make createSubscription`
- **Add a consumer to the subscription**: `make addConsumer`
- **Fund the Chainlink subscription**: `make fundSubscription`


