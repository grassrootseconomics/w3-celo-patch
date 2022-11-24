package eth_test

import (
	"testing"

	"github.com/grassrootseconomics/w3-celo-patch/module/eth"
	"github.com/grassrootseconomics/w3-celo-patch/rpctest"
)

func TestChainID(t *testing.T) {
	tests := []rpctest.TestCase[uint64]{
		{
			Golden:  "chain_id",
			Call:    eth.ChainID(),
			WantRet: 1,
		},
	}

	rpctest.RunTestCases(t, tests)
}
