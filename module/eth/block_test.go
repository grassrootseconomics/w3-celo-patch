package eth_test

import (
	"errors"
	"math/big"
	"sync/atomic"
	"testing"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/rpctest"
)

// func TestBlockByHash(t *testing.T) {
// 	tests := []rpctest.TestCase[types.Block]{
// 		{
// 			Golden: "get_block_by_hash__1",
// 			Call:   eth.BlockByHash(w3.H("0x88e96d4537bea4d9c05d12549907b32561d3bf31f45aae734cdc119f13406cb6")),
// 			WantRet: *types.NewBlockWithHeader(&types.Header{
// 				ParentHash:  w3.H("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3"),
// 				Coinbase:    w3.A("0x05a56E2D52c817161883f50c441c3228CFe54d9f"),
// 				Root:        w3.H("0xd67e4d450343046425ae4271474353857ab860dbc0a1dde64b41b5cd3a532bf3"),
// 				TxHash:      w3.H("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
// 				ReceiptHash: w3.H("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
// 				Bloom:       types.BytesToBloom(w3.B("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
// 				GasUsed:     0x0,
// 				Time:        0x55ba4224,
// 				Extra:       w3.B("0x476574682f76312e302e302f6c696e75782f676f312e342e32"),
// 			}),
// 		},
// 		{
// 			Golden: "get_block_by_hash__46147",
// 			Call:   eth.BlockByHash(w3.H("0x4e3a3754410177e6937ef1f84bba68ea139e8d1a2258c5f85db9f1cd715a1bdd")),
// 			WantRet: *types.NewBlockWithHeader(&types.Header{
// 				ParentHash:  w3.H("0x5a41d0e66b4120775176c09fcf39e7c0520517a13d2b57b18d33d342df038bfc"),
// 				Coinbase:    w3.A("0xe6A7a1d47ff21B6321162AEA7C6CB457D5476Bca"),
// 				Root:        w3.H("0x0e0df2706b0a4fb8bd08c9246d472abbe850af446405d9eba1db41db18b4a169"),
// 				TxHash:      w3.H("0x4513310fcb9f6f616972a3b948dc5d547f280849a87ebb5af0191f98b87be598"),
// 				ReceiptHash: w3.H("0xfe2bf2a941abf41d72637e5b91750332a30283efd40c424dc522b77e6f0ed8c4"),
// 				Bloom:       types.BytesToBloom(w3.B("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
// 				Number:      w3.I("0xb443"),
// 				GasUsed:     0x5208,
// 				Time:        0x55c42659,
// 				Extra:       w3.B("0x657468706f6f6c2e6f7267"),
// 			}).WithBody(
// 				types.Transactions{
// 					types.NewTx(&types.LegacyTx{
// 						GasPrice: w3.I("0x2d79883d2000"),
// 						Gas:      0x5208,
// 						To:       w3.APtr("0x5DF9B87991262F6BA471F09758CDE1c0FC1De734"),
// 						Value:    w3.I("0x7a69"),
// 						V:        w3.I("0x1c"),
// 						R:        w3.I("0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0"),
// 						S:        w3.I("0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a"),
// 					}),
// 				}, nil, nil,
// 			),
// 		},
// 		{
// 			Golden:  "get_block_by_hash__0x00",
// 			Call:    eth.BlockByHash(common.Hash{}),
// 			WantErr: errors.New("w3: call failed: not found"),
// 		},
// 	}

// 	rpctest.RunTestCases(t, tests,
// 		cmp.AllowUnexported(types.Block{}, types.Transaction{}, atomic.Value{}),
// 		cmpopts.IgnoreFields(types.Transaction{}, "time"),
// 	)
// }

func TestBlockByNumber(t *testing.T) {
	tests := []rpctest.TestCase[types.Block]{
		{
			Golden: "get_block_by_number__1",
			Call:   eth.BlockByNumber(big.NewInt(1)),
			WantRet: *types.NewBlockWithHeader(&types.Header{
				ParentHash:  w3.H("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3"),
				Coinbase:    w3.A("0x05a56E2D52c817161883f50c441c3228CFe54d9f"),
				Root:        w3.H("0xd67e4d450343046425ae4271474353857ab860dbc0a1dde64b41b5cd3a532bf3"),
				TxHash:      w3.H("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
				ReceiptHash: w3.H("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"),
				Bloom:       types.BytesToBloom(w3.B("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
				Number:      w3.I("0x1"),
				GasUsed:     0x0,
				Time:        0x55ba4224,
				Extra:       w3.B("0x476574682f76312e302e302f6c696e75782f676f312e342e32"),
			}),
		},
		{
			Golden: "get_block_by_number__46147",
			Call:   eth.BlockByNumber(big.NewInt(46147)),
			WantRet: *types.NewBlockWithHeader(&types.Header{
				ParentHash:  w3.H("0x5a41d0e66b4120775176c09fcf39e7c0520517a13d2b57b18d33d342df038bfc"),
				Coinbase:    w3.A("0xe6A7a1d47ff21B6321162AEA7C6CB457D5476Bca"),
				Root:        w3.H("0x0e0df2706b0a4fb8bd08c9246d472abbe850af446405d9eba1db41db18b4a169"),
				TxHash:      w3.H("0x4513310fcb9f6f616972a3b948dc5d547f280849a87ebb5af0191f98b87be598"),
				ReceiptHash: w3.H("0xfe2bf2a941abf41d72637e5b91750332a30283efd40c424dc522b77e6f0ed8c4"),
				Bloom:       types.BytesToBloom(w3.B("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
				Number:      w3.I("0xb443"),
				GasUsed:     0x5208,
				Time:        0x55c42659,
				Extra:       w3.B("0x657468706f6f6c2e6f7267"),
			}).WithBody(
				types.Transactions{
					types.NewTx(&types.LegacyTx{
						Nonce:    0x0,
						GasPrice: w3.I("0x2d79883d2000"),
						Gas:      0x5208,
						To:       w3.APtr("0x5DF9B87991262F6BA471F09758CDE1c0FC1De734"),
						Value:    w3.I("0x7a69"),
						V:        w3.I("0x1c"),
						R:        w3.I("0x88ff6cf0fefd94db46111149ae4bfc179e9b94721fffd821d38d16464b3f71d0"),
						S:        w3.I("0x45e0aff800961cfce805daef7016b9b675c137a6a41a548f7b60a3484c06a33a"),
					}),
				}, nil, nil,
			),
		},
		{
			Golden:  "get_block_by_number__999999999",
			Call:    eth.BlockByNumber(big.NewInt(999999999)),
			WantErr: errors.New("w3: call failed: not found"),
		},
	}

	rpctest.RunTestCases(t, tests,
		cmp.AllowUnexported(types.Block{}, types.Transaction{}, atomic.Value{}),
		cmpopts.IgnoreFields(types.Transaction{}, "time"),
	)
}

func TestBlockTxCountByHash(t *testing.T) {
	tests := []rpctest.TestCase[uint]{
		{
			Golden:  "block_transaction_count_by_hash__15050000",
			Call:    eth.BlockTxCountByHash(w3.H("0xc43d35f6a64f8a64f046c8deb4069572d622dfe7f028f62301b186f08f0e96f2")),
			WantRet: 32,
		},
		{
			Golden:  "block_transaction_count_by_hash__0x00",
			Call:    eth.BlockTxCountByHash(common.Hash{}),
			WantErr: errors.New("w3: call failed: not found"),
		},
	}

	rpctest.RunTestCases(t, tests)
}

func TestBlockTxCountByNumber(t *testing.T) {
	tests := []rpctest.TestCase[uint]{
		{
			Golden:  "block_transaction_count_by_number__15050000",
			Call:    eth.BlockTxCountByNumber(big.NewInt(15050000)),
			WantRet: 32,
		},
	}

	rpctest.RunTestCases(t, tests)
}

func TestHeaderByHash(t *testing.T) {
	tests := []rpctest.TestCase[types.Header]{
		{
			Golden: "get_block_by_hash__12965000",
			Call:   eth.HeaderByHash(w3.H("0x9b83c12c69edb74f6c8dd5d052765c1adf940e320bd1291696e6fa07829eee71")),
			WantRet: types.Header{
				ParentHash:  w3.H("0x3de6bb3849a138e6ab0b83a3a00dc7433f1e83f7fd488e4bba78f2fe2631a633"),
				Coinbase:    w3.A("0x7777788200B672A42421017F65EDE4Fc759564C8"),
				Root:        w3.H("0x41cf6e8e60fd087d2b00360dc29e5bfb21959bce1f4c242fd1ad7c4da968eb87"),
				TxHash:      w3.H("0xdfcb68d3a3c41096f4a77569db7956e0a0e750fad185948e54789ea0e51779cb"),
				ReceiptHash: w3.H("0x8a8865cd785e2e9dfce7da83aca010b10b9af2abbd367114b236f149534c821d"),
				Bloom:       types.BytesToBloom(w3.B("0x24e74ad77d9a2b27bdb8f6d6f7f1cffdd8cfb47fdebd433f011f7dfcfbb7db638fadd5ff66ed134ede2879ce61149797fbcdf7b74f6b7de153ec61bdaffeeb7b59c3ed771a2fe9eaed8ac70e335e63ff2bfe239eaff8f94ca642fdf7ee5537965be99a440f53d2ce057dbf9932be9a7b9a82ffdffe4eeee1a66c4cfb99fe4540fbff936f97dde9f6bfd9f8cefda2fc174d23dfdb7d6f7dfef5f754fe6a7eec92efdbff779b5feff3beafebd7fd6e973afebe4f5d86f3aafb1f73bf1e1d0cdd796d89827edeffe8fb6ae6d7bf639ec5f5ff4c32f31f6b525b676c7cdf5e5c75bfd5b7bd1928b6f43aac7fa0f6336576e5f7b7dfb9e8ebbe6f6efe2f9dfe8b3f56")),
				Number:      w3.I("0xc5d488"),
				GasUsed:     0x1ca2629,
				Time:        0x610bdaa6,
				Extra:       w3.B("0x68747470733a2f2f7777772e6b7279707465782e6f7267"),
			},
		},
	}

	rpctest.RunTestCases(t, tests)
}

func TestHeaderByNumber(t *testing.T) {
	tests := []rpctest.TestCase[types.Header]{
		{
			Golden: "get_block_by_number__12965000",
			Call:   eth.HeaderByNumber(big.NewInt(12965000)),
			WantRet: types.Header{
				ParentHash:  w3.H("0x3de6bb3849a138e6ab0b83a3a00dc7433f1e83f7fd488e4bba78f2fe2631a633"),
				Coinbase:    w3.A("0x7777788200B672A42421017F65EDE4Fc759564C8"),
				Root:        w3.H("0x41cf6e8e60fd087d2b00360dc29e5bfb21959bce1f4c242fd1ad7c4da968eb87"),
				TxHash:      w3.H("0xdfcb68d3a3c41096f4a77569db7956e0a0e750fad185948e54789ea0e51779cb"),
				ReceiptHash: w3.H("0x8a8865cd785e2e9dfce7da83aca010b10b9af2abbd367114b236f149534c821d"),
				Bloom:       types.BytesToBloom(w3.B("0x24e74ad77d9a2b27bdb8f6d6f7f1cffdd8cfb47fdebd433f011f7dfcfbb7db638fadd5ff66ed134ede2879ce61149797fbcdf7b74f6b7de153ec61bdaffeeb7b59c3ed771a2fe9eaed8ac70e335e63ff2bfe239eaff8f94ca642fdf7ee5537965be99a440f53d2ce057dbf9932be9a7b9a82ffdffe4eeee1a66c4cfb99fe4540fbff936f97dde9f6bfd9f8cefda2fc174d23dfdb7d6f7dfef5f754fe6a7eec92efdbff779b5feff3beafebd7fd6e973afebe4f5d86f3aafb1f73bf1e1d0cdd796d89827edeffe8fb6ae6d7bf639ec5f5ff4c32f31f6b525b676c7cdf5e5c75bfd5b7bd1928b6f43aac7fa0f6336576e5f7b7dfb9e8ebbe6f6efe2f9dfe8b3f56")),
				Number:      w3.I("0xc5d488"),
				GasUsed:     0x1ca2629,
				Time:        0x610bdaa6,
				Extra:       w3.B("0x68747470733a2f2f7777772e6b7279707465782e6f7267"),
			},
		},
	}

	rpctest.RunTestCases(t, tests)
}
