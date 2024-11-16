# HashTracker

## Overview

Hash#AI is a chatbot-based solution that simplifies wallet transaction tracking and makes detecting connections to suspicious or blacklisted addresses easier, safer, and more user-friendly.

## Local environment

### Backend
1. Install Golang
2. Run the project

```go
go run cmd/* 
```

### Endpoints

```bash
curl localhost:8080/transactions/0x8611C17eA68caE77762AdF6446a00f5a71dd7784?level=1
```
The following is the expected result

```
{
	"transactions": [
		{
			"hash": "0x6637a2ab8da7405f1b72dfc325ae80abea0f31149d8bcabf259e402d6a1a2e3e",
			"from": "0x8611C17eA68caE77762AdF6446a00f5a71dd7784",
			"to": "0x7cCD87331574Fe39140697962555848236C75DFA",
			"value": "10000000000000000",
			"timestamp": "2024-11-14T06:10:23Z"
		},
		{
			"hash": "0x396c3a2b3a63efa6ff81b2b268068a405bc4b159bd0c5946ca07880589e5058a",
			"from": "0x28C6c06298d514Db089934071355E5743bf21d60",
			"to": "0x8611C17eA68caE77762AdF6446a00f5a71dd7784",
			"value": "13009150000000000",
			"timestamp": "2024-11-14T06:08:23Z"
		},
	],
	"graph": {
		"nodes": [
			{
				"name": "0x8611C17eA68caE77762AdF6446a00f5a71dd7784"
			},
			{
				"name": "0x7cCD87331574Fe39140697962555848236C75DFA"
			},..
		],
		"links": [
			{
				"source": 0,
				"target": 1,
				"value": 10000000000000000
			},
			{
				"source": 2,
				"target": 0,
				"value": 13009150000000000
			},
			{
				"source": 0,
				"target": 3,
				"value": 3198986360403884
			},..
		]
	},
	"mermaid": "Z3JhcGggTFIKICAgIDB4ODYxMS4uLjc3ODQtLT58My4yMCBtRVRIfDB4MjUzNS4uLjMwM2IKICAgIDB4ODYxMS4uLjc3ODQtLT58MjAuMDAgbUVUSHwweDAwMDAuLi5iRTU5CiAgICAweERGZDUuLi45NjNkLS0+fDIuMzggbUVUSHwweDg2MTEuLi43Nzg0CiAgICAweDg2MTEuLi43Nzg0LS0+fDEwLjAwIG1FVEh8MHg3Y0NELi4uNURGQQogICAgMHgyOEM2Li4uMWQ2MC0tPnw0Mi43MyBtRVRIfDB4ODYxMS4uLjc3ODQ="
}
```

- Sankey Graph support
- Mermaid support


## Production Endpoints

```
GET https://hashtracker-em3r7.ondigitalocean.app/transactions/<WALLET_ADDRESS>?level=1 
GET https://hashtracker-em3r7.ondigitalocean.app/transactions/<ENS>?level=1 
```

## Services
The following are the list of services used.

### Converse Bot
https://converse.xyz/dm/0x61C9AB5968b49905dE120C699E140044ed77Bd2E

### Blockscout endpoints
[/addresses/{address_hash}/transactions](https://eth.blockscout.com/api-docs)

### Diagram Transactions visualization
```
https://hashtracker.vercel.app/?hash=graph+LR%0A0x3154...2C35--%3E%7C3.36+mETH%7C0xf418...EEEE%0A0x3154...2C35--%3E%7C613.41+%C2%B5ETH%7C0x6632...eeeE%0A0x2535...303b--%3E%7C46.15+mETH%7C0x3154...2C35%0A0x3154...2C35--%3E%7C5.00+mETH%7C0x82E0...Ec8A%0A0x7D1A...eBB0--%3E%7C10.12+mETH%7C0x3154...2C35%0A0x3154...2C35--%3E%7C4.00+mETH%7C0x9Ae7...0a29%0A0xEc56...df66--%3E%7C10.34+mETH%7C0x3154...2C35%0A0x3154...2C35--%3E%7C40.00+mETH%7C0xdAC1...1ec7
```


### Subgraphs
ENS Resolution: https://api.thegraph.com/subgraphs/name/ensdomains/ens

#### Creation of a New Subgraph for Indexing Blacklisted Wallets
We developed a new subgraph to index events emitted by a custom smart contract designed to store user-reported addresses. This contract records addresses associated with scams or malicious activities, as reported by users. Each time an address is reported, the contract emits an event containing the reported address, the category of the issue, and the total number of reports for that address. The subgraph indexes these events, enabling efficient querying and retrieval of blacklisted addresses and their associated data.

Example Query
You can test the subgraph using the following curl command:

Test prompt
``` bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "{ reportCreateds(first: 5) { id reportedAddress count category } }", "operationName": "Subgraphs", "variables": {}}' \
  https://api.studio.thegraph.com/query/95028/blacklisted_addresses_v0/version/latest
```
This query fetches the first 5 reportCreated events, including the following details:

- `id`: Unique identifier of the event.
- `reportedAddress`: The wallet address reported by users.
- `count`: The number of times the address has been reported.
- `category`: The type of issue associated with the address (e.g., scam, phishing, fraud).

### Polygon contracts
We deployed a HashReporter Smart Contract on the Polygon Amoy network. This contract is designed to store reported addresses and emit events, enabling efficient indexing in our subgraph powered by The Graph.
You can explore the deployed contract using the Polygon Amoy Explorer at the following URL:
https://amoy.polygonscan.com/address/0x2651F6e80a4295c59Cbb9260A05db591c988676e

## 👨🏻‍💻 Meet the Team
- [**Cristian Chaparro**](https://github.com/cristianchaparroa)
- [**Daniel Calderón**](https://github.com/danielcdz)
- [**Robert Ramirez**](https://github.com/robertram)
- [**Julio Cesar Arango**](https://www.linkedin.com/in/julio-cesar-bog-eth/)
- [**Valentina Díaz**](https://www.linkedin.com/in/valentina-diaz-estevez/)