package virtual

import (
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestClient_GetTokens(t *testing.T) {
	url := "https://base-rpc.publicnode.com"

	rcp_client, err := ethclient.Dial(url)
	if err != nil {
		log.Printf("ethclient dial base node from %s failed, err %v", url, err)
		return
	}

	client := NewClient(rcp_client, Base)

	addr := common.HexToAddress("0x23462808DE54e6F1E3cBD2cA9d51EC95f66f41f1")
	tokens, err := client.GetTokens(addr)
	if err != nil {
		log.Printf("get tokens failed, err %v", err)
		return
	}
	log.Printf("tokens %v", tokens)
}
