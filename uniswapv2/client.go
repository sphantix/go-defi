/* Mysterium network payment library.
 *
 * Copyright (C) 2020 BlockDev AG
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package uniswapv2

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	uniswapv2factory "github.com/sphantix/go-defi/bindings/uniswap/v2factory"
	uniswapv2pair "github.com/sphantix/go-defi/bindings/uniswap/v2pair"
	"github.com/sphantix/go-defi/utils"
)

const (
	Ethereum = iota
	Base
	BSC
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

// GetReserves retursn the available reserves in a pair
func (c *Client) GetReserves(token0, token1 common.Address) (*Reserve, error) {
	addr := GeneratePairAddress(token0, token1, c.chain)
	caller, err := uniswapv2pair.NewUniswapv2pairCaller(addr, c.bc)
	if err != nil {
		return nil, errors.Wrap(err, "new uniswap pair caller")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	reserves, err := caller.GetReserves(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get reserves")
	}
	// This is the tricky bit.
	// The reserve call returns the reserves for token0 and token1 in a sorted order.
	// This means we need to check if our token addresses are sorted or not and flip the reserves if they are not sorted.
	stoken0, _ := sortAddressess(token0, token1)
	if stoken0 != token0 {
		// We're not sorted, so the reserves need to be flipped to represent the actual reserves.
		reserves.Reserve0, reserves.Reserve1 = reserves.Reserve1, reserves.Reserve0
	}
	return &Reserve{Reserve0: reserves.Reserve0, Reserve1: reserves.Reserve1, BlockTimestampLast: reserves.BlockTimestampLast}, nil
}

// GetReserves retursn the available reserves in a pair
func (c *Client) GetReserveFromPair(addr common.Address) (*Reserve, error) {
	//addr := GeneratePairAddress(token0, token1)
	caller, err := uniswapv2pair.NewUniswapv2pairCaller(addr, c.bc)
	if err != nil {
		return nil, errors.Wrap(err, "new uniswap pair caller")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	reserves, err := caller.GetReserves(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get reserves")
	}
	return &Reserve{Reserve0: reserves.Reserve0, Reserve1: reserves.Reserve1, BlockTimestampLast: reserves.BlockTimestampLast}, nil
}

// GetReserves retursn the available reserves in a pair
func (c *Client) GetReservesInBlock(token0, token1 common.Address, bn int64) (*Reserve, error) {
	addr := GeneratePairAddress(token0, token1, c.chain)
	caller, err := uniswapv2pair.NewUniswapv2pairCaller(addr, c.bc)
	if err != nil {
		return nil, errors.Wrap(err, "new uniswap pair caller")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	reserves, err := caller.GetReserves(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: big.NewInt(bn),
	})
	if err != nil {
		return nil, errors.Wrap(err, "get reserves")
	}
	// This is the tricky bit.
	// The reserve call returns the reserves for token0 and token1 in a sorted order.
	// This means we need to check if our token addresses are sorted or not and flip the reserves if they are not sorted.
	stoken0, _ := sortAddressess(token0, token1)
	if stoken0 != token0 {
		// We're not sorted, so the reserves need to be flipped to represent the actual reserves.
		reserves.Reserve0, reserves.Reserve1 = reserves.Reserve1, reserves.Reserve0
	}
	return &Reserve{Reserve0: reserves.Reserve0, Reserve1: reserves.Reserve1, BlockTimestampLast: reserves.BlockTimestampLast}, nil
}

// GetReserves retursn the available reserves in a pair
func (c *Client) GetReservesInBlockByPair(pair common.Address, bn int64) (*Reserve, error) {
	caller, err := uniswapv2pair.NewUniswapv2pairCaller(pair, c.bc)
	if err != nil {
		return nil, errors.Wrap(err, "new uniswap pair caller")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	reserves, err := caller.GetReserves(&bind.CallOpts{
		Context:     ctx,
		BlockNumber: big.NewInt(bn),
	})
	if err != nil {
		return nil, errors.Wrap(err, "get reserves")
	}

	return &Reserve{Reserve0: reserves.Reserve0, Reserve1: reserves.Reserve1, BlockTimestampLast: reserves.BlockTimestampLast}, nil
}

func (c *Client) GetTokens(addr common.Address) ([]common.Address, error) {
	caller, err := uniswapv2pair.NewUniswapv2pairCaller(addr, c.bc)
	if err != nil {
		return nil, errors.Wrap(err, "new uniswap pair caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	token0, err := caller.Token0(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get token0 failed")
	}
	token1, err := caller.Token1(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get token1 failed")
	}
	return []common.Address{token0, token1}, nil
}

// GetFactory returns the factory address for a given pair
func (c *Client) GetFactory(addr common.Address) (common.Address, error) {
	caller, err := uniswapv2pair.NewUniswapv2pairCaller(addr, c.bc)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "new uniswap pair caller")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	factory, err := caller.Factory(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return common.Address{}, errors.Wrap(err, "get factory failed")
	}
	return factory, nil
}

// GetExchangeAmount returns the amount of tokens you'd receive when exchanging the given amount of token0 to token1.
func (c *Client) GetExchangeAmount(amount *big.Int, token0, token1 common.Address) (*big.Int, error) {
	reserves, err := c.GetReserves(token0, token1)
	if err != nil {
		return nil, errors.Wrap(err, "get reserves")
	}
	return Quote(amount, reserves.Reserve0, reserves.Reserve1), nil
}

// GetExchangeAmountForPath calculates the amount for a given path.
func (c *Client) GetExchangeAmountForPath(amount *big.Int, tokens ...common.Address) (*big.Int, error) {
	if len(tokens) <= 1 {
		return nil, errors.New("not enough tokens for path")
	}

	pairs := GetPathPairs(tokens)
	for i := range pairs {
		a, err := c.GetExchangeAmount(amount, pairs[i].Token0, pairs[i].Token1)
		if err != nil {
			return nil, errors.Wrap(err, "get exchange amount")
		}
		amount = a
	}

	return amount, nil
}

// Factory returns a uniswap cactory factory binding
func (c *Client) Factory() (*uniswapv2factory.Uniswapv2factory, error) {
	return uniswapv2factory.NewUniswapv2factory(FactoryAddressEth, c.bc)
}
