package adapters

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Ethereum holds the contract instances
type Ethereum struct {
	Client *ethclient.Client
}

// NewEthereum returns a connected client
func NewEthereum(url string) (*Ethereum, error) {
	bc, err := ethclient.Dial(url)
	if err != nil {
		log.Printf("Unable to connect to Ethereum network: %v\n", err)
		return nil, err
	}
	return &Ethereum{bc}, nil
}

// DexBotInstance returns a reference to the contract instance
func (e *Ethereum) DexBotInstance() *DexBot {
	// DexBot deployed address
	addr := "0x936795f5fa3b0e8b7d39b1a542e92021c3857674"

	// Create a new instance of the contract bound to a specific deployed contract
	contract, err := NewDexBot(common.HexToAddress(addr), e.Client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of DexBot: %v\n", addr)
	}
	log.Printf("Found contract: DexBot %v.\n", addr)

	return contract
}

// EtherDeltaInstance returns a reference to the contract instance
func (e *Ethereum) EtherDeltaInstance() *EtherDelta {
	// EtherDelta deployed address
	addr := "0x8d12a197cb00d4747a1fe03395095ce2a5cc6819"

	// Create a new instance of the contract bound to a specific deployed contract
	contract, err := NewEtherDelta(common.HexToAddress(addr), e.Client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of EtherDelta: %v\n", addr)
	}
	log.Printf("Found contract: EtherDelta %v.\n", addr)

	return contract
}

// TokenBalance returns balance of token
func (e *Ethereum) TokenBalance(name string) string {
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()

	bal, err := e.Client.BalanceAt(context.Background(), common.HexToAddress("0x7CB807347F3E4AE113a4D4128367936eaC78509A"), big.NewInt(0))
	if err != nil {
		log.Printf("Unable to get balance: %v\n", err)
		return "0"
	}

	return bal.String()
}
