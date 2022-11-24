package eth

import (
	"math/big"

	"github.com/grassrootseconomics/w3-celo-patch/internal/module"
	"github.com/grassrootseconomics/w3-celo-patch/w3types"
)

// GasPrice requests the current gas price in wei.
func GasPrice() w3types.CallerFactory[big.Int] {
	return module.NewFactory(
		"eth_gasPrice",
		nil,
		module.WithRetWrapper(module.HexBigRetWrapper),
	)
}
