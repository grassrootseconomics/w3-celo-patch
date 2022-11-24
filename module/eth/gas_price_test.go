package eth_test

import (
	"math/big"
	"testing"

	"github.com/grassrootseconomics/w3-celo-patch"
	"github.com/grassrootseconomics/w3-celo-patch/module/eth"
	"github.com/grassrootseconomics/w3-celo-patch/rpctest"
)

func TestGasPrice(t *testing.T) {
	tests := []rpctest.TestCase[big.Int]{
		{
			Golden:  "gas_price",
			Call:    eth.GasPrice(),
			WantRet: *w3.I("0xc0fe"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
