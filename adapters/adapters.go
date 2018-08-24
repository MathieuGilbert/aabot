package adapters

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mathieugilbert/aabot/adapters/bat"
	"github.com/mathieugilbert/aabot/adapters/dice"
	"github.com/mathieugilbert/aabot/adapters/etherdelta"
	"github.com/mathieugilbert/aabot/adapters/link"
	"github.com/mathieugilbert/aabot/adapters/omg"
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

// BATInstance returns a reference to the contract instance
func (e *Ethereum) BATInstance(addr string) *bat.BAToken {
	c, err := bat.NewBAToken(common.HexToAddress(addr), e.Client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance: %v\n", addr)
	}
	return c
}

// DICEInstance returns a reference to the contract instance
func (e *Ethereum) DICEInstance(addr string) *dice.EtherollToken {
	c, err := dice.NewEtherollToken(common.HexToAddress(addr), e.Client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance: %v\n", addr)
	}
	return c
}

// LINKInstance returns a reference to the contract instance
func (e *Ethereum) LINKInstance(addr string) *link.LinkToken {
	c, err := link.NewLinkToken(common.HexToAddress(addr), e.Client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance: %v\n", addr)
	}
	return c
}

// OMGInstance returns a reference to the contract instance
func (e *Ethereum) OMGInstance(addr string) *omg.OMGToken {
	c, err := omg.NewOMGToken(common.HexToAddress(addr), e.Client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance: %v\n", addr)
	}
	return c
}

// EtherDeltaInstance returns a reference to the contract instance
func (e *Ethereum) EtherDeltaInstance(addr string) *etherdelta.EtherDelta {
	contract, err := etherdelta.NewEtherDelta(common.HexToAddress(addr), e.Client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of EtherDelta: %v\n", addr)
	}
	return contract
}

// TokenBalance returns balance of token, 0 on error
func (e *Ethereum) TokenBalance(addr, tName, tAddr string) string {
	var b *big.Int
	var err error

	switch tName {
	case "ETH":
		b, err = e.Client.BalanceAt(context.Background(), common.HexToAddress(addr), nil)
	case "BAT":
		b, err = e.BATInstance(tAddr).BalanceOf(&bind.CallOpts{}, common.HexToAddress(addr))
	case "DICE":
		b, err = e.DICEInstance(tAddr).BalanceOf(&bind.CallOpts{}, common.HexToAddress(addr))
	case "LINK":
		b, err = e.LINKInstance(tAddr).BalanceOf(&bind.CallOpts{}, common.HexToAddress(addr))
	case "OMG":
		b, err = e.OMGInstance(tAddr).BalanceOf(&bind.CallOpts{}, common.HexToAddress(addr))
	}

	if err != nil {
		log.Printf("Error getting %v(%v) balance for %v: %v\n", tName, tAddr, addr, err)
		return "0"
	}

	return b.String()
}

// ExchangeTokenBalance returns balance of token on exchange for user, 0 on error
func (e *Ethereum) ExchangeTokenBalance(eName, eAddr, tAddr, uAddr string) string {
	var b *big.Int
	var err error

	switch eName {
	case "EtherDelta":
		b, err = e.EtherDeltaInstance(eAddr).BalanceOf(&bind.CallOpts{}, common.HexToAddress(tAddr), common.HexToAddress(uAddr))
	}

	if err != nil {
		log.Printf("Error getting token (%v) balance for %v: %v\n", tAddr, uAddr, err)
		return "0"
	}

	return b.String()
}
