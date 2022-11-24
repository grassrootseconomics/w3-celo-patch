package eth

import (
	"math/big"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/grassrootseconomics/w3-celo-patch/internal/module"
	"github.com/grassrootseconomics/w3-celo-patch/w3types"
)

// Code requests the code of the given common.Address addr at the given
// blockNumber. If blockNumber is nil, the code at the latest known block is
// requested.
func Code(addr common.Address, blockNumber *big.Int) w3types.CallerFactory[[]byte] {
	return module.NewFactory(
		"eth_getCode",
		[]any{addr, module.BlockNumberArg(blockNumber)},
		module.WithRetWrapper(module.HexBytesRetWrapper),
	)
}
