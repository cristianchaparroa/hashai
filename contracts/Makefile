# The .env file should be filled or the next commands will fail
include .env

# Check that required variables are set
ifndef PRIVATE_KEY
    $(error PRIVATE_KEY is undefined. Please set it in the .env file)
endif

ifndef API_KEY
    $(error API_KEY is undefined. Please set it in the .env file)
endif

to_snake_case = $(shell echo $(1) | sed -r 's/([a-z0-9])([A-Z])/\1_\L\2/g' | tr '[:upper:]' '[:lower:]')

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
	forge inspect $(contract) abi > ../backend/abi/$(contract).abi.json

# generate-go will create the contract representation in Golang language
# execute this command on the contracts folder
# make generate-go contract=YourContractName
generate-go:
	$(eval OUTPUT_NAME := $(call to_snake_case,$(contract)))
	forge inspect $(contract) abi > ../backend/abi/$(contract).abi.json
	forge inspect $(contract) bytecode > ../backend/abi/$(contract).bin
	abigen --abi ../backend/abi/$(contract).abi.json --bin ../backend/abi/$(contract).bin --pkg contracts --type $(contract) --out ../backend/internal/contracts/$(OUTPUT_NAME).go
