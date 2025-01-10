package virtual

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	virtualfpair "github.com/sphantix/go-defi/bindings/virtual/fpair"
	"github.com/sphantix/go-defi/utils"
)

const (
	Base = iota
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

func (c *Client) GetTokens(addr common.Address) ([]common.Address, error) {
	caller, err := virtualfpair.NewFPairCaller(addr, c.bc)
	if err != nil {
		return nil, errors.Wrap(err, "new fpair caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	tokenA, err := caller.TokenA(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get tokenA failed")
	}
	tokenB, err := caller.TokenB(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get tokenB failed")
	}
	return []common.Address{tokenA, tokenB}, nil
}
