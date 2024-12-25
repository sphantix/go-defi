package pancakev2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// FactoryAddressEth points to the uniswap factory.
var FactoryAddressEth = common.HexToAddress("0x1097053Fd2ea711dad45caCcc45EfF7548fCB362")
var FactoryAddressBase = common.HexToAddress("0x02a84c1b3BBD7401a5f7fa98a384EBC70bB5749E")
var FactoryAddressBSC = common.HexToAddress("0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73")

// Router02Address points to the uniswap v2 02 router.
var Router02AddressEth = common.HexToAddress("0xEfF92A263d31888d860bD50809A8D171709b7b1c")
var Router02AddressBase = common.HexToAddress("0x8cFe327CEc66d1C090Dd72bd0FF11d690C33a2Eb")
var Router02AddressBSC = common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E")

var pairAddressSuffix = common.FromHex("0x00fb7f630766e6a796048ea87d01acd3068e8ff67d078148a3fa3f4a84f69bd5")

// GeneratePairAddress generates a pair address for the given tokens
func GeneratePairAddress(token0, token1 common.Address, chain int) common.Address {
	// addresses need to be sorted in an ascending order for proper behaviour
	token0, token1 = sortAddressess(token0, token1)

	// 计算salt (keccak256(token0 + token1))
	var salt [32]byte
	copy(salt[:], crypto.Keccak256(
		append(token0.Bytes(), token1.Bytes()...),
	))

	factoryAddress := FactoryAddressEth
	if chain == Base {
		factoryAddress = FactoryAddressBase
	} else if chain == BSC {
		factoryAddress = FactoryAddressBSC
	}

	return crypto.CreateAddress2(factoryAddress, salt, pairAddressSuffix)
}

// Quote gets the exchange quote for a given amount of tokens and the respective pairs reserves.
func Quote(amount, reserve0, reserve1 *big.Int) *big.Int {
	if reserve1.Cmp(big.NewInt(0)) <= 0 ||
		reserve0.Cmp(big.NewInt(0)) <= 0 ||
		amount.Cmp(big.NewInt(0)) <= 0 {

		return new(big.Int)
	}

	// amountB = amountA.mul(reserveB) / reserveA;
	multiplied := new(big.Int).Mul(amount, reserve1)
	res := new(big.Int).Div(multiplied, reserve0)
	return res
}

func sortAddressess(tkn0, tkn1 common.Address) (common.Address, common.Address) {
	token0Rep := big.NewInt(0).SetBytes(tkn0.Bytes())
	token1Rep := big.NewInt(0).SetBytes(tkn1.Bytes())

	if token0Rep.Cmp(token1Rep) > 0 {
		tkn0, tkn1 = tkn1, tkn0
	}

	return tkn0, tkn1
}

// Pair represents a token pair.
type Pair struct {
	Token0, Token1 common.Address
}

// GetPathPairs takes in the given token path and returns the corresponding pairs.
func GetPathPairs(tokens []common.Address) []Pair {
	pairs := make([]Pair, 0)

	for i := 1; i < len(tokens); i++ {
		pair := Pair{
			Token0: tokens[i-1],
			Token1: tokens[i],
		}

		pairs = append(pairs, pair)
	}

	return pairs
}
