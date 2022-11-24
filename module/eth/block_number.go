package eth

import (
	"math/big"

	"github.com/grassrootseconomics/w3-celo-patch/internal/module"
	"github.com/grassrootseconomics/w3-celo-patch/w3types"
)

// BlockNumber requests the number of the most recent block.
func BlockNumber() w3types.CallerFactory[big.Int] {
	return module.NewFactory(
		"eth_blockNumber",
		nil,
		module.WithRetWrapper(module.HexBigRetWrapper),
	)
}
