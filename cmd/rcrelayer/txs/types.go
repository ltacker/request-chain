package txs

import (
	"math/big"

	"github.com/ltacker/request-chain/cmd/rcrelayer/events"
	"github.com/ethereum/go-ethereum/common"
)

// OracleClaim : contains data required to make an OracleClaim
type OracleClaim struct {
	ProphecyID *big.Int
	Message    string
	Signature  []byte
}

// ProphecyClaim : contains data required to make an ProphecyClaim
type ProphecyClaim struct {
	ClaimType            events.Event
	CosmosSender         []byte
	EthereumReceiver     common.Address
	TokenContractAddress common.Address
	Symbol               string
	Amount               *big.Int
}
