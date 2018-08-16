package adapters

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func Setup() (common.Address, *types.Transaction, *DexBot, error) {
	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc)

	//Deploy contract
	address, trans, contract, err := DeployDexBot(
		auth,
		blockchain,
		big.NewInt(666),
	)

	// commit all pending transactions
	blockchain.Commit()

	return address, trans, contract, err
}

func TestDeployDexBot(t *testing.T) {
	address, _, _, err := Setup()

	if err != nil {
		t.Fatalf("Failed to deploy the DexBot contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}
}

func TestNumber(t *testing.T) {
	_, _, contract, _ := Setup()

	if got, _ := contract.Number(nil); got.Cmp(big.NewInt(666)) != 0 {
		t.Errorf("Expected number to be: 666. Got: %s", got)
	}
}

func TestSetNumber(t *testing.T) {
	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc)

	//Deploy contract
	_, _, contract, _ := DeployDexBot(
		auth,
		blockchain,
		big.NewInt(666),
	)

	// commit all pending transactions
	blockchain.Commit()

	contract.SetNumber(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  nil,
	}, big.NewInt(777))

	blockchain.Commit()

	if got, _ := contract.Number(nil); got.Cmp(big.NewInt(777)) != 0 {
		t.Errorf("Expected number to be: 777. Got: %s", got)
	}
}
