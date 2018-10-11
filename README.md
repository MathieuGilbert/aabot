# AABot

## Developer Setup

 * Run local Ethereum node: Ganache
 * truffle compile; truffle test


#### Generating .go adapter from contract .sol

```
abigen -sol ethereum/contracts/EtherDelta.sol -pkg etherdelta -out adapters/etherdelta/contract.go
abigen -sol ethereum/contracts/BAT.sol -pkg bat -out adapters/bat/contract.go
abigen -sol ethereum/contracts/LINK.sol -pkg link -out adapters/link/contract.go
abigen -sol ethereum/contracts/DICE.sol -pkg dice -out adapters/dice/contract.go
abigen -sol ethereum/contracts/OMG.sol -pkg omg -out adapters/omg/contract.go
```

#### Generate contract abi


```
solc --abi -o adapters/dexbot ethereum/contracts/DexBot.sol
```


#### Node Packages

Install dependencies: `npm install`

Live reloading: `npm run watch`


#### Kill process running on a port:

```
lsof -t -i tcp:5000 | xargs kill
```


#### TODO:
