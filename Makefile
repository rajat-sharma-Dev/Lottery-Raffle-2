-include .env

.PHONY: help install clean remove update build test format anvil snapshot deploy createSubscription addConsumer fundSubscription

# Default private key for Anvil
DEFAULT_ANVIL_KEY := 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

help:
	@echo "Available commands:"
	@echo "  make deploy [ARGS=...] \n    example: make deploy ARGS=\"--network sepolia\""
	@echo "  make fundSubscription [ARGS=...] \n    example: make fundSubscription ARGS=\"--network sepolia\""

# Setup the project
install :; forge install foundry-rs/forge-std@v1.8.2 --no-commit && forge install transmissions11/solmate@v6 --no-commit && forge install smartcontractkit/chainlink-brownie-contracts@1.1.1 --no-commit && forge install cyfrin/foundry-devops@0.2.2 --no-commit

# Remove unnecessary files and dependencies
remove :; rm -rf .gitmodules && rm -rf .git/modules/* && rm -rf lib && touch .gitmodules && git add . && git commit -m "Removed submodules"

# Clean up compiled artifacts
clean  :; forge clean

# Update dependencies
update:; forge update

# Build smart contracts
build:; forge build

# Run tests
test :; forge test

# Format Solidity files
format :; forge fmt

# Run Anvil local blockchain
anvil :; anvil -m 'test test test test test test test test test test test junk' --steps-tracing --block-time 1

# Take a snapshot of the test state
snapshot :; forge snapshot

# Set default network arguments for local testing
NETWORK_ARGS := --rpc-url http://localhost:8545 --private-key $(DEFAULT_ANVIL_KEY) --broadcast

# Configure network arguments for Sepolia
ifeq ($(findstring --network sepolia,$(ARGS)),--network sepolia)
	NETWORK_ARGS := --rpc-url $(SEPOLIA_RPC_URL) --private-key $(PRIVATE_KEY) --broadcast --verify --etherscan-api-key $(ETHERSCAN_API_KEY) -vvvv
endif

# Deploy the Raffle contract
deploy:
	@forge script script/DeployRaffle.s.sol:DeployRaffle $(NETWORK_ARGS)

# Create a Chainlink subscription
createSubscription:
	@forge script script/Interactions.s.sol:CreateSubscription $(NETWORK_ARGS)

# Add a consumer to the Chainlink subscription
addConsumer:
	@forge script script/Interactions.s.sol:AddConsumer $(NETWORK_ARGS)

# Fund the Chainlink subscription
fundSubscription:
	@forge script script/Interactions.s.sol:FundSubscription $(NETWORK_ARGS)
