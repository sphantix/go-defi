package uniswapv3

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	uniswapv3pool "github.com/sphantix/go-defi/bindings/uniswap/v3pool"
	"github.com/sphantix/go-defi/utils"
)

const (
	Ethereum = iota
	Base
	Unsupported
)

// Client allows to do operations on uniswap smart contracts.
type Client struct {
	bc    utils.Blockchain
	chain int
}

// NewClient returns a new instance of uniswap client.
func NewClient(bc utils.Blockchain, chain int) *Client {
	if chain >= Unsupported {
		panic("unsupported chain, use correct chain id")
	}

	return &Client{
		bc:    bc,
		chain: chain,
	}
}

// GetLiquidityFromPool returns the liquidity of the pool.
func (c *Client) GetLiquidityFromPool(pool common.Address) (*big.Int, error) {
	caller, err := uniswapv3pool.NewUniv3poolCaller(pool, c.bc)
	if err != nil {
		return nil, errors.Wrap(err, "new uniswap v3 pool caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	liquidity, err := caller.Liquidity(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get liquidity")
	}

	return liquidity, nil
}
