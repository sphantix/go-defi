package testenv

import (
	"context"
	"testing"

	"github.com/sphantix/go-defi/utils"
	"github.com/stretchr/testify/require"
)

func TestTestenv(t *testing.T) {
	testenv, err := NewBlockchain(context.Background())
	require.NoError(t, err)
	testenv.SendETH(testenv.Auth.From, utils.ToWei("100.0", 18))
}
