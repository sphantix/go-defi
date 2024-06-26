package bclient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sphantix/go-defi/sushiswap"
	"github.com/sphantix/go-defi/testenv"
	"github.com/sphantix/go-defi/uniswapv2"
	"github.com/sphantix/go-defi/utils"
)

// BClient wraps ethclient and provides helper functions for commonly used functionality
type BClient struct {
	ctx    context.Context
	cancel context.CancelFunc
	bc     utils.Blockchain
	uc     *uniswapv2.Client
	sc     *sushiswap.Client
}

// New returns a new bclient implementation without any authentication
// as well as wrappers around the uniswap and sushiswap clients. All underlying
// types are exposed in addition to helper functions
func NewClient(ctx context.Context, bc utils.Blockchain, chain int) (*BClient, error) {
	ctx, cancel := context.WithCancel(ctx)
	return &BClient{
		ctx:    ctx,
		cancel: cancel,
		bc:     bc,
		uc:     uniswapv2.NewClient(bc, chain),
		sc:     sushiswap.NewClient(bc),
	}, nil
}

// SimulatedBackend attempts to conver the Blockchain interface to a simulated backend type
// returning an error if unable to type convert the interface. This likely indicates
// that an ethclient backend is being used
func (bc *BClient) SimulatedBackend() (*testenv.Testenv, error) {
	sn, ok := bc.bc.(*testenv.Testenv)
	if !ok {
		return nil, ErrNotSimulatedBackend
	}
	return sn, nil
}

// EthClient attempts to conver the Blockchain interface to an ethclient type
// returning an error if unable to type convert the interface. This likely indicates
// that a simulated backend is being used
func (bc *BClient) EthClient() (*ethclient.Client, error) {
	ec, ok := bc.bc.(*ethclient.Client)
	if !ok {
		return nil, ErrNotEthClient
	}
	return ec, nil
}

// ChainID performs a best effort at determining the chainid
// we do this by first attempting to retrieve the id from
// a simulated backend, and if that fails get the chain id
// from an ethclient
func (bc *BClient) ChainID() (*big.Int, error) {
	if sb, err := bc.SimulatedBackend(); err == nil {
		// get chainid
		return sb.Blockchain().Config().ChainID, nil
	} else {
		if ec, err := bc.EthClient(); err != nil {
			return nil, errors.Wrap(err, "eth client")
		} else {
			return ec.ChainID(bc.ctx)
		}
	}
}

// Blockchain returns the underlying blockchain interface
func (bc *BClient) Blockchain() utils.Blockchain { return bc.bc }

// CurrentBlock returns the current block known by the ethereum client
func (bc *BClient) CurrentBlock() (uint64, error) {
	ec, err := bc.EthClient()
	if err != nil {
		return 0, errors.Wrap(err, "eth client")
	}
	return ec.BlockNumber(bc.ctx)
}

// Sushiswap returns the underlying sushiswap client
func (bc *BClient) Sushiswap() *sushiswap.Client { return bc.sc }

// Uniswap returns the underlying uniswap client
func (bc *BClient) Uniswap() *uniswapv2.Client { return bc.uc }

func (bc *BClient) Context() context.Context { return bc.ctx }

// Close terminates the blockchain connection
func (bc *BClient) Close() {
	// error is nil which means we can close the client
	if ec, err := bc.EthClient(); err == nil {
		ec.Close()
	}
}
