# The .env file should be filled or the next commands will fail
include .env

# Check that required variables are set
ifndef PRIVATE_KEY
    $(error PRIVATE_KEY is undefined. Please set it in the .env file)
endif

ifndef API_KEY
    $(error API_KEY is undefined. Please set it in the .env file)
endif

# Command to deploy and verify a Smart Contract on the Polygon Amoy testnet
deploy-verify:
	@if [ -z "$(contract)" ]; then \
		echo "Error: contract variable is not set. Use 'make deploy-verify contract=YourContractName'"; \
		exit 1; \
	fi
	forge create --rpc-url $(RPC_URL) \
		--private-key $(PRIVATE_KEY) \
		--etherscan-api-key $(API_KEY) \
		--verify \
		"src/$(contract).sol:$(contract)"

# make abi contract=YourContractName
abi:
	jq -r .abi ../out/$(contract).sol/$(contract).json > ../backend/abi/$(contract).abi.json
