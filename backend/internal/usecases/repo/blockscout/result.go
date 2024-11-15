package blockscout

import "time"

type TransactionResult struct {
	Items          []*Item     `json:"items"`
	NextPageParams *TxNextPage `json:"next_page_params"`
}

type Item struct {
	Timestamp              time.Time     `json:"timestamp"`
	Fee                    *Fee          `json:"fee"`
	GasLimit               string        `json:"gas_limit"`
	Block                  int           `json:"block"`
	Status                 string        `json:"status"`
	Method                 interface{}   `json:"method"`
	Confirmations          int           `json:"confirmations"`
	Type                   int           `json:"type"`
	ExchangeRate           string        `json:"exchange_rate"`
	To                     *Account      `json:"to"`
	TxBurntFee             string        `json:"tx_burnt_fee"`
	MaxFeePerGas           string        `json:"max_fee_per_gas"`
	Result                 string        `json:"result"`
	Hash                   string        `json:"hash"`
	GasPrice               string        `json:"gas_price"`
	PriorityFee            string        `json:"priority_fee"`
	BaseFeePerGas          string        `json:"base_fee_per_gas"`
	From                   *Account      `json:"from"`
	TokenTransfers         interface{}   `json:"token_transfers"`
	TxTypes                []string      `json:"tx_types"`
	GasUsed                string        `json:"gas_used"`
	CreatedContract        interface{}   `json:"created_contract"`
	Position               int           `json:"position"`
	Nonce                  int           `json:"nonce"`
	HasErrorInInternalTxs  bool          `json:"has_error_in_internal_txs"`
	Actions                []interface{} `json:"actions"`
	DecodedInput           interface{}   `json:"decoded_input"`
	TokenTransfersOverflow interface{}   `json:"token_transfers_overflow"`
	RawInput               string        `json:"raw_input"`
	Value                  string        `json:"value"`
	MaxPriorityFeePerGas   string        `json:"max_priority_fee_per_gas"`
	RevertReason           interface{}   `json:"revert_reason"`
	TxTag                  interface{}   `json:"tx_tag"`
}

type Fee struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Account struct {
	EnsDomainName   interface{}   `json:"ens_domain_name"`
	Hash            string        `json:"hash"`
	Implementations []interface{} `json:"implementations"`
	IsContract      bool          `json:"is_contract"`
	IsVerified      bool          `json:"is_verified"`
	Metadata        interface{}   `json:"metadata"`
	Name            interface{}   `json:"name"`
	PrivateTags     []interface{} `json:"private_tags"`
	ProxyType       interface{}   `json:"proxy_type"`
	PublicTags      []interface{} `json:"public_tags"`
	WatchlistNames  []interface{} `json:"watchlist_names"`
}

type TxNextPage struct {
	BlockNumber int       `json:"block_number"`
	Fee         string    `json:"fee"`
	Hash        string    `json:"hash"`
	Index       int       `json:"index"`
	InsertedAt  time.Time `json:"inserted_at"`
	ItemsCount  int       `json:"items_count"`
	Limit       string    `json:"limit"`
	Value       string    `json:"value"`
}
